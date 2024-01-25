package azure_tts

const (
	WssUrl    = "wss://%s.tts.speech.microsoft.com/cognitiveservices/websocket/v1"
	TokenUrl  = "https://%s.api.cognitive.microsoft.com/sts/v1.0/issueToken"
	VoicesUrl = "https://%s.tts.speech.microsoft.com/cognitiveservices/voices/list"
)

const (
	headerAuth         = "Authorization"
	headerKey          = "Ocp-Apim-Subscription-Key"
	headerEncoding     = "Accept-Encoding"
	headerContentType  = "Content-Type"
	headerPath         = "Path"
	headerTimestamp    = "X-Timestamp"
	headerConnectionId = "X-ConnectionId"
	headerRequestId    = "X-RequestId"
)

const (
	acceptEncoding = "gzip, deflate, br"
)

const (
	contentTypeJson = "application/json"
	contentTypeSSML = "application/ssml+xml"
)
