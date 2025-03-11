package admin

import v1 "c_program_edu-gin/internal/app/api/handler/admin/v1"

type V1 struct {
	Admin *v1.Admin
}

type Handler struct {
	V1 *V1
}
