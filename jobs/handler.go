package jobs

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	routes := gin.Default()

	routes.GET("/jobrunner/json", JobJson)
	routes.LoadHTMLGlob("./templates/status.html")

	routes.GET("/jobrunner/html", JobHtml)

	return routes
}

func JobJson(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "status.html", jobrunner.StatusPage())
}