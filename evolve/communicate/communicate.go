package communicate

import (
	"bufio"
	"fmt"
	"net/http"
)

const (
	ENDPOINT = "measure-fitness"
)

type FitnessMeasurementRequest struct {
	Programs
}

type Connection struct {
	AgentIP   string
	AgentPort string
}

func NewConnection() *Connection {
	return &Connection{
		AgentIP:   "127.0.0.1",
		AgentPort: "6000",
	}
}

func (c *Connection) Establish() bool {

	url := fmt.Sprintf("http://%s:%s/%s", c.AgentIP, c.AgentPort, ENDPOINT)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return true
}

func (c *Connection) Send(request string, callback func(response string)) {

}
