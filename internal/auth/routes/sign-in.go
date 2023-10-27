package routes

import (
	"API_Gateway/internal/auth/pb"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignInRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func SignIn(ctx *gin.Context, c pb.AuthServiceClient) {
	b := SignInRequestBody{}

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SignIn(context.Background(), &pb.SignInRequest{
		PhoneNumber: b.PhoneNumber,
		Password:    b.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Error != "" {
		ctx.AbortWithError(int(res.Status), errors.New(res.Error))
		return
	}

	ctx.SetCookie("access_token", res.AccessToken, 1800, "/", "localhost", false, false)
	ctx.SetCookie("refresh_token", res.RefreshToken, 604800, "/", "localhost", false, true)

	ctx.JSON(int(res.Status), &res)
}
