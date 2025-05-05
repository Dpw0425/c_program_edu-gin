package router

import (
	"c_program_edu-gin/internal/app/api/handler/admin"
	"c_program_edu-gin/internal/app/middleware"
	ctx "c_program_edu-gin/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(secret string, router *gin.Engine, handler *admin.Handler, storage middleware.IStorage) {
	// 鉴权中间件
	authorizer := middleware.Auth(secret, "admin", storage)

	v1 := router.Group("/admin/v1")
	{
		index := v1.Group("/index").Use(authorizer)
		{
			index.GET("")
		}

		user := v1.Group("/user")
		{
			user.POST("/login", ctx.HandlerFunc(handler.V1.Admin.Login))

			userAuth := user.Group("").Use(authorizer)
			{
				userAuth.GET("/info", ctx.HandlerFunc(handler.V1.Admin.Info))
				userAuth.GET("/logout", ctx.HandlerFunc(handler.V1.Admin.Logout))
			}
		}

		tag := v1.Group("/tag").Use(authorizer)
		{
			tag.POST("/add", ctx.HandlerFunc(handler.V1.Tag.Add))
			tag.GET("/list", ctx.HandlerFunc(handler.V1.Tag.List))
		}

		question := v1.Group("/question").Use(authorizer)
		{
			question.POST("/add", ctx.HandlerFunc(handler.V1.Question.Add))
		}
	}
}
