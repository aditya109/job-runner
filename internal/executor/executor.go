package executor

import (
	cfg "github.com/aditya109/job-runner/pkg/config"
	log "github.com/aditya109/job-runner/pkg/logwrapper"
	st "github.com/aditya109/job-runner/pkg/suite"
)

var (
	logger *log.StandardLogger
)

func ExecutorSuite(suite *st.Suite, config *cfg.Config) error {
	logger = log.NewLogger(config.Server.Env)

	// do something here

	return nil
}
