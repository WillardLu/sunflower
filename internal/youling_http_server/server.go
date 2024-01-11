// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
// @author 陆巍
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_http_server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"sunflower/pkg/youling_string"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

// 路由
type routers struct {
	from  string
	to    string
	TPL   string
	REPLC string
}

// @brief 创建http server
func CreateHttpServer() {
	// 1. 设置运行模式
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 2. 设置路由
	if setupRouter(router) == -1 {
		return
	}
	srv := &http.Server{}
	// 3. 设置服务器参数
	if setServer(srv, router) == -1 {
		return
	}
	// 4. 监听请求
	// 声明一个匿名函数，并创建一个goroutine（有的翻译为协程）
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 关闭（或重启）服务
	// 1）创建通道，用来接收信号
	quit := make(chan os.Signal, 1)
	// 2）监听和捕获信号
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("\n>> 开始关闭 http server……")
	// 3）创建一个子节点的context,5秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("\n>> http server 关闭时出错：", err)
	}
	select {
	case <-ctx.Done():
	default:
	}
	log.Println("\n>> http server 退出。")
	return
}

// 设置 http server 参数
//  srv http服务器
//  返回值：0 成功，-1 失败
func setServer(srv *http.Server, r *gin.Engine) int {
	// 从TOML配置文件中读取http服务器参数
	config, err := toml.LoadFile("config/server_config.toml")
	if err != nil {
		fmt.Println("载入 server_config.toml 文件时发生错误：", err)
		return -1
	}
	// 服务器参数设置
	srv.Addr = config.Get("server.address").(string) + ":" +
		config.Get("server.port").(string)
	srv.Handler = r
	srv.ReadTimeout = time.Duration(config.Get("server.ReadTimeout").(int64)) *
		time.Second
	srv.WriteTimeout = time.Duration(config.Get("server.WriteTimeout").(int64)) *
		time.Second
	return 0
}

// 设置路由
//  gin router
//  返回值：0 成功，-1 失败
func setupRouter(router *gin.Engine) int {
	// 1. 设置静态文件位置
	router.StaticFS("/css", http.Dir("./static/css"))
	router.StaticFS("/images", http.Dir("./static/images"))
	router.StaticFS("/js", http.Dir("./static/js"))

	// 2.读取模板文件内容
	// 通过这样的方式把模板文件内容读入内存，以减少磁盘读取
	templates := make(map[string]string)
	if readTemplates(templates) == -1 {
		return -1
	}
	fmt.Println("\n载入模板文件完成……")

	// 3. 读取路由表
	tree, err := toml.LoadFile("config/routing.toml")
	if err != nil {
		fmt.Println("载入 routing.toml 时发生错误：", err)
		return -1
	}
	route := tree.Get("routing").([]*toml.Tree)
	var r []routers
	for _, route := range route {
		var r1 routers
		r1.from = route.Get("from").(string)
		r1.to = route.Get("to").(string)
		r1.TPL = route.Get("template").(string)
		r1.REPLC = route.Get("replacement").(string)
		r = append(r, r1)
	}
	fmt.Println("\n载入路由完成……")

	// 4. 读取占位符表
	placeHolders, err := toml.LoadFile("config/place_holder.toml")
	if err != nil {
		fmt.Println("载入 place_holder.toml 时发生错误：", err)
		return -1
	}
	fmt.Println("\n读取占位符表完成……")

	// 读取路由表
	for _, r := range r {
		r1 := r // 这里不能直接把 r 交给下面去处理，否则传过去的 r 始终会指向最后一项
		// 设置路由
		// 这里不能直接调用多参数的函数，
		// 需要使用func(c *gin.Context)作为中转来调用多参数的函数
		router.GET(r1.from, func(c *gin.Context) {
			replacePlaceHolder(c, r1, placeHolders.Get(r1.to).([]interface{}),
				templates)
		})
	}
	// 设置刷新模板内容的路由，以方便在不用重启路由的情况下进行一些调试维护
	router.GET("/refresh", func(c *gin.Context) {
		readTemplates(templates)
		fmt.Println("\n刷新模板文件完成……")
	})
	// 获取来自网页提交的内容（演示用）
	router.POST("/submit-data", func(c *gin.Context) {
		handleData(c)
	})
	fmt.Println("\n路由设置完成；")
	return 0
}

// 替换模板中的占位符
//  c           上下文环境
//  r           路由表
//  placeHolder 占位符
//  templates   模板
func replacePlaceHolder(c *gin.Context, r routers, placeHolder []interface{},
	templates map[string]string) {
	str := templates[r.TPL]
	r.REPLC = templates[r.REPLC]

	for _, p := range placeHolder {
		bound := "<!--" + r.to + "." + p.(string) + "-->"
		str = strings.Replace(str, "<!--{{."+p.(string)+"}}-->",
			youling_string.ReadBetween(r.REPLC, bound, bound), -1)
	}
	c.Writer.Write([]byte(str))
	return
}

// 处理接收到的数据
func handleData(c *gin.Context) {
	// 获取来自网页提交的内容
	str, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("读取数据时发生错误：", err)
		return
	}
	fmt.Println(string(str))
	c.String(http.StatusCreated, string(str))
	return
}
