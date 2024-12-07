package httptest

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

// handleError 是一个辅助函数，用于处理测试中的错误
func handleError(tb testing.TB, err error) {
	// 用Helper()方法标记该函数是一个helper函数，当测试失败时，输出的文件和行号将会是调用helper函数的位置，而不是在helper函数中调用的位置。
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

// TestHelloHandler 测试HelloHandler函数的HTTP处理
// go test -v ./httptest
func TestHelloHandler(t *testing.T) {
	// 在本地主机上启动一个TCP监听器
	ln, err := net.Listen("tcp", "localhost:8080")
	handleError(t, err)

	// 注册HelloHandler函数来处理"/hello"路径的请求
	http.HandleFunc("/hello", HelloHandler)

	// 在一个新的goroutine中启动HTTP服务
	// 如果是多个接口，可以使用http.NewServeMux()来注册多个接口
	//// 创建一个新的ServeMux
	//mux := http.NewServeMux()
	//
	//// 注册多个处理函数到不同的路径
	//mux.HandleFunc("/hello", HelloHandler)
	//mux.HandleFunc("/goodbye", GoodbyeHandler)
	//mux.HandleFunc("/api/data", DataHandler)
	//
	//// 在一个新的goroutine中启动HTTP服务
	//go func() {
	//	err := http.Serve(ln, mux)
	//	handleError(t, err)
	//}()

	go func() {
		err := http.Serve(ln, nil)
		handleError(t, err)
	}()

	// 向服务发送GET请求
	resp, err := http.Get("http://localhost:8080/hello")
	handleError(t, err)

	// 检查响应状态码是否为200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	// 确保在函数结束时关闭响应体
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	handleError(t, err)

	// 检查响应体内容是否符合预期
	if string(body) != "Hello, World!" {
		t.Errorf("got: %s, want: Hello, World!", string(body))
	}
}

// TestHelloHandlerMock 使用httptest包来模拟HTTP请求和响应
// go test -v -run TestHelloHandlerMock ./httptest
func TestHelloHandlerMock(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	w := httptest.NewRecorder()
	HelloHandler(w, req)
	body, err := io.ReadAll(w.Result().Body)
	handleError(t, err)
	if string(body) != "Hello, World!" {
		t.Errorf("got: %s, want: Hello, World!", string(body))
	}
}
