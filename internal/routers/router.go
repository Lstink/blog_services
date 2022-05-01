package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lstink/blog/docs"
	"github.com/lstink/blog/internal/middleware"
	v1 "github.com/lstink/blog/internal/routers/api/v1"
	"github.com/lstink/blog/pkg/app"
	"github.com/lstink/blog/pkg/errcode"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 加载中间件
	r.Use(middleware.Translations())
	// 注册swagger API接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()

	r.GET("/", func(ctx *gin.Context) {
		app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
		return
	})

	apiv1 := r.Group("api/v1")
	{
		// 标签
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)
		// 文章
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Update)
		apiv1.GET("/articles", article.List)
	}

	return r
}
