package web

import (
	v1 "c_program_edu-gin/internal/app/api/handler/web/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(v1.Common), "*"),
	wire.Struct(new(v1.User), "*"),
	wire.Struct(new(v1.Upload), "*"),

	wire.Struct(new(V1), "*"),
)
