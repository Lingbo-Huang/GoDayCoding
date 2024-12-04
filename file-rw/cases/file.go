package cases

import (
	"log"
	"os"
	"strings"
)

// 源文件目录
const sourceDir = "source-files/"

// 目标文件目录
const destDir = "dest-files/"

// 获取源文件目录下的所有文件，返回每个文件的相对路径构成的列表
func getFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		// 获取完整路径
		fullName := strings.Trim(dir, "/") + "/" + f.Name()
		list = append(list, fullName)
	}
	return list
}
