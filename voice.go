package azure_tts

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/qieqieplus/azure-tts-go/protocol"
)

type VoiceList []protocol.VoiceItem

func ListVoices(ctx context.Context, subConfig *SubscriptionConfig) (VoiceList, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getVoicesEndpoint(subConfig.Region), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerKey, subConfig.Key)
	req.Header.Set(headerEncoding, acceptEncoding)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var voices VoiceList
	err = json.Unmarshal(body, &voices)
	if err != nil {
		return nil, err
	}

	return voices, nil
}
