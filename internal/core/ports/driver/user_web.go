package driverport

import (
	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	SignInHandler() gin.HandlerFunc
	LoginHandler() gin.HandlerFunc
}
