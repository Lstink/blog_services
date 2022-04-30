package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lstink/blog/global"
	"github.com/lstink/blog/pkg/convert"
)

// GetPage 获取页数量
func GetPage(c *gin.Context) int {
	page := convert.Strto(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

// GetPageSize 获取每页展示条数
func GetPageSize(c *gin.Context) int {
	pageSize := convert.Strto(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// GetPageOffset 获取页数的偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
