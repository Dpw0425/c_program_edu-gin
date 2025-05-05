package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	"context"
)

var _ ITagService = (*TagService)(nil)

type ITagService interface {
	Add(ctx context.Context, name string) error
	List(ctx context.Context, search string) ([]*schema.TagItem, error)
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

func (t *TagService) List(ctx context.Context, search string) ([]*schema.TagItem, error) {
	db := t.TagRepo.DB.WithContext(ctx)
	var items []*schema.TagItem

	err := db.Table("tags").Select("name").Where("name LIKE ?", "%"+search+"%").Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	for _, item := range items {
		var count int64
		if err := db.Table("questions").Where("status = 1 and FIND_IN_SET(?,tag)", item.Name).Count(&count).Error; err == nil {
			item.Count = int(count)
		} else {
			return nil, myErr.NotFound("", err.Error())
		}
	}

	return items, nil
}
