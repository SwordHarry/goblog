package model

import "goblog/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

// 专门用于 swagger 显示返回值
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
