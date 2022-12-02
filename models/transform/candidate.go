package transform

import (
	"models/in_program_models"
	"models/network_models"

	"encoding/base64"

	"github.com/pkg/errors"
)

func Candidate_ProgramToNetwork(c *in_program_models.Candidate) *network_models.Candidate {
	bodyBase64Encoded := base64.StdEncoding.EncodeToString(c.Body)

	return &network_models.Candidate{
		UUID:              string(c.UUID),
		BodyBase64Encoded: bodyBase64Encoded,
	}
}

func Candidate_NetworkToProgram(c *network_models.Candidate) (*in_program_models.Candidate, error) {
	bodyDecodedBytes, err := base64.StdEncoding.DecodeString(c.BodyBase64Encoded)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode function body sent by client")
	}
	// bodyDecoded := string(bodyDecodedBytes)

	return &in_program_models.Candidate{
		UUID:         in_program_models.CandidateID(c.UUID),
		Body:         bodyDecodedBytes,
		Fitness:      0,
		ExecTimeInMs: 0,
	}, nil
}
