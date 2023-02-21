package api
 
import (

	"net/http"
	"github.com/gin-gonic/gin"

	route "backend_community_rest/routes"
	middleware "backend_community_rest/middlewares"

)




var(
	
	Application *gin.Engine

)


func init() { 

	// initialize gin
    Application = gin.New()

	// use middleware
	Application.Use(middleware.CorsMiddleware())

	// routes grouping
	group := Application.Group("/api")
	route.RoutesGroup(group)

	// prevent user if access not found route
	Application.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
				"code": "PAGE_NOT_FOUND",
				"message": "Page not found",
		})
	})

	
}


/*
 *  handler function
 */
func Handler(responses http.ResponseWriter , requests *http.Request) {
	
	Application.ServeHTTP(responses, requests)

}