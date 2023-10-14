package routes

import (
	"API_Gateway/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifyPhoneRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

func VerifyPhone(ctx *gin.Context, c pb.AuthServiceClient) {
	b := VerifyPhoneRequestBody{}

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.VerifyPhone(context.Background(), &pb.VerifyPhoneRequest{
		PhoneNumber: b.PhoneNumber,
		Code:        b.Code,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
