package types

type ProgressEvent struct {
	Requests uint64
	Total    uint64
}

type KOLEventChannel struct {
	Progress   chan ProgressEvent
	JsonOutput chan []byte
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
