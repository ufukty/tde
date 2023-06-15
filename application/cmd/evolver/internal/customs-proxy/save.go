package customs_proxy

import (
	"bytes"
	"encoding/json"
	"go/ast"

	"github.com/pkg/errors"
)

func encodeAsJson(fd *ast.FuncDecl) ([]byte, error) {
	var buf = bytes.NewBuffer([]byte{})
	var err = json.NewEncoder(buf).Encode(fd)
	if err != nil {
		return nil, errors.Wrap(err, "json encoding")
	}
	return buf.Bytes(), nil
}

func decodeFromJson(src []byte) (*ast.FuncDecl, error) {
	var fd = new(ast.FuncDecl)
	var reader = bytes.NewReader(src)
	var err = json.NewDecoder(reader).Decode(fd)
	if err != nil {
		return nil, errors.Wrap(err, "json decoding")
	}
	return fd, nil
}

// func Save(batch map[models.CandidateID]models.Candidate) {

// }
