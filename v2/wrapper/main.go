package wrapper

import (
	"github.com/projectdiscovery/nuclei/v2/internal/runner"
	"github.com/projectdiscovery/nuclei/v2/wrapper/types"
	"github.com/rs/zerolog/log"
)

func RunNuclei(opts *types.NucleiOption) (*types.KOLEventChannel, error) {
	options := &runner.Options{
		Target:            opts.Target,
		Templates:         opts.Templates,
		Timeout:           opts.Timeout,
		JSON:              opts.JSON,
		EnableProgressBar: opts.EnableProgressBar,
		Output:            opts.Output,
		CustomHeaders:     opts.CustomHeaders,
		ProxyURL:          opts.ProxyURL,
		Retries:           1,
		RateLimit:         150,
		BulkSize:          25,
		TemplateThreads:   10,
		JSONRequests:      true,
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

	return nucleiRunner.KOLEventChannel, nil
}
