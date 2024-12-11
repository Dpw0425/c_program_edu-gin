package web

import v1 "c_program_edu-gin/internal/app/api/handler/web/v1"

type V1 struct {
	Common *v1.Common
}

type Handler struct {
	V1 *V1
}
