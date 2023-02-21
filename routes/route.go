package routes

import (
	
	"github.com/gin-gonic/gin"

	controller "backend_community_rest/controllers"

)


func RoutesGroup(r *gin.RouterGroup) {

	r.POST("/join", controller.AddJoinRequest)

}