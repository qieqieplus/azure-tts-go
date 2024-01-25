package azure_tts

import (
	"os"
	"runtime"

	. "github.com/qieqieplus/azure-tts-go/protocol"
)

type SubscriptionConfig struct {
	Key    string
	Region string
}

func (c *SubscriptionConfig) Valid() bool {
	return c.Key != "" && c.Region != ""
}

func NewSubscriptionConfig(key, region string) *SubscriptionConfig {
	return &SubscriptionConfig{
		Key:    key,
		Region: region,
	}
}

func NewDefaultSubscriptionConfig() *SubscriptionConfig {
	return &SubscriptionConfig{
		Key:    os.Getenv("SPEECH_KEY"),
		Region: os.Getenv("SPEECH_REGION"),
	}
}

type SystemConfig SpeechConfig

func NewDefaultSystemConfig() *SystemConfig {
	return &SystemConfig{
		Context: SpeechConfigContext{
			System: SystemInfo{
				Name:    "SpeechSDK",
				Version: "1.34.0",
				Build:   "azure-tts-go",
				Lang:    "Go",
			},
			OS: OSInfo{
				Platform: runtime.GOOS,
				Name:     "Go",
				Version:  runtime.Version(),
			},
		},
	}
}

type AudioConfig SynthesisContext

func NewDefaultAudioConfig() *AudioConfig {
	return &AudioConfig{
		Synthesis: Synthesis{
			Audio: AudioOptions{
				OutputFormat: "audio-24khz-160kbitrate-mono-mp3",
			},
			Language: Language{
				AutoDetection: true,
			},
		},
	}
}
