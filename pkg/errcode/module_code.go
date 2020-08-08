package errcode

// 针对标签模块的错误码
var (
	// tag
	ErrorGetTagListFail       = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail        = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail        = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail        = NewError(20010004, "删除标签失败")
	ErrorCountTagFail         = NewError(20010005, "统计标签失败")
	ErrorCreateTagExistedFail = NewError(20010006, "创建失败，标签已存在")
	// article
	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")
	// upload
	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
