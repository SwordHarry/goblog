package dao

import (
	"goblog/internal/model"
	"goblog/pkg/app"
)

// dao 层的 article 参数封装
type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Model:         &model.Model{CreatedBy: param.CreatedBy},
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
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
	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) ListArticles(state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{State: state}
	return article.ListArticles(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) CountArticles(state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountArticles(d.engine, state)
}
