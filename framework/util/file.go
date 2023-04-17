package util

import (
	"os"
	"path/filepath"
	"strings"
)

// Exists 判断所给路径文件或文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		// 看看是不是因为文件已经存在导致的error
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsHiddenDirectory 判断是否是隐藏路径
func IsHiddenDirectory(path string) bool {
	return len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".")
}
