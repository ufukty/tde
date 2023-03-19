package in_program_models

type RunnerType int

const (
	RUNNER_CLOUD = RunnerType(iota)
	RUNNER_ON_PREM
)

type EvolverParameters struct {
	GenerationsToIterate int
	RunnerType           RunnerType
}

type OnPremisesRunnerConfig struct {
	Address string
	Port    string
	Token   string
}
