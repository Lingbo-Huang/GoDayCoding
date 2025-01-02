package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 接受请求并处理请求
	r.POST("/login", loginHandler)

	// 启动服务器
	r.Run(":8080")
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// loginHandler 处理登录请求
func loginHandler(c *gin.Context) {
	// 获取请求数据
	var reqData Login
	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求参数错误",
		})
		return
	}
	fmt.Printf("reqData:%#v\n", reqData)
	c.JSON(http.StatusOK, reqData)
	// 对数据进行校验
	
	// 返回响应
}
