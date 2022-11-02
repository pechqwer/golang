package route

import (
	controller "servers/restAPI/controllers"
	auth "servers/restAPI/middleware"

	"github.com/gin-gonic/gin"
)

func V1(app *gin.Engine) {
	v1 := app.Group("api/v1")

	v1.POST("/line", controller.Webhook())

	v1.Use(auth.AuthRequired())
	{
		v1.POST("/200", controller.ReadStatus200())
		v1.GET("/200", controller.ReadStatus200())
	}

	// v1.PUT("/read", auth.AuthRequired(), controller.ReadEndpoint())
}
