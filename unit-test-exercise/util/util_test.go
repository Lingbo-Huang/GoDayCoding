package util

import (
	"testing"
)

var commonTestData []commonStruct

type commonStruct struct {
	Group         string
	input         string
	expectSize    int64
	expectSizeStr string
}

func TestParseSize1(t *testing.T) {
	testData := commonTestData
	for _, data := range testData {
		size, sizeStr := ParseSize1(data.input)
		if size != data.expectSize || sizeStr != data.expectSizeStr {
			t.Errorf("测试结果不符合预期：ParseSize1(%s) = %d, %s; want %d, %s", data.input, size, sizeStr, data.expectSize, data.expectSizeStr)
		}
	}
}

// control + G 可以继续选中下一个相同的单词
// control + command + G 全选
// TestParseSizeSub 测试不同组别的ParseSize1函数
// go test -v -short ./util
func TestParseSizeSub(t *testing.T) {
	// 如果是short模式，跳过测试
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// 创建一个map来存储按组分类的测试数据
	testData := make(map[string][]commonStruct)
	// 遍历commonTestData，将测试数据按组别分类
	for _, item := range commonTestData {
		group := item.Group
		// 如果该组别还没有对应的切片，则创建一个
		if _, ok := testData[group]; !ok {
			testData[group] = make([]commonStruct, 0)
		}
		// 将当前测试数据添加到对应的组别中
		testData[group] = append(testData[group], item)
	}

	// 遍历按组分类的测试数据
	for k, _ := range testData {
		// 为每个组别创建一个子测试
		t.Run(k, func(t *testing.T) {
			// 并行测试
			t.Parallel()
			// 遍历当前组别的所有测试用例
			for _, data := range testData[k] {
				size, sizeStr := ParseSize1(data.input)
				if size != data.expectSize || sizeStr != data.expectSizeStr {
					t.Errorf("测试结果不符合预期：ParseSize1(%s) = %d, %s; want %d, %s", data.input, size, sizeStr, data.expectSize, data.expectSizeStr)
				}
			}
		})
	}
}

// go test -v -run ^$ -fuzz . -fuzztime 5 ./util/
// ^$匹配空字符串，所以-run ^$表示不运行任何功能测试用例，-fuzz .表示运行所有fuzz测试用例
// -fuzztime 5表示运行5秒， -fuzztime 5x 表示运行5次
// 运行fuzz测试时，会在当前目录下生成一个fuzz文件夹，里面存放了fuzz测试的语料库
func FuzzParseSize1(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		size, str := ParseSize1(a)
		if size == 0 && str == "" {
			t.Errorf("ParseSize1(%s) = 0, \"\"; want non-zero", a)
		}
	})
}

// go test -v -run ^$ -bench . -benchtime 10s -benchmem ./util
func BenchmarkParseSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseSize1("10gb")
	}
}

func BenchmarkParseSize1Sub(b *testing.B) {
	testData := make(map[string][]commonStruct)
	for _, item := range commonTestData {
		group := item.Group
		if _, ok := testData[group]; !ok {
			testData[group] = make([]commonStruct, 0)
		}
		testData[group] = append(testData[group], item)
	}

	b.ResetTimer()
	for k, _ := range testData {
		b.Run(k, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, data := range testData[k] {
					b.StartTimer()
					ParseSize1(data.input)
					b.StopTimer()
				}
			}
		})
	}
}

func BenchmarkParseSize1Parallel(b *testing.B) {
	// 每个核20个goroutine并行执行
	b.SetParallelism(20)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ParseSize1("10gb")
		}
	})
}

// 测试用例执行前先进入此函数
func TestMain(m *testing.M) {
	initCommonData()
	m.Run()
}

func initCommonData() {
	commonTestData = []commonStruct{
		{Group: "B", input: "1b", expectSize: B, expectSizeStr: "1B"},
		{Group: "B", input: "100B", expectSize: 100 * B, expectSizeStr: "100B"},
		{Group: "KB", input: "1kb", expectSize: KB, expectSizeStr: "1KB"},
		{Group: "KB", input: "10kb", expectSize: 10 * KB, expectSizeStr: "10KB"},
		{Group: "MB", input: "1mb", expectSize: MB, expectSizeStr: "1MB"},
		{Group: "MB", input: "1MB", expectSize: MB, expectSizeStr: "1MB"},
		{Group: "GB", input: "1gb", expectSize: GB, expectSizeStr: "1GB"},
		{Group: "GB", input: "1Gb", expectSize: GB, expectSizeStr: "1GB"},
		{Group: "TB", input: "2tb", expectSize: 2 * TB, expectSizeStr: "2TB"},
		{Group: "TB", input: "1TB", expectSize: TB, expectSizeStr: "1TB"},
		{Group: "PB", input: "1pb", expectSize: PB, expectSizeStr: "1PB"},
		{Group: "PB", input: "12PB", expectSize: 12 * PB, expectSizeStr: "12PB"},
		{Group: "unknown", input: "1k", expectSize: 100 * MB, expectSizeStr: "100MB"},
		{Group: "unknown", input: "1g", expectSize: 100 * MB, expectSizeStr: "100MB"},
	}
}
