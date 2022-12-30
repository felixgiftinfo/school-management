package responses

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type BodyResponse map[string]interface{}

func GetSuccessResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	return rawResponse(status, body, false)
}
func GetErrorResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	return rawResponse(status, body, true)
}

func rawResponse(status int, body interface{}, isError bool) (events.APIGatewayProxyResponse, error) {
	responseBody := BodyResponse{}
	if isError {
		responseBody["msg"] = body
	} else {
		responseBody["data"] = body
	}
	_byte, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}
	stdHeaders := map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
	response := events.APIGatewayProxyResponse{StatusCode: status, Body: string(_byte), Headers: stdHeaders}

	return response, nil
}
