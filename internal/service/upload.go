package service

import (
	"errors"
	"goblog/pkg/upload"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// 获取文件基本信息
	fileName := upload.GetFileName(fileHeader.Filename)
	// 对文件进行检查
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	uploadSavePath := upload.GetSavePath(fileType)
	dst := filepath.Join(uploadSavePath, fileName)
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := upload.GetAccessUrl(fileType, fileName)
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
