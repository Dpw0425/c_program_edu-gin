package router

import (
	"c_program_edu-gin/internal/app/api/handler"
	"c_program_edu-gin/internal/app/middleware"
	"c_program_edu-gin/internal/app/storage/cache"
	"c_program_edu-gin/internal/config"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRouter(conf *config.Config, handler *handler.Handler, session *cache.JwtTokenStorage) *gin.Engine {
	router := gin.New()

	router.Static("/public", "./temp/uploads/public")

	// 新增路由过滤规则
	accessFilterRule := middleware.NewAccessFilterRule()
	accessFilterRule.AddRule("/api/v1/user/login", func(req *middleware.RequestInfo) {
		req.RequestBody, _ = sjson.Set(req.RequestBody, `password`, "过滤敏感字段")
	})

	// 定义路由过滤器日志
	router.Use(middleware.AccessLog(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/logs/router_log/access.log", conf.Log.Path), // 日志文件的位置
		MaxSize:    100,                                                         // 文件最大尺寸（以MB为单位）
		MaxAge:     7,                                                           // 保留旧文件的最大天数
		MaxBackups: 3,                                                           // 保留的最大旧文件数量
		LocalTime:  true,                                                        // 使用本地时间创建时间戳
		Compress:   true,                                                        // 是否压缩/归档旧文件
	}, accessFilterRule))

	// capture panic and response server error
	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		response.ErrResponse(c, myErr.InternalServerError("", "系统错误，请重试!!!"))
	}))

	router.GET("/", func(c *gin.Context) {
		response.NorResponse(c, gin.H{}, "hello world!")
	})

	RegisterWebRouter(conf.Jwt.Secret, router, handler.Web, session)
	RegisterAdminRouter(conf.Jwt.Secret, router, handler.Admin, session)

	return router
}
