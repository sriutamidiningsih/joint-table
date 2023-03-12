package handler

import (
	response "dummy/pkg/api-response"
	userRepository "dummy/project/model"
	userService "dummy/project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	userService userService.Service
}

func NewUserHandler(userService userService.Service) *userHandler {
	return &userHandler{userService}
}

func (handler *userHandler) GetAll(ctx *gin.Context) {
	user, err := handler.userService.FindAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, response.DataNotFound())
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
		return
	}
	var usersResponse []userRepository.ResponseUser
	for _, b := range user {
		userResponse := ResponseUser(b)
		usersResponse = append(usersResponse, userResponse)
	}
	ctx.JSON(http.StatusOK, response.Success(usersResponse))
}

func ResponseUser(b userRepository.User) userRepository.ResponseUser {

	return userRepository.ResponseUser{
		ID:   b.ID,
		Name: b.Name,
		NoHp: b.NoHp,
	}
}
