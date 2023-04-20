package driveradapt

import (
	"errors"
	"net/http"

	"github.com/dany0814/go-apisolutions/internal/core/application"
	"github.com/dany0814/go-apisolutions/internal/core/application/dto"
	"github.com/dany0814/go-apisolutions/internal/core/domain"
	"github.com/dany0814/go-apisolutions/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService application.UserService
}

func NewUserHandler(usrv application.UserService) UserHandler {
	return UserHandler{
		userService: usrv,
	}
}

func (usrh UserHandler) SignInHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.User
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		res, err := usrh.userService.Register(ctx, req)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserConflict):
				ctx.JSON(http.StatusConflict, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
		ctx.JSON(http.StatusCreated, helpers.DataResponse(0, "User created", res))
	}
}

func (usrh UserHandler) LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.LoginReq
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		res, err := usrh.userService.Login(ctx, req)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidPassword):
				ctx.JSON(http.StatusUnauthorized, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, helpers.DataResponse(0, "Welcome to Api Solutions", res))
	}
}
