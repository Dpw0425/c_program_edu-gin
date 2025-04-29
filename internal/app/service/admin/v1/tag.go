package admin_service

import (
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	"context"
)

var _ ITagService = (*TagService)(nil)

type ITagService interface {
	Add(ctx context.Context, name string) error
}

type TagService struct {
	TagRepo *repo.TagRepo
}

func (t *TagService) Add(ctx context.Context, name string) error {
	if t.TagRepo.IsExist(ctx, name) {
		return myErr.BadRequest("", "分类已存在！")
	}

	return t.TagRepo.Create(ctx, &model.Tag{
		Name: name,
	})
}
