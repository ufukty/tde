package runner_communicator

import (
	"tde/internal/utilities"
	models "tde/models/program"
)

type Batch struct {
	FileTemplate string
	Candidates   []*models.Candidate
}

func (batch *Batch) Divide(noBatches int) (batches []*Batch) {
	for _, bucket := range utilities.DivideIntoBuckets(batch.Candidates, noBatches) {
		batches = append(batches, &Batch{
			FileTemplate: batch.FileTemplate,
			Candidates:   bucket,
		})
	}
	return
}
