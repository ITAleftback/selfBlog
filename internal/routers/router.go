package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfblog/global"
	"selfblog/internal/middleware"
	"selfblog/internal/routers/api"
	v1 "selfblog/internal/routers/api/v1"
	"selfblog/pkg/limiter"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "selfblog/docs"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine{
	r:=gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag:=v1.NewTag()
	article:=v1.NewArticle()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/id",tag.Delete)
		apiv1.PUT("/tags/id",tag.Update)
		apiv1.PATCH("/tags/id/state",tag.Update)
		apiv1.GET("/tags",tag.List)

		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/id",article.Delete)
		apiv1.PUT("/articles/id",article.Update)
		apiv1.PATCH("/articles/id/state",article.Update)
		apiv1.GET("/articles/id",article.Get)
		apiv1.GET("/articles",article.List)
	}

	return r
}
