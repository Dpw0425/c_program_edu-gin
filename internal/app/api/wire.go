package api

import (
	"c_program_edu-gin/internal/app/api/handler"
	"c_program_edu-gin/internal/app/api/handler/web"
	"c_program_edu-gin/internal/app/api/router"
	"c_program_edu-gin/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type AppProvider struct {
	Config *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	router.NewRouter,

	handler.ProviderSet,
	web.ProviderSet,

	wire.Struct(new(AppProvider), "*"),
)
