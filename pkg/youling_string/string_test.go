// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
// @author 陆巍
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_string

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

// @brief 测试ReadBetween函数
func TestReadBetween(t *testing.T) {
	type args struct {
		s     string
		start string
		end   string
	}
	type tests struct {
		name string
		args args
		want string
	}
	// 从TOML文件中读取测试数据
	datas, err := toml.LoadFile("testdata/test_data.toml")
	if err != nil {
		t.Error("\n载入测试数据文件 test_data/test_data.toml 时发生错误：", err)
		return
	}
	// 读取测试用源字符串
	var type1 interface{}
	type1 = datas.Get("source_string")
	if fmt.Sprintf("%T", type1) != "*toml.Tree" {
		t.Error("测试数据中没有找到 source_string 表")
		return
	}
	source := type1.(*toml.Tree)
	// 读取测试方案
	type1 = datas.Get("ReadBetween")
	if fmt.Sprintf("%T", type1) != "[]*toml.Tree" {
		t.Error("测试数据中没有找到 ReadBetween 表数组")
		return
	}
	testSchemes := type1.([]*toml.Tree)
	// 读取测试数据
	var tData tests
	for _, tt := range testSchemes {
		type1 := tt.Get("name")
		if fmt.Sprintf("%T", type1) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 name 字段")
			return
		}
		type2 := tt.Get("args.start")
		if fmt.Sprintf("%T", type2) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 args.start 字段")
			return
		}
		type3 := tt.Get("args.end")
		if fmt.Sprintf("%T", type3) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 args.end 字段")
			return
		}
		s1 := tt.Get("source")
		if fmt.Sprintf("%T", s1) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 source 字段")
			return
		}
		tData = tests{
			name: type1.(string),
			args: args{
				s:     source.Get(s1.(string)).(string),
				start: type2.(string),
				end:   type3.(string),
			},
			want: tt.Get("want").(string),
		}
		t.Run(tData.name, func(t *testing.T) {
			got := ReadBetween(tData.args.s, tData.args.start, tData.args.end)
			if got != tData.want {
				t.Errorf("ReadBetween() = %v, want %v", got, tData.want)
			}
		})
	}
	// 错误测试
	t.Run("错误测试：返回内容不符", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "d/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("错误测试：开头字符找不到", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/1AQQ", "d/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("错误测试：结尾字符找不到", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "d1/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("错误测试：结尾字符在开头字符之前", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "inbox"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
}
