package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go/parser"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Server struct{}

type RequestDataType_Program struct {
	UUID              string `json:"uuid"`
	BodyBase64Encoded string `json:"body"`
	Body              string
}

type Request struct {
	Programs []RequestDataType_Program `json:"programs"`
}

func NewServer() *Server {
	s := &Server{}
	s.Start()
	return s
}

func (s *Server) Start() {
	http.HandleFunc("/", s.Controller)
	http.ListenAndServe("127.0.0.1:6000", nil)
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
	for _, program := range request.Programs {
		// program.UUID

		data, err := base64.StdEncoding.DecodeString(program.BodyBase64Encoded)
		if err != nil {
			fmt.Fprintln(w, "Bad request")
			log.Fatalln(errors.Wrap(err, "Could not decode function body sent by client"))
		}
		decodedPrograms[program.UUID] = string(data)

		_, err = parser.ParseExpr(decodedPrograms[program.UUID])
		if err != nil {
			fmt.Fprintf(w, "Bad request, Body for one candidate is not valid Golang code:\n\nUUID: %s\nBody (decoded):\n%s", program.UUID, decodedPrograms[program.UUID])
			log.Println(errors.Wrap(err, "Body is not valid Golang code"))
		}
	}

	// run test `go test TestFoo`
}

func (s *Server) ParseRequest(r *http.Request) (*Request, error) {
	var request = Request{}

	fmt.Printf("%#v\n", r.Header)
	fmt.Printf("%#v\n", r.Body)

	// b, err := io.ReadAll(r.Body)
	// fmt.Println(string(b))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return nil, errors.Wrap(err, "Could not use JSON Decoder to parse the request body")
	}

	return &request, nil
}
