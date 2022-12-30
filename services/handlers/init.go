package handlers

import (
	"school-management/services/auths"

	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/rs/zerolog"
)

type entryHandler struct {
	smg     secretsmanageriface.SecretsManagerAPI
	dbproxy auths.DbConnectable
	//Claims  auths.AwsJwtClaims
	Logger zerolog.Logger
}

func NewHandlerEntry(sMng secretsmanageriface.SecretsManagerAPI, dbproxy auths.DbConnectable, log zerolog.Logger) *entryHandler {

	return &entryHandler{sMng, dbproxy, log}
}
