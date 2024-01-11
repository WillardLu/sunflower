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
	"fmt"
	"os"
	"sunflower/pkg/youling_go_basic"

	"github.com/pelletier/go-toml"
)

// @brief 读取模板文件
//  @param tpl 模板哈希表
//  @return 成功 0，失败 -1
func readTemplates(tpl map[string]string) int {
	// 读取模板列表配置文件
	tree, err := toml.LoadFile("config/templates_list.toml")
	if err != nil {
		fmt.Println("载入 templates_list.toml 时发生错误：", err)
		return -1
	}
	msg := youling_go_basic.CheckDataType(tree.Get("templates"), "[]*toml.Tree")
	if msg != "" {
		fmt.Println("读取模板文件中的数据时发生错误：", msg)
		return -1
	}
	template := tree.Get("templates").([]*toml.Tree)
	// 把模板列表中的名称与实际的文件内容逐一读入哈希表中
	for _, t := range template {
		msg1 := youling_go_basic.CheckDataType(t.Get("name"), "string")
		msg2 := youling_go_basic.CheckDataType(t.Get("file"), "string")
		if msg1 != "" || msg2 != "" {
			fmt.Println("读取模板文件中的数据时发生错误：")
			fmt.Println(msg1)
			fmt.Println(msg2)
			return -1
		}
		contents, err := os.ReadFile(t.Get("file").(string))
		if err != nil {
			fmt.Println("读取模板文件“"+t.Get("file").(string)+"”时发生错误：", err)
			return -1
		}
		tpl[t.Get("name").(string)] = string(contents)
	}
	return 0
}
