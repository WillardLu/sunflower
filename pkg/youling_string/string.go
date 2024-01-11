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
	"strings"
)

// @brief 读取两个特定字符串之间的字符串
//  @param str    源字符串
//  @param start  开头字符串
//  @param end    结尾字符串
//  @return 两个特定字符串之间的字符串
//  @remark 开头字符串是从左边开始搜索，结尾字符串是从右边开始搜索。
func ReadBetween(str string, start string, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.LastIndex(str[s:], end)
	if e == -1 {
		return ""
	}
	return str[s:][:e]
}
