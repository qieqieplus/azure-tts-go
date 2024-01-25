package azure_tts

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func getWebsocketEndpoint(region string) string {
	return fmt.Sprintf(WssUrl, region)
}

func getTokenEndpoint(region string) string {
	return fmt.Sprintf(TokenUrl, region)
}

func getVoicesEndpoint(region string) string {
	return fmt.Sprintf(VoicesUrl, region)
}

func getGUID() string {
	id := uuid.New()
	return strings.ToUpper(hex.EncodeToString(id[:]))
}
