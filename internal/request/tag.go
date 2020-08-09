package request

// 针对入参校验增加绑定和验证结构体
// form 表示表单的映射字段名，binding:入参校验的规则内容
type CountTagRequest struct {
	Name  string `form:"name" json:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" json:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" json:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" json:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" json:"id" binding:"required,gte=1"`
	Name       string `form:"name" json:"name" binding:"min=2,max=100"`
	State      uint8  `form:"state" json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" json:"id" binding:"required,gte=1"`
}
