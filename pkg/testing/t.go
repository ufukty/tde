package testing

import (
	"fmt"
	"tde/internal/evolution/models"

	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Config struct {
	// in kiloflops
	MaxCompute int
	// in kilobytes
	MaxMemory int
	// Maximum program size; in bytes
	MaxSize int
	// Maximum execution time in seconds
	MaxTime int
	// Suggestive value for optimization
	ComputeToMemoryRatio float64
}

type TargetFunctionType any

type T struct {
	UUID                   models.Sid    `json:"uuid"`
	AssertionResults       []bool        `json:"assertion_results"`
	AssertionErrorDistance []float64     `json:"assertion_error_distances"`
	ExecTime               time.Duration `json:"exec_time"`
	Panicked               bool          `json:"panicked"`
	Config                 Config        `json:"-"`
}

func NewT(canditateUUID models.Sid) *T {
	e := T{
		UUID: canditateUUID,
	}
	return &e
}

func (t *T) SetConfig(config Config) {
	t.Config = config
}

func (t *T) Export() {
	f, err := os.Create("results.json")
	if err != nil {
		panic(errors.Wrap(err, "Could not create 'result.json' for writing"))
	}
	json.NewEncoder(f).Encode(t)
}

func (t *T) LoadResults(path string) error {
	fh, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening results.json: %w", err)
	}
	err = json.NewDecoder(fh).Decode(t)
	if err != nil {
		return fmt.Errorf("decoding results.json: %w", err)
	}
	return nil
}
