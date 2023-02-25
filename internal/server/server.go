package server

import (
	// "tde/internal/embedding"
	"tde/models/network_models"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"go/parser"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Server struct {
	Address string
}

func NewServer(port int) *Server {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	s := &Server{address}
	s.Start()
	return s
}

func (s *Server) Start() {
	http.HandleFunc("/", s.Controller)
	http.ListenAndServe(s.Address, nil)
}

func (s *Server) ParseRequest(r *http.Request) (*network_models.FitnessMeasurementRequest, error) {
	var request = network_models.FitnessMeasurementRequest{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return nil, errors.Wrap(err, "Could not use JSON Decoder to parse the request body")
	}

	return &request, nil
}

func (s *Server) Controller(w http.ResponseWriter, rawRequest *http.Request) {

	// marshall request (uuid-code fragments)
	request, err := s.ParseRequest(rawRequest)
	if err != nil {
		fmt.Fprintln(w, "Bad request")
		log.Println(err)
		return
	}

	var decodedPrograms = map[string]string{} // UUID: Body

	// embed candidates into the target file
	for _, candidate := range request.Candidates {
		// program.UUID

		data, err := base64.StdEncoding.DecodeString(candidate.BodyBase64Encoded)
		if err != nil {
			fmt.Fprintln(w, "Bad request")
			log.Fatalln(errors.Wrap(err, "Could not decode function body sent by client"))
		}
		decodedPrograms[candidate.UUID] = string(data)

		_, err = parser.ParseExpr(decodedPrograms[candidate.UUID])
		if err != nil {
			fmt.Fprintf(w, "Bad request, Body for one candidate is not valid Golang code:\n\nUUID: %s\nBody (decoded):\n%s", candidate.UUID, decodedPrograms[candidate.UUID])
			log.Println(errors.Wrap(err, "Body is not valid Golang code"))
		}
	}

	// f := embedding.NewEmbeddingConfig("main_tde.go")

	// f := file.NewFile(request.File)

	// run test `go test TestFoo`
}
