package dao

import (
	"goblog/pkg/app"
)

// dao 层的 article 参数封装
type Article struct {
	*Common
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

// 关联查询
type ArticleRow struct {
	ID            uint32 // article id
	TagID         uint32
	CreatedOn     uint32
	TagName       string
	Title         string
	Desc          string
	CoverImageUrl string
	Content       string
}

// 专门用于 swagger 显示返回值
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// 文章附带标签列表
type ArticleWithTags struct {
	*Article
	Tags []*Tag
}
