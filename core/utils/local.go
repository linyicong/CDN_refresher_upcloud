package utils

import (
	"os"
	"path/filepath"
)

func Local(pwd string) (files []string) {
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path[len(pwd):])
		}
		return nil
	})
	return
}