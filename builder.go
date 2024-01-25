package azure_tts

import (
	"net/http"
	"time"

	. "github.com/qieqieplus/azure-tts-go/protocol"
)

func commonHeaders(id, path, contentType string) http.Header {
	return http.Header{
		headerRequestId:   []string{id},
		headerPath:        []string{path},
		headerContentType: []string{contentType},
		headerTimestamp:   []string{time.Now().UTC().Format(time.RFC3339)},
	}
}

func newSystemConfigMessage(id string, conf *SystemConfig) IMessage {
	message := &SpeechConfigMessage{}
	message.Header = commonHeaders(id, PathSpeechConfig, contentTypeJson)
	message.Entity = (SpeechConfig)(*conf)
	return message
}

func newAudioConfigMessage(id string, conf *AudioConfig) IMessage {
	message := &SynthesisContextMessage{}
	message.Header = commonHeaders(id, PathAudioConfig, contentTypeJson)
	message.Entity = (SynthesisContext)(*conf)
	return message
}

func newTextSynthesisMessage(id, lang, speaker, text string) IMessage {
	message := &SSMLMessage{}
	message.Header = commonHeaders(id, PathSSML, contentTypeSSML)
	message.Entity = SSML{
		Version: "1.0",
		Xmlns:   "http://www.w3.org/2001/10/synthesis",
		Mstts:   "http://www.w3.org/2001/mstts",
		Emo:     "http://www.w3.org/2009/10/emotionml",
		Lang:    lang, // "en-US", "zh-CN"
		Voice: Voice{
			Name: speaker,
			Text: text,
		},
	}
	return message
}

func newSSMLSynthesisMessage(id, ssml string) IMessage {
	return &PlainMessage{
		Header: commonHeaders(id, PathSSML, contentTypeSSML),
		Body:   []byte(ssml),
	}
}
