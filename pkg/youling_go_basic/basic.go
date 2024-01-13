// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
// @author 陆巍
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_go_basic

// @brief 检查类型是否为nil或者是否与预期的数据类型相符
//  @param data 要检查类型的数据
//  @param expectedDataType 预期的数据类型
//  @return 成功：nil，失败：错误信息
/*
func CheckDataType(data interface{}, expectedDataType string) error {
	dataType := fmt.Sprintf("%T", data)
	if dataType == "<nil>" {
		return errors.New("The data type is not " + expectedDataType +
			", it is actually nil.")
	}
	if dataType != expectedDataType {
		return errors.New("The data type is not " + expectedDataType +
			", it is actually " + dataType)
	}
	return nil
}
*/
