//go:build evolve
// +build evolve

package testing

import "testing"

type Input struct {
	x int
	y int
}

type Output struct {
	a int
	b int
}

func TestedFunction(x, y int) (int, int) {
	return -1, -1
}

func Test_SimpleEvolution(e *evolve) {
	cases := map[*Input]*Output{
		{x: 1, y: 1}: {a: 2, b: 1},
		{x: 2, y: 1}: {a: 5, b: 1},
		{x: 3, y: 1}: {a: 10, b: 1},
		{x: 1, y: 2}: {a: 3, b: 2},
		{x: 2, y: 2}: {a: 6, b: 2},
		{x: 3, y: 2}: {a: 11, b: 2},
	}

	for input, expectedOutput := range cases {
		objective.NewRun(func(run *Run) {

			a, b := TestedFunction(input.x, input.y)
			run.Assert(a, expectedOutput.a)
			run.Assert(b, expectedOutput.b)
		})
	}
}

func Test_MultiStageTesting(t *evolve) {
	testing := CreateTesting()

	testing.NewObjective("basic", false, func(objective *Objective) {
		cases := map[*Input]*Output{
			{x: 1, y: 1}: {a: 2, b: 1},
			{x: 2, y: 1}: {a: 5, b: 1},
			{x: 3, y: 1}: {a: 10, b: 1},
			{x: 1, y: 2}: {a: 3, b: 2},
			{x: 2, y: 2}: {a: 6, b: 2},
			{x: 3, y: 2}: {a: 11, b: 2},
		}

		for input, expectedOutput := range cases {
			objective.NewRun(func(run *Run) {

				a, b := TestedFunction(input.x, input.y)
				run.Assert(a, expectedOutput.a)
				run.Assert(b, expectedOutput.b)
			})
		}
	})

	testing.NewObjective("intermediate", false, func(objective *Objective) {

	})

	testing.NewObjective("advanced", false, func(objective *Objective) {

	})
}
