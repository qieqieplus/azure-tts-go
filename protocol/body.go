package protocol

import "encoding/xml"

type SpeechConfig struct {
	Context SpeechConfigContext `json:"context"`
}

type SpeechConfigContext struct {
	System SystemInfo `json:"system"`
	OS     OSInfo     `json:"os"`
}

type SystemInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Build   string `json:"build"`
	Lang    string `json:"lang"`
}

type OSInfo struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

type SynthesisContext struct {
	Synthesis Synthesis `json:"synthesis"`
}

type Synthesis struct {
	Audio    AudioOptions `json:"audio"`
	Language Language     `json:"language"`
}

type Language struct {
	AutoDetection bool `json:"autoDetection"`
}

type AudioOptions struct {
	MetadataOptions MetadataOptions `json:"metadataOptions"`
	OutputFormat    string          `json:"outputFormat"`
}

type MetadataOptions struct {
	BookmarkEnabled            bool `json:"bookmarkEnabled"`
	PunctuationBoundaryEnabled bool `json:"punctuationBoundaryEnabled"`
	SentenceBoundaryEnabled    bool `json:"sentenceBoundaryEnabled"`
	SessionEndEnabled          bool `json:"sessionEndEnabled"`
	VisemeEnabled              bool `json:"visemeEnabled"`
	WordBoundaryEnabled        bool `json:"wordBoundaryEnabled"`
}

type SSML struct {
	XMLName xml.Name `xml:"speak"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Mstts   string   `xml:"xmlns:mstts,attr"`
	Emo     string   `xml:"xmlns:emo,attr"`
	Lang    string   `xml:"xml:lang,attr"`
	Voice   Voice    `xml:"voice"`
}

// Voice represents the <voice> element in SSML.
type Voice struct {
	Name string `xml:"name,attr"`
	Text string `xml:",chardata"`
}

type TurnStart struct {
	Context TurnStartContext `json:"context"`
}

type TurnStartContext struct {
	ServiceTag string `json:"serviceTag"`
}

type TurnEnd struct{}

type Response struct {
	Context ResponseContext `json:"context"`
	Audio   AudioInfo       `json:"audio"`
}

type ResponseContext struct {
	ServiceTag string `json:"serviceTag"`
}

type AudioInfo struct {
	Type     string `json:"type"`
	StreamId string `json:"streamId"`
}

type AudioMetadata struct {
	Metadata []MetadataItem `json:"Metadata"`
}

type MetadataItem struct {
	Type string         `json:"Type"`
	Data NestedMetadata `json:"Data"`
}
type NestedMetadata struct {
	Offset          int64    `json:"Offset"`
	VisemeId        int      `json:"VisemeId"`
	IsLastAnimation bool     `json:"IsLastAnimation"`
	Duration        int64    `json:"Duration"`
	TextData        TextData `json:"text"`
}

type TextData struct {
	Text         string `json:"Text"`
	Length       int    `json:"Length"`
	BoundaryType string `json:"BoundaryType"`
}

type VoiceItem struct {
	Name            string `json:"Name"`
	DisplayName     string `json:"DisplayName"`
	LocalName       string `json:"LocalName"`
	ShortName       string `json:"ShortName"`
	Gender          string `json:"Gender"`
	Locale          string `json:"Locale"`
	LocaleName      string `json:"LocaleName"`
	SampleRateHertz string `json:"SampleRateHertz"`
	VoiceType       string `json:"VoiceType"`
	Status          string `json:"Status"`
	WordsPerMinute  string `json:"WordsPerMinute"`
}
