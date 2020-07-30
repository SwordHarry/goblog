package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goblog/docs"
	"goblog/global"
	"goblog/internal/middleware"
	"goblog/internal/routers/api/v1"
	"net/http"
)

const (
	apiV1Str = "/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	// 注册 swagger 接口文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 静态文件地址
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group(apiV1Str)
	{
		// tags
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)
		// articles
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
		// upload
		apiV1.POST("/upload", UploadFile)
	}
	return r
}
