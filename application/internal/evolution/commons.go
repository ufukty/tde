package evolution

import (
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/pool"
	models "tde/models/program"
)

type commons struct {
	Evaluator *evaluation.Evaluator
	Params    *models.Parameters
	Context   *models.Context
	Pool      *pool.Pool
}
