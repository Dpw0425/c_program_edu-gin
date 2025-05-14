package admin

import v1 "c_program_edu-gin/internal/app/api/handler/admin/v1"

type V1 struct {
	Admin       *v1.Admin
	Question    *v1.Question
	Tag         *v1.Tag
	Auth        *v1.Auth
	Bank        *v1.Bank
	Competition *v1.Competition
}

type Handler struct {
	V1 *V1
}
