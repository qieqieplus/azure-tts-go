package protocol

import "io"

const (
	PathSpeechConfig = "speech.config"
	PathAudioConfig  = "synthesis.context"
	PathSSML         = "ssml"

	PathTurnStart     = "turn.start"
	PathTurnEnd       = "turn.end"
	PathResponse      = "response"
	PathAudioMetadata = "audio.metadata"
	PathAudio         = "audio"
)

// User messages

type SpeechConfigMessage = JsonMessage[SpeechConfig]
type SynthesisContextMessage = JsonMessage[SynthesisContext]
type SSMLMessage = XmlMessage[SSML]

// System messages

type TurnStartMessage = JsonMessage[TurnStart]
type TurnEndMessage = JsonMessage[TurnEnd]
type ResponseMessage = JsonMessage[Response]
type AudioMetadataMessage = JsonMessage[AudioMetadata]

type EventData interface {
	TurnStartMessage | TurnEndMessage | ResponseMessage | AudioMetadataMessage
}

type IMessage interface {
	Decode(io.Reader) error
	Encode() ([]byte, error)
}
