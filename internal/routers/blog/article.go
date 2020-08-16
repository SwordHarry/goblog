package blog

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/internal/dao"
	"goblog/internal/request"
	"goblog/internal/routers/api"
	"goblog/internal/routers/common"
	"goblog/internal/service"
	"goblog/pkg/app"
	"goblog/pkg/convert"
	"goblog/pkg/errcode"
	"goblog/pkg/upload"
	"goblog/pkg/util"
	"html/template"
	"net/http"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

const (
	indexHtml   = "index.html"
	articleHtml = "article.html"
)

// @Summary 获取单个文章
// @Produce html
// @Param id path int true "文章ID"
// @Success 200 {object} dao.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /articles/{id} [get]
func (a *Article) Get(c *gin.Context) {
	param := request.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		common.ViewErr(c, errcode.ErrorGetArticleFail)
		return
	}
	if article.CoverImageUrl != "" {
		article.CoverImageUrl = global.AppSetting.UploadImageServerUrl + "/" + article.CoverImageUrl
	}
	// 获取 md 文件
	s, err := util.GetMd(article.Content)
	if err != nil {
		global.Logger.Errorf(c, "util.GetMd err: %v", err)
		common.ViewErr(c, err)
	}
	// 获取 标签 列表
	tags, err := getTags(&svc)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		common.ViewErr(c, err)
	}
	common.View(c, articleHtml, gin.H{
		"article": article,
		"content": template.HTML(s),
		"tags":    tags,
	})
}

// @Summary 通过 tagID 获取多个文章
// @Produce json
// @Param name query string false "文章名称"
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dao.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles_tag [get]
func (a *Article) ListByTagID(c *gin.Context) {
	param := request.ArticleListByTIDRequest{TagID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, count, err := svc.GetArticleListByTagID(&param, &pager)
	pager.TotalRows = count
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleListByTagID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	// 获取 标签 列表
	tags, err := getTags(&svc)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		common.ViewErr(c, err)
	}
	common.View(c, indexHtml, gin.H{
		"articles":  articles,
		"tags":      tags,
		"pager":     &pager,
		"totalPage": getTotalPage(&pager),
	})
	return
}

// @Summary 获取多个文章
// @Produce json
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dao.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a *Article) ViewIndex(c *gin.Context) {
	param := request.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	// 获取 文章 列表
	articles, count, err := svc.GetArticleList(&param, &pager)
	pager.TotalRows = count
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v", err)
		common.ViewErr(c, errcode.ErrorGetArticlesFail)
		return
	}

	// 获取 标签 列表
	tags, err := getTags(&svc)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		common.ViewErr(c, err)
	}
	common.View(c, indexHtml, gin.H{
		"articles":  articles,
		"tags":      tags,
		"pager":     &pager,
		"totalPage": getTotalPage(&pager),
	})
}

// —————————————————————————————— api begin ————————————————————————————————

// @Summary 创建文章
// @Produce json
// @Accept multipart/form-data
// @Param tag_id formData int false "标签ID列表"
// @Param title formData string true "文章标题"
// @Param desc formData string true "文章简述"
// @Param created_by formData string true "创建者"
// @Param state formData int false "状态"
// @Param md formData file true "md博客文件"
// @Param img formData file true "cover封面图"
// @Success 200 {object} dao.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a *Article) Create(c *gin.Context) {
	param := request.CreateArticleRequest{}
	response := app.NewResponse(c)
	// 上传 md 和 image
	mdInfo, ok := delWithUploadFile(c, response, upload.TypeMd, false)
	if ok {
		param.Content = mdInfo.Name
		param.Title = mdInfo.OriginName
	} else {
		return
	}
	if imageInfo, ok := delWithUploadFile(c, response, upload.TypeImage, true); ok {
		if imageInfo != nil {
			param.CoverImageUrl = imageInfo.Name
		}
	}
	// 绑定剩余参数
	valid, errs := app.BindAndValid(c, &param)
	if param.Desc == "" {
		// 读取 md 文件
		s, err := util.GetMd(mdInfo.Name)
		// 对中文的处理 rune
		rs := []rune(string(s))
		if err != nil {
			global.Logger.Errorf(c, "util.GetMd err: %v", err)
			response.ToErrorResponse(errcode.ErrorUploadFileFail)
			return
		}
		// 获取 md 文件一部分字符串生成 desc
		desc := util.Html2Text(string(rs[:800]))
		param.Desc = desc + "..."
	}
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	article, err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{"data": article})
	return
}

// @Summary 更新文章
// @Produce json
// @Param article_id path string true "文章ID"
// @Param title formData string false "文章标题"
// @Param desc formData string false "文章简述"
// @Param cover_image_url formData file false "封面图片地址"
// @Param content formData file false "文章内容"
// @Param modified_by formData string true "修改者"
// @Param tag_id formData string false "附属标签"
// @Success 200 {object} dao.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a *Article) Update(c *gin.Context) {
	param := request.UpdateArticleRequest{ArticleID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	// 上传 md 和 image
	if mdInfo, ok := delWithUploadFile(c, response, upload.TypeMd, true); ok {
		if mdInfo != nil {
			param.Content = mdInfo.Name
			param.Title = mdInfo.OriginName
		}
	}
	if imageInfo, ok := delWithUploadFile(c, response, upload.TypeImage, true); ok {
		if imageInfo != nil {
			param.CoverImageUrl = imageInfo.Name
		}
	}
	// 绑定剩余参数
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(nil)
	return
}

// @Summary 删除文章
// @Produce  json
// @Param article_id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a *Article) Delete(c *gin.Context) {
	param := request.DeleteArticleRequest{ArticleID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(nil)
	return
}

// —————————————————————————————— api end ————————————————————————————————

// 对 controller 层的一些通用逻辑的封装
func delWithUploadFile(c *gin.Context, response *app.Response, fileType upload.FileType, isOptional bool) (*service.FileInfo, bool) {
	fileInfo, err := api.UploadFile(c, fileType)
	// 是否可选
	if isOptional {
		// 若可选：上传了文件出错
		if err != nil && err != http.ErrMissingFile {
			global.Logger.Errorf(c, "api.UploadFile err: %v", err)
			response.ToErrorResponse(errcode.ErrorUploadFileFail)
			return nil, false
		} else {
			// 没上传文件或者上传了没出错
			if err == nil {
				return fileInfo, true
			} else {
				return nil, true
			}
		}
	} else {
		// 不可选：上传了文件出错
		if err != nil {
			global.Logger.Errorf(c, "api.UploadFile err: %v", err)
			response.ToErrorResponse(errcode.ErrorUploadFileFail)
			return nil, false
		}
		return fileInfo, true
	}
}

func getTotalPage(pager *app.Pager) int {
	// 总页数逻辑
	totalPage := pager.TotalRows / pager.PageSize
	if pager.TotalRows%pager.PageSize != 0 {
		totalPage++
	}
	return totalPage
}

func getTags(svc *service.Service) ([]*dao.Tag, error) {
	// 获取 标签 列表
	return svc.GetTagList(&request.TagListRequest{
		State: 1,
	}, &app.Pager{Page: 1, PageSize: 10}) // 暂时写死 1页 10 标签
}
