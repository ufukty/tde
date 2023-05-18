package batch

import (
	models "tde/models/dto"

	"github.com/davecgh/go-spew/spew"
)

func Controller(request models.RunnerService_Batch_Request) (response models.RunnerService_Batch_Response) {
	spew.Println(request)
	response.Registered = true
	return
}
