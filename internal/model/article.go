package model

import (
	"github.com/jinzhu/gorm"
	"goblog/internal/dao"
)

func (m *Model) CreateArticle(a *dao.Article) (*dao.Article, error) {
	if err := m.engine.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (m *Model) UpdateArticle(param *dao.Article) error {
	article := dao.Article{Common: &dao.Common{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["Title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return m.engine.Model(&article).Updates(values).Where("id = ? and is_del = ?", article.ID, 0).Error
}

func (m *Model) GetArticle(id uint32, state uint8) (*dao.Article, error) {
	article := new(dao.Article)
	db := m.engine.Where("id = ? and state = ? and is_del = ?", id, state, 0)
	// first 方法用于查询单条记录
	err := db.First(article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

// list articles with pagination
func (m *Model) ListArticles(state uint8, pageOffset, pageSize int) ([]*dao.Article, error) {
	var result []*dao.Article
	db := m.engine
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	err := db.Where("state = ? and is_del = ?", state, 0).Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Model) DeleteArticle(id uint32) error {
	article := dao.Article{Common: &dao.Common{ID: id}}
	return m.engine.Where("id = ? and is_del = ?", id, 0).Delete(&article).Error
}

func (m *Model) CountArticles(state uint8) (int, error) {
	var count int
	a := dao.Article{State: state}
	err := m.engine.Table(a.TableName()).Where("state = ?", state).Count(&count).Error
	return count, err
}

// list articles by tag id
func (m *Model) ListArticleByTID(tagID uint32, state uint8, pageOffset, pageSize int) ([]*dao.ArticleRow, error) {
	fields := []string{
		"ar.id as article_id",
		"ar.title as article_title",
		"ar.desc as article_desc",
		"ar.cover_image_url",
		"ar.content",
	}
	fields = append(fields, []string{
		"t.id as tag_id",
		"t.name as tag_name",
	}...)
	db := m.engine
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	articleTag := new(dao.ArticleTag)
	tag := new(dao.Tag)
	article := new(dao.Article)
	// select：指定要检索的字段，若不指定，则为 select *
	// joins：指定关联查询的语句
	// rows：执行sql语句并获取查询结果
	rows, err := db.Select(fields).Table(articleTag.TableName()+" as at").
		Joins("left join `"+tag.TableName()+"` as t on at.tag_id = t.id").
		Joins("left join `"+article.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.`tag_id` = ? and ar.state = ? and ar.is_del = ?", tagID, state, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*dao.ArticleRow
	for rows.Next() {
		r := &dao.ArticleRow{}
		if err := rows.Scan(
			&r.ArticleID,
			&r.Title,
			&r.Desc,
			&r.CoverImageUrl,
			&r.Content,
			&r.TagID,
			&r.TagName,
		); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (m *Model) CountArticleByTID(tagID uint32, state uint8) (int, error) {
	var count int
	a := new(dao.Article)
	articleTag := new(dao.ArticleTag)
	tag := new(dao.Tag)
	err := m.engine.Table(articleTag.TableName()+" as at").
		Joins("left join `"+tag.TableName()+"` as t on at.tag_id= t.id").
		Joins("left join `"+a.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.`tag_id` = ? and ar.state = ? and ar.is_del = ?", tagID, state, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
