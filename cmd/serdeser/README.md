# serdeser - Serializer/Deserializer

## Install

```sh
go install tde/cmd/serdeser
```

## Usage

```go
//go:generate serdeser filename.go

type MyEndpoint1_Request struct {}

type MyEndpoint1_Response struct {}

type MyEndpoint2_Request struct {}

type MyEndpoint2_Response struct {}
```

It will create a file named `filename.sd.go` which will contain:

```go
func (s *MyEndpoint1_Request) Serialize(w http.ResponseWriter) { ... }
func (s *MyEndpoint1_Request) Deserialize(r *http.Request) { ... }

func (s *MyEndpoint1_Response) Serialize(w http.ResponseWriter) { ... }
func (s *MyEndpoint1_Response) Deserialize(r *http.Request) { ... }

func (s *MyEndpoint2_Request) Serialize(w http.ResponseWriter) { ... }
func (s *MyEndpoint2_Request) Deserialize(r *http.Request) { ... }

func (s *MyEndpoint2_Response) Serialize(w http.ResponseWriter) { ... }
func (s *MyEndpoint2_Response) Deserialize(r *http.Request) { ... }
```
