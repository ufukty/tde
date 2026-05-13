package evolution

import (
	"tde/internal/evaluation"
	"tde/internal/evolution/models"
	"tde/internal/evolution/pool"
)

type commons struct {
	Evaluator *evaluation.Evaluator
	Params    *models.Parameters
	Context   *models.Context
	Pool      *pool.Pool
}
