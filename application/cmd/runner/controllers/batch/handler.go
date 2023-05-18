package batch

import (
	"fmt"
	"net/http"
	models "tde/models/dto"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	reqDTO := models.RunnerService_Batch_Request{}
	err := reqDTO.ParseRequest(r)
	if err != nil {
		fmt.Fprintln(w, "Error")
		return
	}

	resDTO := Controller(reqDTO)
	err = resDTO.SerializeIntoResponseWriter(w)
	if err != nil {
		fmt.Fprintln(w, "Error")
		return
	}

}
