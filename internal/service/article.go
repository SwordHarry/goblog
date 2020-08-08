package service

import (
	"goblog/internal/dao"
	"goblog/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" json:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListByTIDRequest struct {
	TagID uint32 `form:"tag_id" json:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagIDList     []uint32 `form:"tag_ids" json:"tag_ids"`
	Title         string   `form:"title" json:"title" binding:"required,min=2,max=100"`
	Desc          string   `form:"desc" json:"desc" binding:"required,min=2,max=255"`
	Content       string   `form:"content" json:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string   `form:"cover_image_url" json:"cover_image_url" binding:"url"`
	CreatedBy     string   `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	State         uint8    `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

// TODO: 理清楚 update article 的时候，tag 的关系
type UpdateArticleRequest struct {
	ArticleID     uint32 `form:"article_id" json:"article_id" binding:"required,gte=1"`
	Title         string `form:"title" json:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" json:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" json:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" json:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ArticleID uint32 `form:"article_id" json:"article_id" binding:"required,gte=1"`
}

// service 层 article 返回结构体
type Article struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	// TODO: 将 articles 与 tag 进行 一对多 关系
	Tags []*tag `json:"tags"`
}

type tag struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

// 获取单个文章
func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTagList, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	var tagList []*tag
	for _, articleTag := range articleTagList {
		t, err := svc.dao.GetTag(articleTag.TagID, 1)
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
func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	totalRow, err := svc.dao.CountArticles(param.State)
	if err != nil {
		return nil, 0, err
	}
	var result []*Article
	modelArticles, err := svc.dao.ListArticles(param.State, pager.Page, pager.PageSize)
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
func (svc *Service) GetArticleListByTagID(param *ArticleListByTIDRequest, pager *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
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
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
	if err != nil {
		return err
	}

	if len(param.TagIDList) > 0 {
		for _, tagID := range param.TagIDList {
			err = svc.dao.CreateArticleTag(article.ID, tagID, param.CreatedBy)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// 更新文章
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ArticleID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
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
func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ArticleID)
	if err != nil {
		return err
	}
	err = svc.dao.DeleteArticleTagByAID(param.ArticleID)
	if err != nil {
		return err
	}
	return nil
}
