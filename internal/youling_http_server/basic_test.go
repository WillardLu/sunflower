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
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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
	t.Run("错误测试：错误表名", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl3.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("表名错误的情况下却没有报错")
		}
	})
	// 4.4 错误键名读取测试
	t.Run("错误测试：第一个键名错误", func(t *testing.T) {
		templates := make(map[string]string)
		// 调用被测试的函数
		err := readTemplates("testdata/tpl2.toml", templates)
		// 断言结果
		if err == nil {
			t.Errorf("文件中第一个键名错误时没有报错")
		}
	})
	t.Run("错误测试：第二个键名错误", func(t *testing.T) {
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

// @brief 测试readPlaceHolderList函数
func TestReadPlaceHolderList(t *testing.T) {
	// 1. 正常测试
	t.Run("正常测试", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holder.toml", ph)
		// 断言结果
		if err != nil {
			t.Errorf("测试失败：%v", err)
		}
	})
	// 2. 错误测试
	// 2.1 文件名错误
	t.Run("错误测试：文件名错误", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holders.toml", ph)
		// 断言结果
		if err == nil {
			t.Errorf("文件名错误，但没有报错。")
		}
	})
	// 2.2 表头错误
	t.Run("错误测试：place_holder 表头错误", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holder1.toml", ph)
		// 断言结果
		if err == nil {
			t.Errorf("place_holder 表头错误，但没有报错。")
		}
	})
	// 2.3 字段错误
	t.Run("错误测试：name 字段错误", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holder2.toml", ph)
		// 断言结果
		if err == nil {
			t.Errorf("name 字段错误，却没有报错。")
		}
	})
	t.Run("错误测试：contents 字段错误", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holder3.toml", ph)
		// 断言结果
		if err == nil {
			t.Errorf("contents 字段错误，却没有报错。")
		}
	})
	// 3. 内容测试
	t.Run("内容测试", func(t *testing.T) {
		ph := make(map[string][]string)
		// 调用被测试的函数
		err := readPlaceHolderList("testdata/place_holder.toml", ph)
		// 断言结果
		if err != nil {
			t.Errorf("测试失败：%v", err)
		}
		// 断言内容
		if ph["homepage"][0] != "subdir" {
			t.Errorf("homepage 内容错误。")
		}
		if ph["homepage"][1] != "funcmenu" {
			t.Errorf("homepage 内容错误。")
		}
		if ph["task-list"][0] != "subdir" {
			t.Errorf("homepage 内容错误。")
		}
		if ph["task-list"][2] != "contents" {
			t.Errorf("homepage 内容错误。")
		}
	})
}

// @brief 测试 replacePlaceHolder 函数
func TestReplacePlaceHolder(t *testing.T) {
	// 定义输入参数
	r := Router{
		"template",
		"/",
		"homepage",
		"general1",
		"replacement",
		"",
	}
	placeHolder := []string{"subdir", "funcmenu"}
	template, error := os.ReadFile("testdata/general1.html")
	if error != nil {
		t.Errorf("读取测试用模板文件失败：%v", error)
	}
	replacement, error := os.ReadFile("testdata/replacement.html")
	if error != nil {
		t.Errorf("读取测试用替换文件失败：%v", error)
	}
	page, error := os.ReadFile("testdata/page.html")
	if error != nil {
		t.Errorf("读取测试用页面文件失败：%v", error)
	}
	// 调用被测试的函数
	str := replacePlaceHolder(r, placeHolder, string(template),
		string(replacement))
	t.Run("内容测试", func(t *testing.T) {
		if string(page) != str {
			t.Errorf("替换后的内容与预期不一致。")
		}
	})
}

// @brief 测试 setServer 函数
func TestSetServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	srv := &http.Server{}
	t.Run("正常测试", func(t *testing.T) {
		err := setServer("testdata/server_config.toml", srv, router)
		if err != nil {
			t.Errorf("设置服务器失败：%v", err)
		}
	})
	t.Run("错误测试：参数文件名错误", func(t *testing.T) {
		err := setServer("testdata/server_config_error.toml", srv, router)
		if err == nil {
			t.Errorf("参数文件名不正确，但没有报错。")
		}
	})
}
