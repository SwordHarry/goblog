package model

import (
	"goblog/internal/dao"
)

// gorm for sql 相关操作
// where 设置筛选条件，接收 map struct string 作为条件
// offset 偏移量，用于指定开始返回记录之前要跳过的记录数
// limit 限制检索的记录数
// find 查找符合条件记录
// update 更新字段
// delete 删除
// count 统计记录数

func (m *Model) GetTagById(tagId uint32, tagState uint8) (*dao.Tag, error) {
	tag := new(dao.Tag)
	err := m.engine.Where("id = ? and is_del = ? and state = ?", tagId, 0, tagState).First(&tag).Error
	return tag, err
}
func (m *Model) GetTagByName(tagName string, tagState uint8) (*dao.Tag, error) {
	tag := new(dao.Tag)
	err := m.engine.Where("name = ? and is_del = ? and state = ?", tagName, 0, tagState).First(&tag).Error
	return tag, err
}

// count tag
func (m *Model) CountTag(tagName string, tagState uint8) (int, error) {
	var count int
	db := m.engine
	t := dao.Tag{
		Name:  tagName,
		State: tagState,
	}
	if tagName != "" {
		db = db.Where("name = ?", tagName)
	}
	db = db.Where("state = ?", tagState)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// list tag
func (m *Model) ListTags(tagName string, tagState uint8, pageOffset, pageSize int) ([]*dao.Tag, error) {
	var tags []*dao.Tag
	var err error
	db := m.engine
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if tagName != "" {
		db = db.Where("name = ?", tagName)
	}
	db = db.Where("state = ?", tagState)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (m *Model) CreateTag(tagName string, tagState uint8, createdBy string) error {
	t := dao.Tag{
		Name:   tagName,
		State:  tagState,
		Common: &dao.Common{CreatedBy: createdBy},
	}
	return m.engine.Create(&t).Error
}

func (m *Model) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	t := dao.Tag{
		Common: &dao.Common{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return m.engine.Model(&t).Where("id = ? and is_del = ?", t.ID, 0).Update(values).Error
}

func (m *Model) DeleteTag(id uint32) error {
	t := dao.Tag{Common: &dao.Common{ID: id}}
	return m.engine.Where("id = ? and is_del = ?", t.ID, 0).Delete(&t).Error
}
