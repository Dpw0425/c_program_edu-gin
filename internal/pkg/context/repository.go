package ctx

import (
	"context"
	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
}

type Repo[T ITable] struct {
	model T        // 数据模型
	DB    *gorm.DB // 数据库
}

func NewRepo[T ITable](db *gorm.DB) Repo[T] {
	return Repo[T]{DB: db}
}

func (r Repo[T]) Model(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx).Model(r.model)
}

// IsExist 根据条件查询数据是否存在
func (r Repo[T]) IsExist(ctx context.Context, where string, args ...any) (bool, error) {
	var count int64
	err := r.Model(ctx).Select("1").Where(where, args...).Limit(1).Scan(&count).Error
	if err != nil {
		return false, nil
	}

	return count == 1, nil
}

// FindByID 根据主键 ID 查询
func (r Repo[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var item *T
	err := r.DB.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

// FindByIDs 根据主键 ID 批量查询
func (r Repo[T]) FindByIDs(ctx context.Context, ids []uint) ([]*T, error) {
	var items = make([]*T, 0)
	err := r.DB.WithContext(ctx).Find(&items, ids).Error
	if err != nil {
		return nil, err
	}

	return items, err
}

// Create 创建记录
func (r Repo[T]) Create(ctx context.Context, data *T) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

// BatchCreation 批量创建
func (r Repo[T]) BatchCreation(ctx context.Context, data []*T) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

// FindByWhere 根据条件查询
func (r Repo[T]) FindByWhere(ctx context.Context, where string, args ...any) (*T, error) {
	var item *T
	err := r.DB.WithContext(ctx).Where(where, args...).First(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}
