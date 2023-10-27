package auth

import (
	"API_Gateway/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
	svc *ServiceClient
}

func InitMiddleware(svc *ServiceClient) Middleware {
	return Middleware{svc: svc}
}

func (m *Middleware) AuthRequired(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
	}
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	res, err := m.svc.AuthClient.Validate(context.Background(), &pb.ValidateRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Error != "" {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
}
