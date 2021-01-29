// common_test.go
// 辅助测试函数

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpMessageList []message

// 在执行测试函数之前进行setup
func TestMain(m *testing.M) {
	//设置gin为测试模式
	gin.SetMode(gin.TestMode)

	// 运行其他测试
	os.Exit(m.Run())
}

// 在测试期间创建getRouter函数
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// 处理请求并测试其响应的函数
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 创建service并处理上述请求
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// 这个函数用于将主列表存储到临时列表中进行测试
func saveLists() {
	tmpMessageList = messageList
}

// 此函数用于从临时列表恢复主列表
func restoreLists() {
	messageList = tmpMessageList
}
