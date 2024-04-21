package bucket

import (
	"tde/internal/microservices/errors/loggable"
	"tde/internal/microservices/errors/tagged"

	"bytes"
	"encoding/json"
	"strings"
)

type Bucket []loggable.Loggable

func New() *Bucket {
	return &Bucket{}
}

func (bucket *Bucket) Add(err loggable.Loggable) {
	*bucket = append(*bucket, err)
}

func (bucket *Bucket) Error() string {
	var arr = []string{}
	for _, err := range *bucket {
		arr = append(arr, err.Error())
	}
	return strings.Join(arr, ", in addition to: ")
}

func (bucket *Bucket) Log() string {
	var buf = bytes.NewBuffer([]byte{})
	var err = json.NewEncoder(buf).Encode(bucket)
	if err != nil {
		panic("failed on encoding error bucket into json")
	}
	return buf.String()
}

func (bucket *Bucket) Tag(e loggable.Loggable) *tagged.TaggedError {
	return tagged.New(e, bucket)
}

func (bucket *Bucket) IsAny() bool {
	return len(*bucket) > 0
}
