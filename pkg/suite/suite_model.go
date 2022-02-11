package suite

// Suite contains the scenario queue - containing both sequential and parallel
type Suite struct {
	Scenarios [][]Scenario
}

type Scenario struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Parallelism bool   `yaml:"parallelism"`
	Steps       []Step `yaml:"steps"`
}

type Step struct {
	Name                       string   `yaml:"name"`
	Type                       string   `yaml:"type"`
	Api                        string   `yaml:"api"`
	Method                     string   `yaml:"method"`
	Headers                    []Header `yaml:"headers"`
	PersistentKeysFromResponse []string `yaml:"persistentKeysFromResponse"`
	VirtualUsers               int64    `yaml:"virtualUsers"`
	Timeout                    Timeout  `yaml:"timeout"`
}

type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Timeout struct {
	Duration int32 `yaml:"duration"`
}
