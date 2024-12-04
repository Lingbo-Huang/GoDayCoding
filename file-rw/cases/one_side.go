package cases

import (
	"io"
	"log"
	"os"
	"path"
)

// OneSideReadWriteToDest 按缓冲大小边读边写，把源文件夹的文件拷贝到目标文件夹
func OneSideReadWriteToDest() {
	list := getFiles(sourceDir)
	for _, srcFile := range list {
		_, fName := path.Split(srcFile)
		destFile := destDir + "one-side/" + fName
		OneSideReadWrite(srcFile, destFile)
	}
}

// OneSideReadWrite 分片读取文件内容并写入新文件
func OneSideReadWrite(srcName, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	// 边读边写要给O_APPEND权限，否则会覆盖原文件
	dest, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		if _, err := dest.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
}
