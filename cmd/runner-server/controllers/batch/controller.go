package batch

import (
	models "tde/models/transfer"

	"github.com/davecgh/go-spew/spew"
)

func Controller(request models.RunnerService_Batch_Request) (response models.RunnerService_Batch_Response) {
	spew.Println(request)
	response.Registered = true
	return
}
