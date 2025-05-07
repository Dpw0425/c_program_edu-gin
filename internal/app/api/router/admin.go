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

		user := v1.Group("/user") // 鉴权
		{
			user.POST("/login", ctx.HandlerFunc(handler.V1.Admin.Login)) // 管理员登录

			userAuth := user.Group("").Use(authorizer)
			{
				userAuth.GET("/info", ctx.HandlerFunc(handler.V1.Admin.Info))     // 管理员用户信息
				userAuth.GET("/logout", ctx.HandlerFunc(handler.V1.Admin.Logout)) // 退出登录
			}
		}

		tag := v1.Group("/tag").Use(authorizer) // 分类管理
		{
			tag.POST("/add", ctx.HandlerFunc(handler.V1.Tag.Add))  // 添加分类
			tag.GET("/list", ctx.HandlerFunc(handler.V1.Tag.List)) // 分类列表
		}

		question := v1.Group("/question").Use(authorizer) // 题目管理
		{
			question.POST("/add", ctx.HandlerFunc(handler.V1.Question.Add))  // 添加题目
			question.GET("/list", ctx.HandlerFunc(handler.V1.Question.List)) // 题目列表
		}

		auth := v1.Group("/auth").Use(authorizer) // 权限管理
		{
			auth.GET("/user_list", ctx.HandlerFunc(handler.V1.Auth.UserList))   // 用户列表
			auth.GET("/admin_list", ctx.HandlerFunc(handler.V1.Auth.AdminList)) // 管理员列表
		}
	}
}
