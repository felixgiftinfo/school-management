package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"school-management/services/handlers"
	"school-management/services/responses"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/rs/zerolog"
)

var sess *session.Session
var smg *secretsmanager.SecretsManager

func init() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}
func main() {
	log.Println("Hello Felix")

	smg = secretsmanager.New(sess, aws.NewConfig().WithRegion(os.Getenv("REGION")))

	//svc := dynamodb.New(sess)
	lambda.Start(handler)
}

//	GOOS=linux
//
// go build main.go
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("AWS request")
	var response events.APIGatewayProxyResponse
	var err error

	mockLogBuffer := &bytes.Buffer{}
	logger := zerolog.New(mockLogBuffer)

	handlerEntryPoint := handlers.NewHandlerEntry(smg, nil, logger)

	switch request.HTTPMethod {
	case "POST":
		response, err = handlerEntryPoint.Get(request)
	case "PUT":
	case "GET":
	case "DELETE":

	default:
		response, err = responses.GetErrorResponse(http.StatusInternalServerError, "Service not implemented")

	}
	return response, err
}
