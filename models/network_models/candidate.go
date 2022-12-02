package network_models

type Candidate struct {
	UUID              string `json:"uuid"`
	BodyBase64Encoded string `json:"bodyBase64Encoded"`
}

type FitnessMeasurementRequest struct {
	File       string      `json:"file"`
	TestFile   string      `json:"testfile"`
	Candidates []Candidate `json:"candidates"`
}
