package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lstink/blog/internal/routers/api/v1"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	article := v1.NewArticle()
	tag := v1.NewTag()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is a go project!")
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
