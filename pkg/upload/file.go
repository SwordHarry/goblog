package upload

import (
	"goblog/global"
	"goblog/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 文件上传服务
type FileType int

const (
	TypeImage FileType = iota + 1
	TypeMd
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath(fileType FileType) string {
	switch fileType {
	case TypeImage:
		return global.AppSetting.UploadImageSavePath
	case TypeMd:
		return global.AppSetting.UploadMDSavePath
	}
	return global.AppSetting.UploadSavePath
}

// 对文件的检查

// 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	// err 是否为 oserror.ErrNotExist
	return os.IsNotExist(err)
}

// 检查后缀是否合理
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	case TypeMd:
		for _, allowExt := range global.AppSetting.UploadMDAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

// 检查文件大小是否超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	case TypeMd:
		if size >= global.AppSetting.UploadMDMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 文件相关操作
// 创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	return os.MkdirAll(dst, perm)
}

// 保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

func GetAccessUrl(fileType FileType, fileName string) string {
	switch fileType {
	case TypeImage:
		return filepath.Join(global.AppSetting.UploadImageServerUrl, fileName)
	case TypeMd:
		return filepath.Join(global.AppSetting.UploadMDServerUrl, fileName)
	}
	return ""
}
