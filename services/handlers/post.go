package handlers

import (
	"encoding/json"
	"net/http"
	"school-management/services/dbcrud"
	"school-management/services/models"
	"school-management/services/responses"

	"github.com/aws/aws-lambda-go/events"
)

func (entry *entryHandler) Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var payload models.School
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return responses.GetErrorResponse(http.StatusInternalServerError, err)
	}
	db, err := entry.dbproxy.GetConnection(entry.smg)
	if err != nil {
		return responses.GetErrorResponse(http.StatusInternalServerError, err)
	}
	crud := dbcrud.NewDbCrudEntry(db)
	result, err := crud.Create(payload)
	return responses.GetSuccessResponse(http.StatusInternalServerError, result)
}
