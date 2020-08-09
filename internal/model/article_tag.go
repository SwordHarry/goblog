package model

import (
	"github.com/jinzhu/gorm"
	"goblog/internal/dao"
)

// select articleTag on article_id = ? and is_del = ?
func (m *Model) GetArticleTagByAID(articleID uint32) ([]*dao.ArticleTag, error) {
	var tagList []*dao.ArticleTag
	err := m.engine.Where("article_id = ? and is_del = ?", articleID, 0).Find(&tagList).Error
	if err != nil {
		return nil, err
	}
	return tagList, nil
}

// select articleTag on tag_id = ? and is_del = ?
func (m *Model) ListArticleTagByTID(tagID uint32) ([]*dao.ArticleTag, error) {
	var articleTags []*dao.ArticleTag
	if err := m.engine.Where("tag_id = ? and is_del = ?", tagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}
	return articleTags, nil
}

// select articleTags on article_id in ? and is_del = ?
func (m *Model) ListArticleTagByAIDs(articleIDs []uint32) ([]*dao.ArticleTag, error) {
	var articleTags []*dao.ArticleTag
	err := m.engine.Where("article_id in (?) and is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (m *Model) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := &dao.ArticleTag{
		Common: &dao.Common{
			CreatedBy: createdBy,
		},
		ArticleID: articleID,
		TagID:     tagID,
	}
	return m.engine.Create(articleTag).Error
}

func (m *Model) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := dao.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return m.engine.Model(&articleTag).
		Where("article_id = ? and is_del = ?", articleID, 0).Limit(1).Updates(values).Error
}

func (m *Model) DeleteArticleTagByAID(articleID uint32) error {
	articleTag := dao.ArticleTag{ArticleID: articleID}
	return m.engine.Where("article_id = ? and is_del = ?", articleID, 0).Delete(&articleTag).Error
}

func (m *Model) DeleteArticleTagByTID(tagID uint32) error {
	articleTag := dao.ArticleTag{TagID: tagID}
	return m.engine.Where("tag_id = ? and is_del = ?", tagID, 0).Delete(&articleTag).Error
}

func (m *Model) DeleteOne(articleID uint32) error {
	articleTag := dao.ArticleTag{ArticleID: articleID}
	return m.engine.Where("article_id = ? and is_del = ?", articleID, 0).Delete(&articleTag).Limit(1).Error
}
