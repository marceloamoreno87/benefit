// Routes/Routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/benefit/api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp := r.Group("/api")
	{
		grp.GET("/benefit", controllers.GetBenfitByDoc)
	}

	return r
}
