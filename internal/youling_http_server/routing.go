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

// @brief 路由配置结构
type Router struct {
	Type        string // 类型
	Path        string // 路径
	Function    string // 函数
	Template    string // 模板
	Replacement string // 替换
	Dir         string // 目录
}

type Routers struct {
	Routing []Router // 路由列表
}

// @brief 设置路由
//  @param gin router
//  @return 0 成功，-1 失败
func setupRouter(router *gin.Engine) error {
	// 1.读取模板文件内容
	// 通过这样的方式把模板文件内容读入内存，以减少磁盘读取
	templates := make(map[string]string)
	err := readTemplates("config/templates_list.toml", templates)
	if err != nil {
		return err
	}
	// 2. 读取路由配置表
	var router_list Routers
	err = readRouting(&router_list)
	if err != nil {
		return err
	}
	// 3. 读取占位符列表
	placeHolder := make(map[string][]string)
	err = readPlaceHolderList("config/place_holder.toml", placeHolder)
	if err != nil {
		return err
	}
	// 4. 设置路由
	// 按照路由配置表设置路由
	// 这里不能直接调用多参数的函数，
	// 需要使用func(c *gin.Context)作为中转来调用多参数的函数
	for _, r := range router_list.Routing {
		r1 := r // 这里不能直接把 r 交给下面去处理，否则传过去的 r 始终会指向最后一项
		switch r1.Type {
		case "static":
			router.Static(r1.Path, r1.Dir)
		case "template":
			router.GET(r1.Path, func(c *gin.Context) {
				str := replacePlaceHolder(r1, placeHolder[r1.Function],
					templates[r1.Template], templates[r1.Replacement])
				c.Writer.Write([]byte(str))
			})
		case "function":
			switch r1.Function {
			case "refreshTemplates":
				router.GET(r1.Path, func(c *gin.Context) {
					refreshTemplates(templates)
				})
			case "handleData":
				router.POST(r1.Path, func(c *gin.Context) {
					handleData(c)
				})
			default:
				return errors.New(fmt.Sprintf("未知路由功能：%s", r1.Function))
			}
		default:
			return errors.New(fmt.Sprintf("未知路由类型：%s", r1.Type))
		}
	}
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

// @brief 读取路由配置表
func readRouting(r *Routers) error {
	conf, err := toml.LoadFile("config/routing.toml")
	if err != nil {
		return errors.New("载入 routing.toml 时发生错误：" + err.Error())
	}
	err = conf.Unmarshal(r)
	if err != nil {
		return errors.New("解析 routing.toml 时发生错误：" + err.Error())
	}
	return nil
}

// @brief 刷新模板文件
//  @param tpl 模板文件哈希表
//  @return 成功：nil，失败：错误信息
func refreshTemplates(tpl map[string]string) error {
	err := readTemplates("config/templates_list.toml", tpl)
	if err != nil {
		return err
	}
	fmt.Println("\n刷新模板文件完成……")
	return nil
}
