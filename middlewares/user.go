package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"weblog/constants"
	"weblog/utils"
)

func Authentication(c *gin.Context) {
	fmt.Println("[AuthenticationMiddleware] Checking for authentication of user")

	authorization := c.Request.Header.Get(constants.Authorization)
	token := strings.Replace(authorization, constants.Bearer, "", 1)

	if token == "" {
		errMessage := constants.ErrNoAuthHeader
		fmt.Println("[AuthenticationMiddleware]", errMessage)
		errRes := utils.CreateErrorResponse(http.StatusBadRequest, errMessage)
		c.JSON(http.StatusBadRequest, errRes)
		c.Abort()
		return
	}

	claims, err := utils.ValidateToken(token)

	if err != nil {
		fmt.Println("[AuthenticationMiddleware]", err.Error())
		errRes := utils.CreateErrorResponse(http.StatusUnauthorized, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}

	fmt.Println("[AuthenticationMiddleware] User is authenticated")
	c.Set(constants.Email, claims.Email)
	c.Set(constants.Role, claims.Role)
	c.Next()
}
