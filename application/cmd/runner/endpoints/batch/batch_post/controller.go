package batch_post

import (
	"github.com/davecgh/go-spew/spew"
)

func Controller(request Request) (response Response) {
	spew.Println(request)
	response.Registered = true
	return
}
