package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/tag"
	"context"
)

var _ ITagService = (*TagService)(nil)

type ITagService interface {
	Add(ctx context.Context, name string) error
	List(ctx context.Context, request *admin.TagListRequest) ([]*schema.TagItem, error)
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

func (t *TagService) List(ctx context.Context, request *admin.TagListRequest) ([]*schema.TagItem, error) {
	db := t.TagRepo.DB.WithContext(ctx)
	var items []*schema.TagItem

	err := db.Table("tags").Select("id", "name").Where("name LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
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
