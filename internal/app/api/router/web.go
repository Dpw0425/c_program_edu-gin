package router

import (
	"c_program_edu-gin/internal/app/api/handler/web"
	"c_program_edu-gin/internal/app/middleware"
	"c_program_edu-gin/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(secret string, router *gin.Engine, handler *web.Handler, storage middleware.IStorage) {
	// authorizer := middleware.Auth(secret, "api", storage)

	v1 := router.Group("/api/v1")
	{
		common := v1.Group("/common")
		{
			common.POST("/send_email_code", ctx.HandlerFunc(handler.V1.Common.SendEmailCode)) // 邮箱验证码
		}

		user := v1.Group("/user")
		{
			user.POST("/register", ctx.HandlerFunc(handler.V1.User.Register)) // 用户注册
			user.POST("/login", ctx.HandlerFunc(handler.V1.User.Login))       // 用户登录
		}
	}
}
