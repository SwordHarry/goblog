package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a *Tag) TableName() string {
	return "blog_tag"
}
