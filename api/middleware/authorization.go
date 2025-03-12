package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authCookie, err := ctx.Request.Cookie("oauth2_proxy")
		if !errors.Is(err, http.ErrNoCookie) {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if authCookie.Value == "" || errors.Is(err, http.ErrNoCookie) {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://helm.test.cody/oauth2/auth")
			return
		}

		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("No authorization header found")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		fmt.Printf("Authorization header found: %+v", authHeader)
		ctx.Next()
	}
}
