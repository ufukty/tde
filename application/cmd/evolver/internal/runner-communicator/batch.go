package runner_communicator

import (
	"go/ast"
	"tde/cmd/runner/endpoints/batch/batch_post"
	"tde/internal/utilities"
	models "tde/models/program"
)

type Batch struct {
	File     *ast.File
	Subjects []*models.Subject
}

func (batch *Batch) Divide(noBatches int) (batches []*Batch) {
	for _, bucket := range utilities.DivideIntoBuckets(batch.Subjects, noBatches) {
		batches = append(batches, &Batch{
			File:     batch.File,
			Subjects: bucket,
		})
	}
	return
}

func (batch *Batch) GetRequestDTO() *batch_post.Request {
	var req = batch_post.Request{
		Subjects:     []batch_post.Subject{},
		ArchiveID:    "", // FIXME:
		FileTemplate: batch.File,
	}

	for _, subject := range batch.Subjects {
		req.Subjects = append(req.Subjects, batch_post.Subject{
			Sid:      string(subject.Sid),
			FuncDecl: subject.AST,
		})
	}

	return &req
}
