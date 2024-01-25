package azure_tts

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/qieqieplus/azure-tts-go/protocol"
)

type synthesizeSession struct {
	ctx  context.Context
	sid  string
	conn *websocket.Conn
}

func (s *synthesizeSession) sendSystemConfig(config *SystemConfig) error {
	data, err := newSystemConfigMessage(s.sid, config).Encode()
	if err != nil {
		return err
	}

	return s.conn.WriteMessage(websocket.TextMessage, data)
}

func (s *synthesizeSession) sendAudioConfig(config *AudioConfig) error {
	data, err := newAudioConfigMessage(s.sid, config).Encode()
	if err != nil {
		return err
	}
	return s.conn.WriteMessage(websocket.TextMessage, data)
}

func (s *synthesizeSession) sendTextSynthesis(lang, speaker, text string) error {
	data, err := newTextSynthesisMessage(s.sid, lang, speaker, text).Encode()
	if err != nil {
		return err
	}
	return s.conn.WriteMessage(websocket.TextMessage, data)
}

func (s *synthesizeSession) sendSSMLSynthesis(ssml string) error {
	data, err := newSSMLSynthesisMessage(s.sid, ssml).Encode()
	if err != nil {
		return err
	}
	return s.conn.WriteMessage(websocket.TextMessage, data)
}

func (s *synthesizeSession) readLoop(task SynthesizeTask) {
	defer func() {
		task.Close()
		_ = s.conn.Close()
	}()

	for {
		mtype, data, err := s.conn.ReadMessage()
		if err != nil {
			task.Error <- err
			return
		}

		rd := bytes.NewReader(data)

		switch mtype {
		case websocket.TextMessage:
			pm := &protocol.PlainMessage{}
			if err = pm.Decode(rd); err != nil {
				task.Error <- err
				return
			}

			switch pm.Header.Get(headerPath) {
			case protocol.PathTurnStart:
				handleEventAsync(task, pm.Body, &TurnStartEvent{})
			case protocol.PathResponse:
				handleEventAsync(task, pm.Body, &ResponseEvent{})
			case protocol.PathAudioMetadata:
				handleEventAsync(task, pm.Body, &AudioMetadataEvent{})
			case protocol.PathTurnEnd:
				handleEventAsync(task, pm.Body, &TurnEndEvent{})
				return
			}

		case websocket.BinaryMessage:
			bm := &protocol.BinaryMessage{}
			if err = bm.Decode(rd); err != nil {
				task.Error <- err
				return
			}
			task.Audio <- bm.Body
		}
	}
}

func handleEventAsync[T any](task SynthesizeTask, body []byte, v T) {
	err := json.Unmarshal(body, v)
	if err != nil {
		task.Error <- err
		return
	}
	task.Event <- v
}
