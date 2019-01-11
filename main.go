package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)
	// 绑定端口，然后启动应用
	engine.Run(":8000")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
*/
func WebRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}
//
//func main(){
//	fmt.Println("hello")
//}