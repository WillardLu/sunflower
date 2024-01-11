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
	"sunflower/pkg/youling_go_basic"
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
		fmt.Println("\n载入测试数据文件 test_data/test_data.toml 时发生错误：", err)
		return
	}
	// 读取测试用源字符串
	dataType := youling_go_basic.CheckDataType(datas.Get("source_string"),
		"*toml.Tree")
	if dataType != "" {
		fmt.Println(dataType)
		return
	}
	source := datas.Get("source_string").(*toml.Tree)
	// 读取测试方案
	dataType = youling_go_basic.CheckDataType(datas.Get("ReadBetween"),
		"[]*toml.Tree")
	if dataType != "" {
		fmt.Println(dataType)
		return
	}
	testSchemes := datas.Get("ReadBetween").([]*toml.Tree)
	var testData tests
	for _, tt := range testSchemes {
		dataType1 := youling_go_basic.CheckDataType(tt.Get("name"), "string")
		dataType2 := youling_go_basic.CheckDataType(tt.Get("args.start"), "string")
		dataType3 := youling_go_basic.CheckDataType(tt.Get("args.end"), "string")
		if dataType1 != "" || dataType2 != "" || dataType3 != "" {
			fmt.Println(dataType1)
			fmt.Println(dataType2)
			fmt.Println(dataType3)
			return
		}
		testData = tests{
			name: tt.Get("name").(string),
			args: args{
				s:     source.Get(tt.Get("source").(string)).(string),
				start: tt.Get("args.start").(string),
				end:   tt.Get("args.end").(string),
			},
			want: tt.Get("want").(string),
		}
		t.Run(testData.name, func(t *testing.T) {
			if got := ReadBetween(testData.args.s, testData.args.start,
				testData.args.end); got != testData.want {
				t.Errorf("ReadBetween() = %v, want %v", got, testData.want)
			}
		})
	}
	// 错误测试
	t.Run("错误测试", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "d/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
}
