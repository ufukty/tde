package services

type ServiceName string

const (
	ApiGateway = ServiceName("api_gateway")
	Customs    = ServiceName("customs")
	Evolver    = ServiceName("evolver")
	Runner     = ServiceName("runner")
)
