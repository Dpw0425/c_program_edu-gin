package handler

import (
	"c_program_edu-gin/internal/app/api/handler/admin"
	"c_program_edu-gin/internal/app/api/handler/web"
)

type Handler struct {
	Web   *web.Handler
	Admin *admin.Handler
}
