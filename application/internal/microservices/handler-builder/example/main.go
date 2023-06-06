package endpoint

import (
	"mime/multipart"
)

//go:generate boilerer -b ../builder.go main.go
type (
	RequestHeader struct {
		Authorization string `header:"Authorization"`
	}

	RequestBody struct {
		Checksum string         `form:"value"`
		File     multipart.File `form:"file" header:"filename"`
	}

	Request struct {
		Header RequestHeader
		Body   RequestBody `Content-Type:"multipart/form-data"`
	}
)

type (
	ResponseHeader struct {
		StatusCode int `header:"status"`
	}

	Response struct {
		Header ResponseHeader
		Body   []byte `Content-Type:"application/zip"`
	}
)

type (
	Failure struct {
	}
)
