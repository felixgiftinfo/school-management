package handlers

import (
	"encoding/json"
	"net/http"
	"school-management/services/dbcrud"
	"school-management/services/models"
	"school-management/services/responses"

	"github.com/aws/aws-lambda-go/events"
)

func (entry *entryHandler) Put(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := entry.dbproxy.GetConnection(entry.smg)
	if err != nil {
		return responses.GetErrorResponse(http.StatusInternalServerError, err)
	}
	crud := dbcrud.NewDbCrudEntry(db)
	var payload models.School
	err = json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return responses.GetErrorResponse(http.StatusInternalServerError, err)
	}
	switch request.Resource {
	case "/school/{schoolId}":
		schoolId := request.PathParameters["schoolId"]
		result, err := crud.Update(schoolId, payload)
		if err != nil {
			return responses.GetErrorResponse(http.StatusInternalServerError, err)
		}
		return responses.GetSuccessResponse(http.StatusInternalServerError, result)
	default:
		return responses.GetErrorResponse(http.StatusInternalServerError, "Service not implemented")
	}
}
