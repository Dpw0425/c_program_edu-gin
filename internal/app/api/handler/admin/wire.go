package admin

import (
	v1 "c_program_edu-gin/internal/app/api/handler/admin/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(v1.Admin), "*"),
	wire.Struct(new(v1.Question), "*"),
	wire.Struct(new(v1.Tag), "*"),

	wire.Struct(new(V1), "*"),
)
