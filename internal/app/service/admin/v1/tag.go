package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/tag"
	"context"
	"strings"
)

var _ ITagService = (*TagService)(nil)

type ITagService interface {
	Add(ctx context.Context, name string) error
	List(ctx context.Context, request *admin.TagListRequest) ([]*schema.TagItem, int, error)
	Update(ctx context.Context, request *admin.TagUpdateRequest) error
	Delete(ctx context.Context, request *admin.DeleteTagRequest) error
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

func (t *TagService) List(ctx context.Context, request *admin.TagListRequest) ([]*schema.TagItem, int, error) {
	db := t.TagRepo.DB.WithContext(ctx)
	var items []*schema.TagItem

	err := db.Table("tags").Select("id", "name").Where("name LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	for _, item := range items {
		var count int64
		if err := db.Table("questions").Where("status = 1 and FIND_IN_SET(?,tag)", item.Name).Count(&count).Error; err == nil {
			item.Count = int(count)
		} else {
			return nil, 0, myErr.NotFound("", err.Error())
		}
	}

	var count int64
	db.Table("tags").Count(&count)

	return items, int(count), nil
}

func (t *TagService) Update(ctx context.Context, request *admin.TagUpdateRequest) error {
	db := t.TagRepo.DB.WithContext(ctx).Begin()

	// 查找原标签
	var tag *model.Tag
	if err := db.Table("tags").Where("id = ?", request.Id).Find(&tag).Error; err != nil {
		db.Rollback()
		return myErr.BadRequest("", "所选标签不存在！")
	}
	// 修改标签
	if err := db.Table("tags").Where("id = ?", request.Id).UpdateColumn("name", request.Name).Error; err != nil {
		db.Rollback()
		return myErr.BadRequest("", "修改失败！")
	}

	// 修改已使用该标签的题目
	var queList []*model.Question
	db.Table("questions").Where("tag LIKE ?", "%"+tag.Name+"%").Scan(&queList)
	for _, question := range queList {
		// 构造新字符串
		oldTags := strings.Split(question.Tag, tag.Name)
		oldTags[0] = oldTags[0] + request.Name
		newTags := strings.Join(oldTags, "")
		// 写入新字符串
		if err := db.Table("questions").Where("id = ?", question.ID).UpdateColumn("tag", newTags).Error; err != nil {
			db.Rollback()
			return myErr.BadRequest("", "修改失败！")
		}
	}

	db.Commit()

	return nil
}

func (t *TagService) Delete(ctx context.Context, request *admin.DeleteTagRequest) error {
	db := t.TagRepo.DB.WithContext(ctx)
	return db.Unscoped().Where("id = ?", request.Id).Delete(&model.Tag{}).Error
}
