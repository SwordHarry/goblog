package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goblog/docs"
	"goblog/global"
	"goblog/internal/middleware"
	"goblog/internal/routers/api"
	"goblog/internal/routers/blog"
	"goblog/pkg/limiter"
	"goblog/pkg/util"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

const (
	apiPrefix = "/api/v1"
)

// 限流器
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

var StaticConfig = map[string]string{
	"/css":    "frontend/css/",
	"/js":     "frontend/js/",
	"/images": "frontend/images",
}

func NewRouter() *gin.Engine {
	r := gin.New()
	apiV1 := r.Group(apiPrefix)
	// ————————————————————— begin 通用中间件 ——————————————————————
	// 国际化中间件 注意数据竞态
	r.Use(middleware.Translations())
	// 限流中间件
	r.Use(middleware.RateLimiter(methodLimiters))
	// 超时中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	// 链路追踪中间件
	r.Use(middleware.Tracing())
	apiV1.Use(middleware.Tracing())
	if global.ServerSetting.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		// 访问日志记录中间件
		apiV1.Use(middleware.AccessLog())
		// recover 加 邮件报警 中间件
		r.Use(middleware.Recovery())
	}
	// ————————————————————— end 通用中间件 ——————————————————————
	// 静态文件地址
	r.StaticFS("/static/image", http.Dir(global.AppSetting.UploadImageSavePath))
	// jwt 验证接口
	r.GET("/auth", api.GetAuth)
	// 注册渲染函数
	r.SetFuncMap(template.FuncMap{
		"timeStampToDate": util.TimeStampToDate,
		"add":             util.Add,
		"sub":             util.Sub,
	})
	// 加载 html
	r.LoadHTMLGlob(filepath.Join("frontend", "**", "*.html"))
	// 加载静态资源
	for k, v := range StaticConfig {
		r.Static(k, v)
	}
	// ————————————————————— begin 业务api ——————————————————————
	article := blog.NewArticle()
	tag := blog.NewTag()

	// JWT 中间件
	apiV1.Use(middleware.JWT())
	{
		// tags
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		//r.GET("/tags", tag.List)
		// articles
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		r.GET("/articles/:id", article.Get)
		r.GET("/articles_tag/:id", article.ListByTagID)
		r.GET("/", article.ViewIndex)
		r.POST("/search", article.Search)
		// 注册 swagger 接口文档路由
		apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// ————————————————————— end 业务api ——————————————————————
	return r
}
