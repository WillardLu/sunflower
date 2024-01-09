package sunflowerwebserver

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// 生成主页
//  c - 指向gin.Context的指针
//  page - 模板文件内容
func homePage(c *gin.Context, template string) {
	var msg []byte
	str := "<a href=\"#\" onclick=\"location.reload()\">主页</a>"
	msg = []byte(strings.Replace(string(template), "{{.sub-dir}}", str, -1))
	c.Writer.Write(msg)
	// c.File("./static/index.html") // 这种方式可以直接读取文件内容并显示在浏览器上
}

// 生成任务列表页
//  c - 指向gin.Context的指针
//  page - 模板文件内容
func taskList(c *gin.Context, template string) {
	var msg []byte
	str := "<a href=\"\">任务列表</a>"
	msg = []byte(strings.Replace(template, "{{.sub-dir}}", str, -1))
	c.Writer.Write(msg)
}

// 设置路由
//  函数名首字母大写才能被其他文件调用
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 设置CSS样式的位置
	router.StaticFS("/css", http.Dir("./static/css"))
	// 设置静态图片的位置
	router.StaticFS("/images", http.Dir("./static/images"))

	// 读取模板文件内容
	// 通过这样的方式把模板文件内容读入内存，以减少磁盘读取
	template, error := os.ReadFile("./templates/general-page.html")
	if error != nil {
		log.Println("Error reading file:", error)
		return nil
	}

	// 这里不能直接调用多参数的函数，
	// 需要使用func(c *gin.Context)作为中转来调用多参数的函数
	router.GET("/", func(c *gin.Context) {
		homePage(c, string(template))
	})
	router.GET("/task-list", func(c *gin.Context) {
		taskList(c, string(template))
	})

	return router
}
