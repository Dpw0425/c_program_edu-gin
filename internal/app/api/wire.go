package api

import (
	"c_program_edu-gin/internal/config"
	"github.com/google/wire"
)

type AppProvider struct {
	Config *config.Config
	// Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
)
