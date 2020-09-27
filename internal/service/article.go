package service

import (
	"goblog/internal/dao"
	"goblog/internal/request"
	"goblog/pkg/app"
)

// 获取单个文章
func (svc *Service) GetArticle(param *request.ArticleRequest) (*dao.ArticleWithTags, error) {
	article, err := svc.model.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTagList, err := svc.model.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	var tagList []*dao.Tag
	for _, articleTag := range articleTagList {
		t, err := svc.model.GetTagById(articleTag.TagID, 1)
		if err != nil {
			return nil, err
		}
		tagList = append(tagList, t)
	}

	return &dao.ArticleWithTags{
		Article: &dao.Article{
			Common: &dao.Common{
				ID:         article.ID,
				CreatedBy:  article.CreatedBy,
				ModifiedBy: article.ModifiedBy,
				CreatedOn:  article.CreatedOn,
				ModifiedOn: article.ModifiedOn,
			},
			Title:         article.Title,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
		},
		Tags: tagList,
	}, nil
}

// 分页：获取文章列表
func (svc *Service) GetArticleList(param *request.ArticleListRequest, pager *app.Pager) ([]*dao.Article, int, error) {
	totalRow, err := svc.model.CountArticles(param.State)
	if err != nil {
		return nil, 0, err
	}
	//var result []*Article
	modelArticles, err := svc.model.ListArticles(param.State, app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return modelArticles, totalRow, nil
}

// 通过 tagId 获取文章列表
func (svc *Service) GetArticleListByTagID(param *request.ArticleListByTIDRequest, pager *app.Pager) ([]*dao.ArticleRow, int, error) {
	articleCount, err := svc.model.CountArticleByTID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := svc.model.ListArticleByTID(param.TagID, param.State, app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return articles, articleCount, nil
}

// 创建文章
func (svc *Service) CreateArticle(param *request.CreateArticleRequest) (*dao.Article, error) {
	return svc.model.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Common: &dao.Common{
			CreatedBy: param.CreatedBy,
		},
	}, param.TagID)
}

// 更新文章
func (svc *Service) UpdateArticle(param *request.UpdateArticleRequest) error {
	return svc.model.UpdateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Common: &dao.Common{
			ID:         param.ArticleID,
			ModifiedBy: param.ModifiedBy,
		},
	}, param.TagID)
}

// 删除文章
func (svc *Service) DeleteArticle(param *request.DeleteArticleRequest) error {
	return svc.model.DeleteArticle(param.ArticleID)
}

// 根据 title 获取文章
func (svc *Service) SearchArticlesByTitle(param *request.SearchArticleRequest, pager *app.Pager) ([]*dao.Article, int, error) {
	param.Title = "%" + param.Title + "%"
	totalRow, err := svc.model.CountArticlesByTitle(param.State, param.Title)
	if err != nil {
		return nil, 0, err
	}
	articles, err := svc.model.SearchArticlesByTitle(
		param.Title, param.State,
		app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return articles, totalRow, nil
}
