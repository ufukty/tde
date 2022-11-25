package evolution

import (
	codefragment "GoGP/evolve/code-fragment"
	"go/format"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Individual struct {
	ID           uuid.UUID
	Program      *Program
	Fitness      float64
	ExecTimeInMs int
}

func NewIndividual() *Individual {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create an UUID for new Individual"))
	}
	return &Individual{
		ID: newUUID,
	}
}

func (i *Individual) RandomInit() {
	*i.Program = codefragment.ProduceRandomFragment()
}

func (i *Individual) CheckSyntax() bool {
	_, err := format.Source(*i.Program)
	return err == nil
}

func (i *Individual) PickCrossoverPoint() int {
	length := len(*(i.Program))
	randomPoint := rand.Intn(length)
	return randomPoint
}

func (i *Individual) Measure() {
	if !i.CheckSyntax() {
		i.Fitness = 1.1 // fitness for invalid-syntax programs exceeds the "1.0" treshold
		return
	}

	// t := &Testing{}
	// var timeStart = time.Now()
	// (*(i.TestFunction))(t)

	// i.Fitness = float64(t.TotalErrors) / float64(t.TotalCalls)
	// i.ExecTimeInMs = int(time.Since(timeStart))
}
