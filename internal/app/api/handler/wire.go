package handler

import (
	"c_program_edu-gin/internal/app/api/handler/web"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(web.Handler), "*"),
	wire.Struct(new(Handler), "*"),
)
