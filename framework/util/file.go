package util

import (
	"io"
	"net/http"
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

// SubDir 输出所有子目录和目录名
func SubDir(folder string) ([]string, error) {
	subs, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, sub := range subs {
		if sub.IsDir() {
			ret = append(ret, sub.Name())
		}
	}
	return ret, nil
}

// DownloadFile 下载的时候就写, 而不是把整个文件全部放入内存
func DownloadFile(filepath string, url string) error {
	// 先拿到数据
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	out, err1 := os.Create(filepath)
	if err1 != nil {
		return err1
	}
	defer out.Close()

	// 将数据的body写入文件中
	_, err1 = io.Copy(out, resp.Body)
	return err1
}
