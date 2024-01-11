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
	"os"
	"sunflower/pkg/youling_go_basic"

	"github.com/pelletier/go-toml"
)

// @brief 读取模板文件
//  @param file 模板文件
//  @param tpl 模板哈希表
//  @return 成功：nil，失败：错误信息
func readTemplates(file string, tpl map[string]string) error {
	// 读取模板列表配置文件
	tree, err := toml.LoadFile(file)
	if err != nil {
		return errors.New("载入 " + file + " 文件时发生错误：" + err.Error())
	}
	msg := youling_go_basic.CheckDataType(tree.Get("templates"), "[]*toml.Tree")
	if msg != nil {
		return errors.New("读取模板文件中的数据时发生错误：" + msg.Error())
	}
	template := tree.Get("templates").([]*toml.Tree)
	// 把模板列表中的名称与实际的文件内容逐一读入哈希表中
	for _, t := range template {
		msg := youling_go_basic.CheckDataType(t.Get("name"), "string")
		if msg != nil {
			return errors.New("读取模板文件中的数据时发生错误\n" + msg.Error())
		}
		msg = youling_go_basic.CheckDataType(t.Get("file"), "string")
		if msg != nil {
			return errors.New("读取模板文件中的数据时发生错误\n" + msg.Error() + "\n" +
				msg.Error())
		}
		contents, err := os.ReadFile(t.Get("file").(string))
		if err != nil {
			return errors.New("读取模板文件“" + t.Get("file").(string) +
				"”时发生错误：" + err.Error())
		}
		tpl[t.Get("name").(string)] = string(contents)
	}
	return nil
}
