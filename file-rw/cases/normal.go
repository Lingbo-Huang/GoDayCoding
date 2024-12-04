package cases

import (
	"log"
	"os"
	"path"
)

// ReadWriteFiles reads files from sourceDir and writes them to destDir
func ReadWriteFiles() {
	list := getFiles(sourceDir)
	for _, srcFile := range list {
		bytes, err := os.ReadFile(srcFile)
		if err != nil {
			log.Fatal(err)
		}
		_, fName := path.Split(srcFile)
		destFile := destDir + "normal/" + fName
		err = os.WriteFile(destFile, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
