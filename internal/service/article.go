package service

import (
	"goblog/internal/dao"
	"goblog/internal/request"
	"goblog/pkg/app"
)

// service 层 article 返回结构体
type Article struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	Tags          []*tag `json:"tags"`
}

type tag struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

// 获取单个文章
func (svc *Service) GetArticle(param *request.ArticleRequest) (*Article, error) {
	article, err := svc.model.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTagList, err := svc.model.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	var tagList []*tag
	for _, articleTag := range articleTagList {
		t, err := svc.model.GetTagById(articleTag.TagID, 1)
		if err != nil {
			return nil, err
		}
		tagList = append(tagList, &tag{
			ID:   t.ID,
			Name: t.Name,
		})
	}

	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tags:          tagList,
	}, nil
}

// 分页：获取文章列表
func (svc *Service) GetArticleList(param *request.ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	totalRow, err := svc.model.CountArticles(param.State)
	if err != nil {
		return nil, 0, err
	}
	var result []*Article
	modelArticles, err := svc.model.ListArticles(param.State, app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	for _, article := range modelArticles {
		result = append(result, &Article{
			ID:            article.ID,
			Title:         article.Title,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
		})
	}
	return result, totalRow, nil
}

// 通过 tagId 获取文章列表
func (svc *Service) GetArticleListByTagID(param *request.ArticleListByTIDRequest, pager *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.model.CountArticleByTID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := svc.model.ListArticleByTID(param.TagID, param.State, app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.Title,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
		})
	}
	return articleList, articleCount, nil
}

// 创建文章
func (svc *Service) CreateArticle(param *request.CreateArticleRequest) error {
	article, err := svc.model.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Common: &dao.Common{
			CreatedBy: param.CreatedBy,
		},
	})
	if err != nil {
		return err
	}

	if len(param.TagIDList) > 0 {
		for _, tagID := range param.TagIDList {
			err = svc.model.CreateArticleTag(article.ID, tagID, param.CreatedBy)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// 更新文章
func (svc *Service) UpdateArticle(param *request.UpdateArticleRequest) error {
	err := svc.model.UpdateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Common: &dao.Common{
			ID:         param.ArticleID,
			ModifiedBy: param.ModifiedBy,
		},
	})
	if err != nil {
		return err
	}
	//err = svc.dao.UpdateArticleTag(param.ArticleID, param.TagID, param.ModifiedBy)
	//if err != nil {
	//	return err
	//}
	return nil
}

// 删除文章
func (svc *Service) DeleteArticle(param *request.DeleteArticleRequest) error {
	err := svc.model.DeleteArticle(param.ArticleID)
	if err != nil {
		return err
	}
	err = svc.model.DeleteArticleTagByAID(param.ArticleID)
	if err != nil {
		return err
	}
	return nil
}
