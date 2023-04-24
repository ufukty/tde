package evolution

import (
	"fmt"
	"net/http"
	"tde/models/dto"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	reqDTO := dto.EvolverService_Evolve_Request{}
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
