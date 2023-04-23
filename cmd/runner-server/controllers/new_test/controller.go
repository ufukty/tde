package new_test

import (
	"fmt"
	models "tde/models/transfer/artifacts/services/runner/models"
)

func Controller(request models.NewTest_Request) (response models.NewTest_Response) {
	fmt.Println(request.ArchiveID)
	return
}
