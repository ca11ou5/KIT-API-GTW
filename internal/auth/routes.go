package auth

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine, cfg *configs.Config) *ServiceClient {
	svc := &ServiceClient{AuthClient: InitAuthClient(cfg)}

	auth := r.Group("auth")
	auth.GET("sign-up", svc.RenderSignUp)
	auth.POST("sign-up", svc.SignUp)
	auth.POST("verify-phone", svc.VerifyPhone)
	auth.POST("send-code", svc.CheckUser)

	return svc
}

func (c *ServiceClient) RenderSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "sign-up.html", gin.H{})
}

func (c *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, c.AuthClient)
}

func (c *ServiceClient) VerifyPhone(ctx *gin.Context) {
	routes.VerifyPhone(ctx, c.AuthClient)
}

func (c *ServiceClient) CheckUser(ctx *gin.Context) {
	routes.CheckUser(ctx, c.AuthClient)
}
