package runner_communicator

import (
	"go/ast"
	"tde/internal/utilities"
	models "tde/models/program"
	transfer_models "tde/models/transfer"
)

type Batch struct {
	File       *ast.File
	Candidates []*models.Candidate
}

func (batch *Batch) Divide(noBatches int) (batches []*Batch) {
	for _, bucket := range utilities.DivideIntoBuckets(batch.Candidates, noBatches) {
		batches = append(batches, &Batch{
			File:       batch.File,
			Candidates: bucket,
		})
	}
	return
}

func (batch *Batch) GetRequestDTO() *transfer_models.RunnerService_Batch_Request {
	dto := transfer_models.RunnerService_Batch_Request{
		Candidates:   []transfer_models.Candidate{},
		ArchiveID:    "", // FIXME:
		FileTemplate: batch.File,
	}

	for _, candidate := range batch.Candidates {
		dto.Candidates = append(dto.Candidates, transfer_models.Candidate{
			CandidateID: string(candidate.UUID),
			FuncDecl:    candidate.AST.FuncDecl,
		})
	}

	return &dto
}
