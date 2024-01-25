package azure_tts

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Synthesizer struct {
	subConfig    *SubscriptionConfig
	systemConfig *SystemConfig
	audioConfig  *AudioConfig
	ctx          context.Context
}

func NewSynthesizer(ctx context.Context, sub *SubscriptionConfig) *Synthesizer {
	return &Synthesizer{
		subConfig:    sub,
		systemConfig: NewDefaultSystemConfig(),
		audioConfig:  NewDefaultAudioConfig(),
		ctx:          ctx,
	}
}

func (s *Synthesizer) newSession() (*synthesizeSession, error) {
	endpoint := getWebsocketEndpoint(s.subConfig.Region)
	queries := url.Values{
		// headerAuth:         []string{fmt.Sprintf("Bearer %s", ...)},
		headerConnectionId: []string{getGUID()},
	}
	wssUrl := endpoint + "?" + queries.Encode()
	dialer := websocket.Dialer{
		HandshakeTimeout:  10 * time.Second,
		EnableCompression: true,
	}
	conn, _, err := dialer.DialContext(s.ctx, wssUrl, http.Header{
		headerEncoding: []string{acceptEncoding},
		headerKey:      []string{s.subConfig.Key},
	})

	if err != nil {
		return nil, err
	}

	return &synthesizeSession{
		ctx:  s.ctx,
		sid:  getGUID(),
		conn: conn,
	}, nil
}

func (s *Synthesizer) speakAsync(sender func(sess *synthesizeSession) error) (task SynthesizeTask) {
	var (
		sess *synthesizeSession
		err  error
	)

	task = newSynthesizeTask(s.ctx)
	defer func() {
		if err != nil {
			task.Error <- err
			task.Close()
		}
	}()

	if sess, err = s.newSession(); err != nil {
		return
	}
	if err = sess.sendSystemConfig(s.systemConfig); err != nil {
		return
	}
	if err = sess.sendAudioConfig(s.audioConfig); err != nil {
		return
	}
	if err = sender(sess); err != nil {
		return
	}

	go sess.readLoop(task)

	return
}

func (s *Synthesizer) SetConfig(audioConfig *AudioConfig) {
	s.audioConfig = audioConfig
}

func (s *Synthesizer) SpeakTextAsync(lang, speaker, text string) (task SynthesizeTask) {
	return s.speakAsync(func(sess *synthesizeSession) error {
		return sess.sendTextSynthesis(lang, speaker, text)
	})
}

func (s *Synthesizer) SpeakSSMLAsync(ssml string) (task SynthesizeTask) {
	return s.speakAsync(func(sess *synthesizeSession) error {
		return sess.sendSSMLSynthesis(ssml)
	})
}
