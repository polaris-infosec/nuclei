package wrapper

import (
	"fmt"
	"github.com/projectdiscovery/nuclei/v2/internal/runner"
	"github.com/projectdiscovery/nuclei/v2/wrapper/model"
	"github.com/rs/zerolog/log"
)

func RunNuclei(job *model.Job, templateDir string, debug bool) error {
	options := &runner.Options{
		Target: job.Target.URL,
		Templates: []string {
			templateDir + "/nuclei-templates/cves",
			templateDir + "/nuclei-templates/default-logins",
			templateDir + "/nuclei-templates/dns",
			templateDir + "/nuclei-templates/exposed-panels",
			templateDir + "/nuclei-templates/exposed-tokens",
			templateDir + "/nuclei-templates/exposures",
			templateDir + "/nuclei-templates/fuzzing",
			templateDir + "/nuclei-templates/miscellaneous",
			templateDir + "/nuclei-templates/misconfiguration",
			templateDir + "/nuclei-templates/takeovers",
			templateDir + "/nuclei-templates/technologies",
			templateDir + "/nuclei-templates/vulnerabilities",
			templateDir + "/nuclei-templates/workflows",
		},
		Debug: debug,
		Timeout: 5,
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
