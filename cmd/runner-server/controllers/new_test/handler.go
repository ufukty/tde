package new_test

import (
	"fmt"
	"net/http"
	models "tde/models/transfer/artifacts/services/runner/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	reqDTO := models.NewTest_Request{}
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
