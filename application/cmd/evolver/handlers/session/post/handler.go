package session_post

import (
	"fmt"
	"net/http"
)

// var (
// 	ErrTargetInaccessible = errors.New("Target service is unaccassable or can not accept the request at the moment")
// )

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
		// if errors.Is(err, ErrTargetInaccessible) {
		// 	queueRequest(reqDTO)
		// }
		return
	}

}
