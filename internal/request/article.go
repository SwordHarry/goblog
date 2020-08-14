package request

type ArticleRequest struct {
	ID    uint32 `form:"id" json:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListByTIDRequest struct {
	TagID uint32 `form:"tag_id" json:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" json:"tag_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" json:"desc" binding:"max=255"`
	Content       string `form:"content" json:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url"`
	CreatedBy     string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ArticleID     uint32 `form:"article_id" json:"article_id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" json:"tag_id"`
	Title         string `form:"title" json:"title"`
	Desc          string `form:"desc" json:"desc"`
	Content       string `form:"content" json:"content"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url"`
	ModifiedBy    string `form:"modified_by" json:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ArticleID uint32 `form:"article_id" json:"article_id" binding:"required,gte=1"`
}
