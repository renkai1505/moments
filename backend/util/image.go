package util

import (
	"strings"
)

// GetThumbUrl 获取缩略图URL
func GetThumbUrl(originalUrl string) string {
	if originalUrl == "" {
		return ""
	}
	
	// 如果是本地图片，尝试添加缩略图后缀
	if strings.HasPrefix(originalUrl, "/upload/") {
		// 获取文件扩展名
		lastDot := strings.LastIndex(originalUrl, ".")
		if lastDot > 0 {
			ext := originalUrl[lastDot:]
			nameWithoutExt := originalUrl[:lastDot]
			return nameWithoutExt + "_thumb" + ext
		}
	}
	
	// 默认返回原图
	return originalUrl
}