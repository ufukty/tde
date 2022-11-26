package evolution

import (
	codefragment "GoGP/evolve/code-fragment"
	"go/format"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Candidate struct {
	ID           uuid.UUID
	Program      *Program
	Fitness      float64
	ExecTimeInMs int
}

func NewCandidate() *Candidate {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create an UUID for new Individual"))
	}
	return &Candidate{
		ID: newUUID,
	}
}

func (c *Candidate) RandomInit() {
	*c.Program = codefragment.ProduceRandomFragment()
}

func (c *Candidate) CheckSyntax() bool {
	_, err := format.Source(*c.Program)
	return err == nil
}

func (c *Candidate) PickCrossoverPoint() int {
	length := len(*(c.Program))
	randomPoint := rand.Intn(length)
	return randomPoint
}

func (c *Candidate) Measure() {
	if !c.CheckSyntax() {
		c.Fitness = 1.1 // fitness for invalid-syntax programs exceeds the "1.0" treshold
		return
	}

	// t := &Testing{}
	// var timeStart = time.Now()
	// (*(i.TestFunction))(t)

	// i.Fitness = float64(t.TotalErrors) / float64(t.TotalCalls)
	// i.ExecTimeInMs = int(time.Since(timeStart))
}
