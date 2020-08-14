package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/internal/service"
	"goblog/pkg/upload"
	"path"
	"strings"
)

// 上传文件
func UploadFile(c *gin.Context, fileType upload.FileType) (*service.FileInfo, error) {
	var fileName string
	switch fileType {
	case upload.TypeMd:
		fileName = "md"
	case upload.TypeImage:
		fileName = "img"
	default:
		fileName = "md"
	}
	file, fileHeader, err := c.Request.FormFile(fileName)
	if err != nil {
		return nil, err
	}
	if fileHeader == nil || fileType <= 0 {
		//response.ToErrorResponse(errcode.InvalidParams)
		return nil, errors.New("invalidParams")
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(fileType, file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		return nil, err
	}
	// 记录原始文件名，去除扩展名
	ext := path.Ext(fileHeader.Filename)
	fileInfo.OriginName = strings.TrimSuffix(fileHeader.Filename, ext)
	return fileInfo, nil
}
