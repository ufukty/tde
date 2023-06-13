package runner_communicator

import (
	"go/ast"
	"tde/cmd/runner/endpoints/batch/batch_post"
	"tde/internal/utilities"
	models "tde/models/program"
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

func (batch *Batch) GetRequestDTO() *batch_post.Request {
	var req = batch_post.Request{
		Candidates:   []batch_post.Candidate{},
		ArchiveID:    "", // FIXME:
		FileTemplate: batch.File,
	}

	for _, candidate := range batch.Candidates {
		req.Candidates = append(req.Candidates, batch_post.Candidate{
			CandidateID: string(candidate.UUID),
			FuncDecl:    candidate.AST.FuncDecl,
		})
	}

	return &req
}
