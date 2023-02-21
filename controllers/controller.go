package controllers


import (

	"net/http"
	"github.com/gin-gonic/gin"

	model "backend_community_rest/models"
	response "backend_community_rest/responses"
	exception "backend_community_rest/exceptions"
	repository "backend_community_rest/repositories"

)



type response_data struct {

	Message string `json:"message"`

}


// add join request
func AddJoinRequest(c *gin.Context) {

	// declare payload struct
	var user_request model.JoinRequest


	// default value
	http_response, status, message := response.AddJoinResponse(
		http.StatusOK,
	)


	// bind json payload
	err := c.BindJSON(&user_request)
	if err != nil {

		http_response, status, message = response.AddJoinResponse(
			http.StatusBadRequest,
		)

	}

	exception.TryCatchError(err)


	// call repository
	if !repository.AddJoinRequestRepo(user_request) {

		http_response, status, message = response.AddJoinResponse(
			http.StatusUnauthorized,
		)

	}


	response := model.JoinResponse {

		HttpResponse: http_response,
		Status: status,
		Data: response_data{

			Message: message,

		},
	}

	c.JSON(http_response, response)

	
}
