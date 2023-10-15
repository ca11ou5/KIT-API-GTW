package routes

import (
	"API_Gateway/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	b := SignUpRequestBody{}

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SignUp(context.Background(), &pb.SignUpRequest{
		PhoneNumber: b.PhoneNumber,
		Password:    b.Password,
		Name:        b.Name,
		Surname:     b.Surname,
		DateOfBirth: b.DateOfBirth,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": &res})
}
