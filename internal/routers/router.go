package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goblog/docs"
	"goblog/global"
	"goblog/internal/middleware"
	"goblog/internal/routers/api"
	"goblog/internal/routers/api/v1"
	"goblog/pkg/limiter"
	"net/http"
	"time"
)

const (
	apiV1Str = "/api/v1"
)

// 限流器
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	// ————————————————————— begin 通用中间件 ——————————————————————
	if global.ServerSetting.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		// 访问日志记录中间件
		r.Use(middleware.AccessLog())
		// recover 加 邮件报警 中间件
		r.Use(middleware.Recovery())
	}
	// 国际化中间件
	r.Use(middleware.Translations())
	// 限流中间件
	r.Use(middleware.RateLimiter(methodLimiters))
	// 超时中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	// 链路追踪中间件
	r.Use(middleware.Tracing())
	// ————————————————————— end 通用中间件 ——————————————————————
	// 注册 swagger 接口文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 静态文件地址
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	// jwt 验证接口
	r.GET("/auth", api.GetAuth)

	// ————————————————————— begin 业务api ——————————————————————
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group(apiV1Str)
	// JWT 中间件
	apiV1.Use(middleware.JWT())
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
		apiV1.POST("/upload", api.UploadFile)
	}
	// ————————————————————— end 业务api ——————————————————————
	return r
}
