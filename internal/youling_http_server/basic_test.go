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
	"testing"

	"github.com/pelletier/go-toml"
)

// @brief 测试readTemplates函数
func TestReadTemplates(t *testing.T) {
	// 1. 一般文件读取测试
	t.Run("一般测试：正常读取", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl1.toml", templates)
		// 断言结果
		if err != nil {
			t.Errorf("正常读取模板文件失败")
		}
	})
	// 2. 错误文件名读取测试
	t.Run("错误测试：错误文件名读取", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/error_file_name.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("读取错误文件名时没有报错")
		}
	})
	// 3. 文件读取边界测试
	t.Run("边界测试：空文件读取", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("", templates)
		// 断言结果
		if err == nil {
			t.Errorf("读取空文件名时没有报错")
		}
	})
	// 4. 键值读取测试
	// 4.1 从TOML配置文件中读取测试数据
	toml, err := toml.LoadFile("testdata/test_data.toml")
	if err != nil {
		t.Errorf("读取TOML测试文件失败！")
		return
	}
	// 4.2 正常读取测试
	t.Run("键值测试：正常读取", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl1.toml", templates)
		// 断言结果
		if err != nil {
			t.Errorf("读取模板文件失败")
		}
		value := toml.Get("readTemplates_key_value.value1").(string)
		if templates["general1"] != value {
			t.Errorf("读取的键值对不正确")
		}
		value = toml.Get("readTemplates_key_value.value2").(string)
		if templates["replacement"] != value {
			t.Errorf("读取的键值对不正确")
		}
	})
	// 4.3 错误表名读取测试
	t.Run("错误测试：文件中没有找到指定表名", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl3.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("文件中没有找到指定表名时没有报错")
		}
	})
	// 4.4 错误键名读取测试
	t.Run("错误测试：文件中第一个键名错误", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl2.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("文件中第一个键名错误时没有报错")
		}
	})
	t.Run("错误测试：文件中第二个键名错误", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl4.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("文件中第二个键名错误时没有报错")
		}
	})
	// 4.5 测试错误读取第二个键值指向的文件
	t.Run("错误测试：错误读取第二个键值指向的文件", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl5.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("错误读取第二个键值指向的文件时没有报错")
		}
	})
}
