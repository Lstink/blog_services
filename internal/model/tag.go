package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lstink/blog/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Paper *app.Pager
}

// Count 查询数据长度
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("state = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var (
		tags []*Tag
		err  error
	)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// Create 新增数据
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新数据
func (t Tag) Update(db *gorm.DB, values any) error {
	err := db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID).Updates(values).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete 删除数据
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
