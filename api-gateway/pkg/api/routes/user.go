package routes

import (
	"api-gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler) {
	user := api.Group("/user")
	{
		user.GET("/getbyid/:userId", userHandler.GetUserData)
		user.POST("/getbylist", userHandler.GetUserDataList)
	}
}
