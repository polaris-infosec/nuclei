package types

import "context"

type ProgressEvent struct {
	Requests uint64
	Total    uint64
}

type KOLNucleiRunner struct {
	ProgressChannel   chan ProgressEvent
	JsonOutputChannel chan []byte
	Ctx               context.Context
	Cancel            context.CancelFunc
}

type NucleiOption struct {
	Target            string
	Templates         []string
	Debug             bool
	Timeout           int
	JSON              bool
	EnableProgressBar bool
	Output            string
	CustomHeaders     []string
	ProxyURL          string
}
