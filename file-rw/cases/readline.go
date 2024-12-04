package cases

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const README = "README.md"

// ReadLine 一次读取，按行打印文件内容
func ReadLine1() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 一次性读取文件内容
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// 按行切分
	list := strings.Split(string(bytes), "\n")
	// 逐行打印
	for _, line := range list {
		log.Println(line)
	}
}

// ReadLine2 利用 bufio 按行读取文件内容
func ReadLine2() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 读取一行内容, 包含换行符
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Print(line)
	}
}

// ReadLine3 利用 scanner 按行读取文件内容
func ReadLine3() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 获取一行内容, 不包含换行符
		line := scanner.Text()
		fmt.Println(line)
	}
}
