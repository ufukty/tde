package in_program_models

import (
	"tde/internal/cfg/character"

	"go/format"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CandidateID string

type Candidate struct {
	UUID         CandidateID
	Body         []byte
	Fitness      float64
	ExecTimeInMs int
}

func NewCandidate() *Candidate {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not create an UUID for new Individual"))
	}
	return &Candidate{
		UUID: CandidateID(newUUID.String()),
	}
}

func (c *Candidate) RandomInit() {
	c.Body = character.ProduceRandomFragment()
}

func (c *Candidate) CheckSyntax() bool {
	_, err := format.Source(c.Body)
	return err == nil
}

func (c *Candidate) PickCrossoverPoint() int {
	length := len(c.Body)
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
