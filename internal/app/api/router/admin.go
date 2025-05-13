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
			tag.POST("/add", ctx.HandlerFunc(handler.V1.Tag.Add))         // 添加分类
			tag.GET("/list", ctx.HandlerFunc(handler.V1.Tag.List))        // 分类列表
			tag.POST("/update", ctx.HandlerFunc(handler.V1.Tag.Update))   // 修改分类
			tag.DELETE("/delete", ctx.HandlerFunc(handler.V1.Tag.Delete)) // 删除分类
			tag.GET("/get_all", ctx.HandlerFunc(handler.V1.Tag.GetAll))   // 获取全部分类
		}

		question := v1.Group("/question").Use(authorizer) // 题目管理
		{
			question.POST("/add", ctx.HandlerFunc(handler.V1.Question.Add))                      // 添加题目
			question.GET("/list", ctx.HandlerFunc(handler.V1.Question.List))                     // 题目列表
			question.POST("/update", ctx.HandlerFunc(handler.V1.Question.Update))                // 修改题目
			question.DELETE("/delete", ctx.HandlerFunc(handler.V1.Question.Delete))              // 删除题目
			question.POST("/add_test", ctx.HandlerFunc(handler.V1.Question.AddTestData))         // 添加测试数据
			question.GET("/get_test", ctx.HandlerFunc(handler.V1.Question.GetTestData))          // 查询测试数据
			question.POST("/update_test", ctx.HandlerFunc(handler.V1.Question.UpdateTestData))   // 更新测试数据
			question.DELETE("/delete_test", ctx.HandlerFunc(handler.V1.Question.DeleteTestData)) // 删除测试数据
		}

		auth := v1.Group("/auth").Use(authorizer) // 权限管理
		{
			auth.GET("/user_list", ctx.HandlerFunc(handler.V1.Auth.UserList))          // 用户列表
			auth.GET("/admin_list", ctx.HandlerFunc(handler.V1.Auth.AdminList))        // 管理员列表
			auth.POST("/add_admin", ctx.HandlerFunc(handler.V1.Auth.AddAdmin))         // 新增管理员
			auth.POST("/update_admin", ctx.HandlerFunc(handler.V1.Auth.UpdateAdmin))   // 修改管理员信息
			auth.DELETE("/delete_admin", ctx.HandlerFunc(handler.V1.Auth.DeleteAdmin)) // 删除管理员
			auth.POST("/update_user", ctx.HandlerFunc(handler.V1.Auth.UpdateUser))     // 修改用户信息
			auth.DELETE("/delete_user", ctx.HandlerFunc(handler.V1.Auth.DeleteUser))   // 删除用户
		}

		bank := v1.Group("/question_bank").Use(authorizer) // 题库管理
		{
			bank.POST("/add", ctx.HandlerFunc(handler.V1.Bank.Add))                            // 创建题库
			bank.GET("/list", ctx.HandlerFunc(handler.V1.Bank.List))                           // 题库列表
			bank.POST("/update", ctx.HandlerFunc(handler.V1.Bank.Update))                      // 更新题库信息
			bank.DELETE("/delete", ctx.HandlerFunc(handler.V1.Bank.Delete))                    // 删除题库
			bank.GET("/get_question", ctx.HandlerFunc(handler.V1.Bank.GetQuestionInBank))      // 获取题库内题目列表
			bank.GET("/get_beside", ctx.HandlerFunc(handler.V1.Bank.GetQuestionBesideBank))    // 获取非题库内题目列表
			bank.DELETE("/exclude_question", ctx.HandlerFunc(handler.V1.Bank.ExcludeQuestion)) // 删除题库内题目
		}
	}
}
