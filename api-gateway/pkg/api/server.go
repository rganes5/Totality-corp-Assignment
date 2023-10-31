package api

import (
	_ "api-gateway/docs"
	"api-gateway/pkg/api/routes"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-gateway/pkg/api/handlers"
	"api-gateway/pkg/config"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(cfg *config.Config, userHandler handlers.UserHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.RegisterUserRoutes(engine.Group("/"), userHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"statuscode": 404,
			"message":    "invalid url",
		})
	})
	return &Server{
		Engine: engine,
		Port:   cfg.Port,
	}, nil
}

func (cfg *Server) Start() {
	cfg.Engine.Run(cfg.Port)
}
