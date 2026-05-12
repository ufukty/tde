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

## Next Milestone: rest api boilerplate generator

### Steps

1. List `*.go` files
1. list `GenDecl` that has `Tok: token.TYPE` (type declaration)
1. each `GenDecl` should end with `*Service` named interface declaration (service declaration)
1. each members of service interface should have 1 input 1 output parameter each type should be defined in same `GenDecl`. Line comment should be included in format: `// POST /user`
1. input parameter's type should be named as `<ServiceName>_<EndpointName>_Request` && output parameter's type should be named as `<ServiceName>_<EndpointName>_Response`

```go
type {
    UserService interface {
        Create(UserService_Create_Request) UserService_Create_Response            // POST /user
        DeletionProtocole(UserService_Delete_Request) UserService_Delete_Response // POST /deletion
    }

    UserService_Create_Request struct {
        Name        string
        Email       string
        Password    string
        RedirectURI string
        CSRF        string
    }

    UserService_Create_Response struct {
        RedirectURI string
    }

    UserService_Deletion_Request struct {
        CSRF string
        Step int
    }

    UserService_Deletion_Response struct {

    }
}
```

### Behaviours

1. `serdeser` doesn't rewrite previously generated and user-modified file
1. If a rewrite is seen as improvement, user will be presented with choice in CLI
1. DSD
