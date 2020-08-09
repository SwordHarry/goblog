package dao

type ArticleTag struct {
	*Common
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a *ArticleTag) TableName() string {
	return "blog_article_tag"
}
