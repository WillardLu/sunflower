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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

// @brief 设置路由
//  @param gin router
//  @return 0 成功，-1 失败
func setupRouter(router *gin.Engine) error {
	// 1. 设置静态文件位置
	router.StaticFS("/css", http.Dir("./static/css"))
	router.StaticFS("/images", http.Dir("./static/images"))
	router.StaticFS("/js", http.Dir("./static/js"))

	// 2.读取模板文件内容
	// 通过这样的方式把模板文件内容读入内存，以减少磁盘读取
	templates := make(map[string]string)
	err := readTemplates("config/templates_list.toml", templates)
	if err != nil {
		return err
	}
	fmt.Println("\n载入模板文件完成……")

	// 3. 读取路由表
	tree, err := toml.LoadFile("config/routing.toml")
	if err != nil {
		return errors.New("载入 routing.toml 时发生错误：" + err.Error())
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
		return errors.New("载入 place_holder.toml 时发生错误：" + err.Error())
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
		readTemplates("config/templates_list.toml", templates)
		fmt.Println("\n刷新模板文件完成……")
	})
	// 获取来自网页提交的内容（演示用）
	router.POST("/submit-data", func(c *gin.Context) {
		handleData(c)
	})
	fmt.Println("\n路由设置完成；")
	return nil
}

// @brief 处理接收到的数据
//  @param c 上下文
func handleData(c *gin.Context) error {
	// 获取来自网页提交的内容
	str, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return errors.New("读取数据时发生错误：" + err.Error())
	}
	fmt.Println(string(str))
	c.String(http.StatusCreated, string(str))
	return nil
}
