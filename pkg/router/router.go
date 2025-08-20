package router

import (
	"github.com/diiineeei/mvc-sample/handlers"
	"github.com/diiineeei/mvc-sample/pkg/middleware"
	"github.com/diiineeei/mvc-sample/providers"
	"github.com/diiineeei/mvc-sample/repo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware(logger))

	userRepo := repo.NewUserRepo()
	userProvider := providers.NewUserProvider(userRepo)
	userHandler := handlers.NewUserHandler(userProvider)

	r.GET("/webhook", userHandler.Webhook)
	r.POST("/webhook", userHandler.Webhook)

	return r
}
