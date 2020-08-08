package dao

import "goblog/internal/model"

// dao for article_tag
// get
func (d *Dao) GetArticleTagByAID(articleID uint32) ([]*model.ArticleTag, error) {
	articleTag := &model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := &model.ArticleTag{TagID: tagID}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

// create
func (d *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := &model.ArticleTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		ArticleID: articleID,
		TagID:     tagID,
	}
	return articleTag.Create(d.engine)
}

// update
func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

// delete
func (d *Dao) DeleteArticleTagByAID(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.DeleteByAID(d.engine)
}

func (d *Dao) DeleteArticleTagByTID(tagID uint32) error {
	articleTag := model.ArticleTag{TagID: tagID}
	return articleTag.DeleteByTID(d.engine)
}
