package suite

import (
	"fmt"
	"log"
	"os"

	"github.com/aditya109/job-runner/pkg/helper"
	"gopkg.in/yaml.v2"
)

func GetCompiledSuite(specs []string) *Suite {
	var scenarios = [][]Scenario{}
	var rowCounter int32 = -1
	for _, spec := range specs {
		var scenario Scenario = getScenarioFromSpec(spec)
		if !scenario.Parallelism {
			scenarios = append(scenarios, []Scenario{})
			rowCounter += 1
		} 
		scenarios[rowCounter] = append(scenarios[rowCounter], scenario)
	}

	suite := Suite{
		scenarios,
	}

	return &suite
}

func getScenarioFromSpec(file string) Scenario {
	var scenario Scenario
	var specFilePath string = helper.GetAbsolutePath(fmt.Sprintf("/spec/%s", file))

	specFile, err := os.Open(specFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer specFile.Close()

	decoder := yaml.NewDecoder(specFile)
	err = decoder.Decode(&scenario)
	if err != nil {
		log.Fatal(err)
	}
	return scenario
}
