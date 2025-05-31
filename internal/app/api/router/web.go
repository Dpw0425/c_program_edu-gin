package router

import (
	"c_program_edu-gin/internal/app/api/handler/web"
	"c_program_edu-gin/internal/app/middleware"
	"c_program_edu-gin/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(secret string, router *gin.Engine, handler *web.Handler, storage middleware.IStorage) {
	authorizer := middleware.Auth(secret, "api", storage)

	v1 := router.Group("/api/v1")
	{
		common := v1.Group("/common")
		{
			common.POST("/send_email_code", ctx.HandlerFunc(handler.V1.Common.SendEmailCode)) // 邮箱验证码
		}

		upload := v1.Group("/upload")
		{
			upload.POST("/avatar", ctx.HandlerFunc(handler.V1.Upload.Avatar)) // 上传头像
		}

		user := v1.Group("/user")
		{
			user.POST("/register", ctx.HandlerFunc(handler.V1.User.Register)) // 用户注册
			user.POST("/login", ctx.HandlerFunc(handler.V1.User.Login))       // 用户登录
			userAuth := user.Group("").Use(authorizer)
			{
				userAuth.GET("/info", ctx.HandlerFunc(handler.V1.User.Info))        // 用户信息
				userAuth.DELETE("/logout", ctx.HandlerFunc(handler.V1.User.Logout)) // 退出登录
			}
		}

		question := v1.Group("/question")
		{
			question.GET("/list", ctx.HandlerFunc(handler.V1.Question.List))                   // 题目列表
			question.GET("/detail", ctx.HandlerFunc(handler.V1.Question.Detail))               // 题目详情
			question.GET("/test_data", ctx.HandlerFunc(handler.V1.Question.TestData))          // 获取测试数据
			question.POST("/commit_answer", ctx.HandlerFunc(handler.V1.Question.CommitAnswer)) // 提交答案
		}

		bank := v1.Group("/bank").Use(authorizer)
		{
			bank.GET("/list", ctx.HandlerFunc(handler.V1.Bank.List)) // 题库列表
		}

		competition := v1.Group("/competition")
		{
			competition.GET("/list", ctx.HandlerFunc(handler.V1.Competition.List)) // 比赛列表
		}
	}
}
