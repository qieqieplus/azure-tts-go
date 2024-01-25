package azure_tts

import (
	"bytes"
	"context"
	"os"
	"testing"
)

var (
	subConfig = NewDefaultSubscriptionConfig()
)

func TestSynthesizer(t *testing.T) {
	if !subConfig.Valid() {
		t.Skip("no config")
	}

	synthesizer := NewSynthesizer(context.Background(), subConfig)
	audioConfig := NewDefaultAudioConfig()
	audioConfig.Synthesis.Audio.MetadataOptions.WordBoundaryEnabled = true
	synthesizer.SetConfig(audioConfig)

	task := synthesizer.SpeakTextAsync("zh-CN", "zh-CN-XiaoxiaoNeural", "Hello, world! 你好，世界！")
	audioStream := &bytes.Buffer{}

	stop := false
	for !stop {
		select {
		case <-task.Done():
			stop = true
		case err := <-task.Error:
			t.Fatal(err)
		case event := <-task.Event:
			switch event.(type) {
			case *AudioMetadataEvent:
				metadata := event.(*AudioMetadataEvent)
				for _, item := range metadata.Metadata {
					t.Log(item.Type)
					t.Log(item.Data)
				}
			}
		case audio := <-task.Audio:
			audioStream.Write(audio)
		}
	}

	os.WriteFile("synth.mp3", audioStream.Bytes(), 0644)
}
