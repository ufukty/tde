package batch_post

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	reqDTO := Request{}
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
