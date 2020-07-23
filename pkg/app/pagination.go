package app

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/pkg/convert"
)

// 分页处理

// 将参数中的 page 参数提取出来
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MuseInt()
	if page <= 0 {
		return 1
	}
	return page
}

// 将参数中的 pageSize 参数提取出来
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MuseInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

// 获取页数偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
