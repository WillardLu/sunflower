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
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
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
	err := setupRouter(router)
	if err != nil {
		return
	}
	srv := &http.Server{}

	// 3. 设置服务器参数
	if setServer(srv, router) != nil {
		return
	}

	// 4. 监听请求
	// 声明一个匿名函数，并创建一个goroutine（有的翻译为协程）
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 5.关闭（或重启）服务
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
