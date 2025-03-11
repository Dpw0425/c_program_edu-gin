package router

import (
	"c_program_edu-gin/internal/app/api/handler/admin"
	"c_program_edu-gin/internal/app/middleware"
	ctx "c_program_edu-gin/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(secret string, router *gin.Engine, handler *admin.Handler, storage middleware.IStorage) {
	// 鉴权中间件
	authorizer := middleware.Auth(secret, "user", storage)

	v1 := router.Group("/user/v1")
	{
		index := v1.Group("/index").Use(authorizer)
		{
			index.GET("")
		}

		user := v1.Group("/user")
		{
			user.GET("/login", ctx.HandlerFunc(handler.V1.Admin.Login))
		}
	}
}
