package mutation

import (
	"fmt"
	"tde/internal/evolution/models"
)

// TODO: RegenerateSubtree (cfg/node_constructor)
// TODO: Merge declared variables

type MutationClass int

const (
	Create = MutationClass(iota)
	Delete
	Update
)

type MutationSurface int

const (
	Statement = MutationSurface(iota)
	Expression
	Decleration
	Specification
)

type Probabilities struct {
	BlockStmt struct {
		DeclareVariable     float64
		ConditionalizeRange float64
		RangeStdDeviation   float64
	}
	Statement struct {
		IfStmt struct{ RegenerateCond, RegenerateBody, RegenerateElse, AddElif, RemoveElif float64 }
	}
	Expression struct {
	}
	Decleration struct {
	}
}

// LiteralShuffle
// TokenShuffle
// Conditionalize
// GenerateLine
// CreateVariableDeclaration
// SwapLines
// CallPackageFunction
// GenerateExpression
var defaults = Probabilities{
	BlockStmt: struct {
		DeclareVariable     float64
		ConditionalizeRange float64
		RangeStdDeviation   float64
	}{
		DeclareVariable:     0,
		ConditionalizeRange: 0,
		RangeStdDeviation:   0,
	},
	Statement: struct {
		IfStmt struct {
			RegenerateCond float64
			RegenerateBody float64
			RegenerateElse float64
			AddElif        float64
			RemoveElif     float64
		}
	}{
		IfStmt: struct {
			RegenerateCond float64
			RegenerateBody float64
			RegenerateElse float64
			AddElif        float64
			RemoveElif     float64
		}{
			RegenerateCond: 0,
			RegenerateBody: 0,
			RegenerateElse: 0,
			AddElif:        0,
			RemoveElif:     0,
		},
	},
	Expression:  struct{}{},
	Decleration: struct{}{},
}

type Mutator func(models.Parameters, models.Context, models.Subject) error

var mutations = map[Mutation]Mutator{
	MLiteralShuffle:            literalShuffle,
	MTokenShuffle:              tokenShuffle,
}

var ErrUnavailable = fmt.Errorf("No available point for mutation")
