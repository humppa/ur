/* main.go */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// RequestBody is the schema for the body of the client request
type RequestBody struct {
	Name string `json:"name"`
}

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
}

func makeErr(msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, errors.New(msg)
}

// Ur is the main program logic
func Ur(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body RequestBody
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return makeErr(fmt.Sprintf("Failed to parse request body: %+v", req.Body))
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("The name is %s", body.Name),
	}, nil
}

func main() {
	log.Printf("The Pleb Game of Ur")
	lambda.Start(Ur)
}
