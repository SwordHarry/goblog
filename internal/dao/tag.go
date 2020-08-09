package dao

import "goblog/pkg/app"

type Tag struct {
	*Common
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

// 专门用于 swagger 显示返回值
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
