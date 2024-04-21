package mutation

import (
	"tde/internal/evolution/models"
)

type MutationManager struct {
	params  *models.Parameters
	context *models.Context

	availables    []Mutation
	probabilities []float64 // cumulative
}

func NewMutationManager(params *models.Parameters, context *models.Context) *MutationManager {
	mngr := &MutationManager{
		params:        params,
		context:       context,
		availables:    []Mutation{},
		probabilities: []float64{},
	}

	if len(params.Packages) > 0 {
		mngr.availables = append(mngr.availables, MExprCallPackageFunction)
	}

	// baseprobs := []float64{}
	// for _, op := range mngr.availables {

	// }

	return mngr
}

func (mngr *MutationManager) Mutate(subj *models.Subject) error {
	return nil
}
