package cases

import (
	"io"
	"os"
	"path"
)

// CopyDirToDir 拷贝文件夹到文件夹
func CopyDirToDir() {
	list := getFiles(sourceDir)
	for _, srcFile := range list {
		// 拿到文件名
		_, fName := path.Split(srcFile)
		// 拼接目标文件路径
		destFile := destDir + "copy/" + fName
		// 拷贝文件
		CopyFile(srcFile, destFile)
	}
}

// CopyFile 拷贝文件(拷贝内容从源文件到目标文件)，返回的第一个参数是拷贝的字节数
func CopyFile(srcFile string, destFile string) (int64, error) {
	// 打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	// 打开目标文件
	// Perm: 0644 表示所有者有读写权限，同组用户和其他用户只有读权限
	dst, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	// 拷贝文件内容
	return io.Copy(dst, src)
}
