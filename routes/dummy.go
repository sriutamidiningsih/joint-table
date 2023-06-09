package routes

import (
	handler "dummy/project/handler"
	repo "dummy/project/repository"
	service "dummy/project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, router *gin.Engine) {
	userRepository := repo.NewRepositoryUser(db)
	userService := service.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService)

	v1 := router.Group("/v1/user")

	v1.GET("/", userHandler.GetAll)
	v1.GET("/joins", userHandler.GetOrders)
	v1.GET("/:id_user", userHandler.GetByUserId)
	v1.POST("/", userHandler.Create)
	v1.POST("/order", userHandler.CreateOrders)

}
