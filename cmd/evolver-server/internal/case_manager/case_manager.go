package case_manager

import (
	"errors"
	"tde/internal/evolution"
	"tde/models/common_models"

	"github.com/google/uuid"
)

type CaseID string

func NewCaseID() CaseID {
	return CaseID(uuid.NewString())
}

type Case struct {
	EvolutionManager *evolution.EvolutionManager
	EvolutionConfig  *common_models.EvolutionConfig
}

type CaseManager struct {
	cases map[CaseID]*Case
}

func NewCaseManager() *CaseManager {
	return &CaseManager{}
}

func (cm *CaseManager) NewCase(cs *Case) (caseID CaseID) {
	caseID = NewCaseID()
	cm.cases[caseID] = cs
	return
}

func (cm *CaseManager) Iterate(caseID CaseID) error {
	var (
		cs *Case
		ok bool
	)
	if cs, ok = cm.cases[caseID]; !ok {
		return errors.New("")
	}

	cs.EvolutionManager.InitPopulation(cs.EvolutionConfig.Population)
	// TODO: make request to runner
	cs.EvolutionManager.IterateLoop()

	return nil
}
