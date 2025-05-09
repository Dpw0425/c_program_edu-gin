package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type AdminRepo struct {
	ctx.Repo[model.Admin]
}

func NewAdmins(db *gorm.DB) *AdminRepo {
	return &AdminRepo{Repo: ctx.NewRepo[model.Admin](db)}
}

func (a *AdminRepo) FindByTeacherID(ctx context.Context, username string) (*model.Admin, error) {
	return a.Repo.FindByWhere(ctx, "teacher_id = ?", username)
}

func (a *AdminRepo) IsExist(ctx context.Context, teacherId string) bool {
	if len(teacherId) == 0 {
		return false
	}

	isExist, _ := a.Repo.IsExist(ctx, "teacher_id = ?", teacherId)
	return isExist
}

func (a *AdminRepo) Create(ctx context.Context, admin *model.Admin) error {
	if err := a.Repo.Create(ctx, admin); err != nil {
		return err
	}

	return nil
}
