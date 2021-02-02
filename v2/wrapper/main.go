package wrapper

import (
	"fmt"
	"github.com/projectdiscovery/nuclei/v2/internal/runner"
	"github.com/rs/zerolog/log"
)

type NucleiOption struct {
	Target string
	Templates []string
	Debug bool
	Timeout int
	JSON bool
	EnableProgressBar bool
	Output string
	CustomHeaders []string
}

func RunNuclei(opts *NucleiOption) error {
	options := &runner.Options{
		Target: opts.Target,
		Templates: opts.Templates,
		Debug: opts.Debug,
		Timeout: opts.Timeout,
		JSON: opts.JSON,
		EnableProgressBar: opts.EnableProgressBar,
		Output: opts.Output,
		CustomHeaders: opts.CustomHeaders,
	}

	nucleiRunner, err := runner.New(options)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("Could not create runner: %s\n", err))
		return err
	}

	nucleiRunner.RunEnumeration()
	nucleiRunner.Close()

	return nil
}
