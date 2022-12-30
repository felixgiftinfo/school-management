package handlers

import (
	"net/http"
	"school-management/services/dbcrud"
	"school-management/services/responses"

	"github.com/aws/aws-lambda-go/events"
)

func (entry *entryHandler) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := entry.dbproxy.GetConnection(entry.smg)
	if err != nil {
		return responses.GetErrorResponse(http.StatusInternalServerError, err)
	}
	crud := dbcrud.NewDbCrudEntry(db)

	switch request.Resource {
	case "/school":
		result, err := crud.Get()
		if err != nil {
			return responses.GetErrorResponse(http.StatusInternalServerError, err)
		}
		return responses.GetSuccessResponse(http.StatusInternalServerError, result)
	case "/school/{schoolId}":
		schoolId := request.PathParameters["schoolId"]
		result, err := crud.GetById(schoolId)
		if err != nil {
			return responses.GetErrorResponse(http.StatusInternalServerError, err)
		}
		return responses.GetSuccessResponse(http.StatusInternalServerError, result)
	default:
		return responses.GetErrorResponse(http.StatusInternalServerError, "Service not implemented")
	}

}
