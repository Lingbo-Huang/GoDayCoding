package util

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

const defaultNum = 100

func ParseSize1(size string) (int64, string) {
	time.Sleep(time.Nanosecond * 500)
	re, _ := regexp.Compile("[0-9]+")
	// 利用正则拿到字符串去掉其中所有数字剩余的部分（这里就是单位）
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	// 利用strings.Replace去掉unit部分，拿到字符串里数字部分
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num * B
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}
	if num == 0 {
		log.Println("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
		// 试验fuzz测试时，此处会panic
		//panic("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
		num = defaultNum
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizeStr
}

func ParseSize(size string) (int64, string, error) {
	time.Sleep(time.Nanosecond * 500)

	// 去除空白字符并转换为大写
	size = strings.ToUpper(strings.TrimSpace(size))

	// 使用正则表达式分离数字和单位
	// (?:)
	re := regexp.MustCompile(`^(\d+)\s*(?:([KMGTP]B?)|B)?$`)
	matches := re.FindStringSubmatch(size)

	if len(matches) < 2 {
		return 0, "", fmt.Errorf("无效的大小格式: %s", size)
	}

	// 解析数字部分
	num, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, "", fmt.Errorf("无法解析数字: %v", err)
	}

	// 处理单位
	unit := "B"
	if len(matches) == 3 && matches[2] != "" {
		unit = matches[2]
		if !strings.HasSuffix(unit, "B") {
			unit += "B"
		}
	}

	// 计算字节数
	var byteNum int64
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		return 0, "", fmt.Errorf("不支持的单位: %s", unit)
	}

	sizeStr := fmt.Sprintf("%d%s", num, unit)
	return byteNum, sizeStr, nil
}
