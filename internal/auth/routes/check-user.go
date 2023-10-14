package routes

import (
	"API_Gateway/internal/auth/pb"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func CheckUser(ctx *gin.Context, c pb.AuthServiceClient) {
	b := CheckUserRequestBody{}

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CheckUser(context.Background(), &pb.CheckUserRequest{
		PhoneNumber: b.PhoneNumber,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Error != "" {
		ctx.AbortWithError(int(res.Status), errors.New(res.Error))
		return
	}

	ctx.JSON(int(res.Status), &res)
}
