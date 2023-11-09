package utils

import (
	"github.com/gin-gonic/gin"
	"weblog/constants"
)

func SetAccessTokenCookie(c *gin.Context, accessToken string) {
	c.SetCookie(
		constants.AccessTokenCookie,
		accessToken, constants.JwtAccessTokenTimeDuration,
		constants.HomePath,
		constants.LocalHost,
		true,
		true,
	)
}
