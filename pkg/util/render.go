package util

import "time"

// 将时间戳变为日期
func TimeStampToDate(timeStamp uint32) string {
	return time.Unix(int64(timeStamp), 0).Format("2006-01-02")
}

// 变量加减
func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
