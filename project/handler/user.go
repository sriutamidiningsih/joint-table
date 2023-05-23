package handler

import (
	response "dummy/pkg/api-response"
	Models "dummy/project/model"
	userService "dummy/project/services"
	helper "dummy/utility"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userHandler struct {
	userService userService.Service
}

func NewUserHandler(userService userService.Service) *userHandler {
	return &userHandler{userService}
}

func (handler *userHandler) GetAll(ctx *gin.Context) {
	user, err := handler.userService.FindUserAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, response.DataNotFound())
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
		return
	}
	var usersResponse []Models.ResponseUser
	for _, b := range user {
		userResponse := ResponseUser(b)
		usersResponse = append(usersResponse, userResponse)
	}
	ctx.JSON(http.StatusOK, response.Success(usersResponse))
}

func (handler *userHandler) GetOrders(ctx *gin.Context) {
	users, err := handler.userService.FindJoins()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, response.DataNotFound())
			return
		}
	}
	var responseData []gin.H
	for _, user := range users {
		responseData = append(responseData, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"orderlist": func() []gin.H {
				var orderlist []gin.H
				for _, orderlists := range user.OrderList {
					orderlist = append(orderlist, gin.H{
						"id_user":      orderlists.IdUser,
						"name_product": orderlists.NameProduct,
						"total_order":  orderlists.TotalOrder,
					})
				}
				return orderlist
			}(),
		})
	}

	ctx.JSON(http.StatusOK, response.Success(responseData))

}

// func (handler *userHandler)GetOrders(ctx *gin.Context)  {
// 	orders, err := handler.userService.FindOrders()
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			ctx.JSON(http.StatusNotFound, response.DataNotFound())
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
// 		return
// 	}

// }

func (handler *userHandler) GetByUserId(ctx *gin.Context) {
	idString := ctx.Param("id_user")
	id, _ := strconv.Atoi(idString)

	if id == 0 {
		ctx.JSON(http.StatusNotFound, response.NotFound("ID User"))
		return
	}

	users, err := handler.userService.FindByUserId(id)

	if users.ID == 0 {
		ctx.JSON(http.StatusNotFound, response.NotFound("ID User"))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(users))

}

func (handler *userHandler) Create(ctx *gin.Context) {
	var joinRequest Models.RequestOrders

	err := ctx.ShouldBindJSON(&joinRequest)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := helper.ErrorMessages(err)
			ctx.JSON(http.StatusBadRequest, response.BadRequest(errorMessages))
			return
		}
		ctx.JSON(http.StatusBadRequest, response.BadRequest(err.Error()))
		return
	}

	joins, err := handler.userService.Create(joinRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(joins))
}

func (handler *userHandler) CreateOrders(ctx *gin.Context) {
	var orders Models.RequestOrder

	err := ctx.ShouldBindJSON(&orders)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := helper.ErrorMessages(err)
			ctx.JSON(http.StatusBadRequest, response.BadRequest(errorMessages))
			return
		}
		ctx.JSON(http.StatusBadRequest, response.BadRequest(err.Error()))
		return
	}

	order, err := handler.userService.CreateOrders(orders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(order))
}

func ResponseUser(b Models.Users) Models.ResponseUser {

	return Models.ResponseUser{
		ID:    b.ID,
		Name:  b.Name,
		Email: b.Email,
	}
}
