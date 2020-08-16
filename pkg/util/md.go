package util

import (
	"github.com/russross/blackfriday/v2"
	"goblog/global"
	"io/ioutil"
	"regexp"
	"strings"
)

// 获取 markdown 文件并转移成 html
func GetMd(fileName string) ([]byte, error) {
	filePath := global.AppSetting.UploadMDSavePath + "/" + fileName
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return blackfriday.Run(b), nil
}

// 将 html 内容转化为普通文本
func Html2Text(html string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllString(html, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s+")
	html = re.ReplaceAllString(html, " ")
	// 去除特殊字符
	re, _ = regexp.Compile("&[A-Za-z]+?;")
	html = re.ReplaceAllString(html, "")
	return strings.TrimSpace(html)
}
