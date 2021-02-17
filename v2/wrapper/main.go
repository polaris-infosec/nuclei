package wrapper

import (
	"github.com/projectdiscovery/nuclei/v2/internal/progress"
	"github.com/projectdiscovery/nuclei/v2/internal/runner"
	"github.com/rs/zerolog/log"
)

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

type ProgressEvent struct {
	Requests uint64
	Total    uint64
}

func RunNuclei(opts *NucleiOption) (chan progress.ProgressEvent, error) {
	options := &runner.Options{
		Target:            opts.Target,
		Templates:         opts.Templates,
		Timeout:           opts.Timeout,
		JSON:              opts.JSON,
		EnableProgressBar: opts.EnableProgressBar,
		Output:            opts.Output,
		CustomHeaders:     opts.CustomHeaders,
		ProxyURL:          opts.ProxyURL,
	}

	nucleiRunner, err := runner.New(options)
	if err != nil {
		log.Info().Msgf("Could not create runner: %s\n", err)
		return nil, err
	}

	go func() {
		nucleiRunner.RunEnumeration()
		nucleiRunner.Close()
	}()

	return nucleiRunner.ProgressEventChannel, nil
}
