package main

import (
	"fmt"

	cfg "github.com/aditya109/job-runner/pkg/config"
	log "github.com/aditya109/job-runner/pkg/logwrapper"
	st "github.com/aditya109/job-runner/pkg/suite"
	exe "github.com/aditya109/job-runner/internal/executor"
)

var (
	config      *cfg.Config
	logger      *log.StandardLogger
	environment string
	suite       *st.Suite
)

func main() {
	config = cfg.GetConfiguration()           // retrieving configuration
	logger = log.NewLogger(config.Server.Env) // initializing logger
	environment = config.Server.Env           // getting environment from config
	logger.Info(fmt.Sprintf("starting runner 🔥 in %s mode", environment))

	// prepping suite object
	suite = st.GetCompiledSuite(config.Scenarios)

	// feed suite object into executor
	fmt.Println(suite)

	// #TODO:
	err := exe.ExecutorSuite(suite, config)
	if err != nil {
		logger.Error(err)
	}

}
