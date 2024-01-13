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
	"net/http"
	"os"
	"strings"
	"sunflower/pkg/youling_string"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

type Template struct {
	Name string
	File string
}

type Tpl struct {
	Templates []Template
}

type PH struct {
	Name     string
	Contents []string
}

type PHs struct {
	Place_holder []PH
}

// @brief 读取模板文件
//  @param file 模板文件
//  @param tpl 模板哈希表
//  @return 成功：nil，失败：错误信息
func readTemplates(f string, tpl map[string]string) error {
	// 读取模板列表配置文件
	conf, err := toml.LoadFile(f)
	if err != nil {
		return errors.New("载入模板配置文件 " + f + " 失败。")
	}
	// 解析模板列表配置文件
	if fmt.Sprintf("%T", conf.Get("templates")) != "[]*toml.Tree" {
		return errors.New("模板配置文件 " + f + " 中找不到 templates 表头。")
	}
	tp1 := conf.Get("templates").([]*toml.Tree)
	var tp Tpl
	for _, v := range tp1 {
		if fmt.Sprintf("%T", v.Get("name")) != "string" {
			return errors.New("模板配置文件 " + f + " 中 name 字段找不到。")
		}
		name := v.Get("name").(string)
		if fmt.Sprintf("%T", v.Get("file")) != "string" {
			return errors.New("模板配置文件 " + f + " 中 file 字段找不到。")
		}
		file := v.Get("file").(string)
		tp.Templates = append(tp.Templates, Template{Name: name, File: file})
	}
	// 把模板列表中的名称与实际的文件内容逐一读入哈希表中
	for _, t := range tp.Templates {
		contents, err := os.ReadFile(t.File)
		if err != nil {
			return errors.New("读取模板文件 " + t.File + " 失败。")
		}
		tpl[t.Name] = string(contents)
	}
	return nil
}

// @brief 读取占位符列表
//  @param file 占位符配置文件
//  @param ph 占位符列表
//  @return 成功：nil，失败：错误信息
func readPlaceHolderList(file string, ph map[string][]string) error {
	// 读取占位符列表配置文件
	conf, err := toml.LoadFile(file)
	if err != nil {
		return errors.New("载入占位符文件 " + file + " 时发生错误：")
	}
	if fmt.Sprintf("%T", conf.Get("place_holder")) != "[]*toml.Tree" {
		return errors.New("占位符配置文件 " + file + " 中找不到 place_holder 表头。")
	}
	phs := conf.Get("place_holder").([]*toml.Tree)
	var ph1 PHs
	for _, v := range phs {
		if fmt.Sprintf("%T", v.Get("name")) != "string" {
			return errors.New("占位符配置文件 " + file + " 中找不到 name 字段。")
		}
		name := v.Get("name").(string)
		if fmt.Sprintf("%T", v.Get("contents")) != "[]interface {}" {
			return errors.New("占位符配置文件 " + file + " 中找不到 contents 字段。")
		}
		contentsTemp := v.Get("contents").([]interface{})
		var contents []string
		for _, w := range contentsTemp {
			contents = append(contents, w.(string))
		}
		ph1.Place_holder = append(ph1.Place_holder, PH{Name: name,
			Contents: contents})
	}
	for _, p := range ph1.Place_holder {
		ph[p.Name] = p.Contents
	}
	return nil
}

// @brief 替换模板中的占位符
//  @param r 路由表
//  @param placeHolder 占位符
//  @param templates   模板
//  @return 页面内容
func replacePlaceHolder(r Router, placeHolder []string,
	template string, replacement string) string {
	for _, p := range placeHolder {
		bound := "<!--" + r.Function + "." + p + "-->"
		template = strings.Replace(template, "<!--{{."+p+"}}-->",
			youling_string.ReadBetween(replacement, bound, bound), -1)
	}
	return template
}

// @brief 设置 http server 参数
//  @param srv http服务器
//  @return 成功：nil，失败：错误信息
func setServer(file string, srv *http.Server, r *gin.Engine) error {
	// 从TOML配置文件中读取http服务器参数
	config, err := toml.LoadFile(file)
	if err != nil {
		return errors.New("载入服务器参数文件 server_config.toml 时发生错误。")
	}
	// 服务器参数设置
	srv.Addr = config.Get("server.address").(string) + ":" +
		config.Get("server.port").(string)
	srv.Handler = r
	srv.ReadTimeout = time.Duration(config.Get("server.ReadTimeout").(int64)) *
		time.Second
	srv.WriteTimeout = time.Duration(config.Get("server.WriteTimeout").(int64)) *
		time.Second
	return nil
}
