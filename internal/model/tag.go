package model

import (
	"github.com/jinzhu/gorm"
	"goblog/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

// 专门用于 swagger 显示返回值
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// gorm for sql 相关操作
// where 设置筛选条件，接收 map struct string 作为条件
// offset 偏移量，用于指定开始返回记录之前要跳过的记录数
// limit 限制检索的记录数
// find 查找符合条件记录
// update 更新字段
// delete 删除
// count 统计记录数
// count tag
func (t *Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// list tag
func (t *Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
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

func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Tag) Update(db *gorm.DB, value interface{}) error {
	return db.Model(t).Where("id = ? and is_del = ?", t.ID, 0).Update(value).Error
}

func (t *Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.ID, 0).Delete(t).Error
}
