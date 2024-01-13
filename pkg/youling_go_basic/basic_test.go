// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
// @author 陆巍
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_go_basic

/*
// @brief 测试CheckingDataType函数
func TestCheckDataType(t *testing.T) {
	// 基础数据类型
	var data0 int = 10
	var data1 int32 = 20
	var data2 int64 = 30
	var data3 float32 = 40.5
	var data4 float64 = 50.5
	var data5 complex64 = 60 + 0i
	var data6 complex128 = 70 + 0i
	var data7 bool = true
	var data8 string = "hello"
	// 复合数据类型
	var data9 []int = []int{1, 2, 3}
	var data10 map[string]int = map[string]int{"a": 1, "b": 2}
	data11 := struct {
		Name string
		Age  int
	}{Name: "John", Age: 30}
	// 引用数据类型
	var data12 func(int, int) int = func(a, b int) int { return a + b }
	var data13 chan int = make(chan int)
	var data14 error = errors.New("this is an error")
	// 把不同类型的数据放入同一个数组中
	type dataType interface{}
	data := [...]dataType{data0, data1, data2, data3, data4, data5, data6, data7,
		data8, data9, data10, data11, data12, data13, data14}
	type1 := [...]string{"int", "int32", "int64", "float32", "float64",
		"complex64", "complex128", "bool", "string", "[]int", "map[string]int",
		"struct { Name string; Age int }", "func(int, int) int", "chan int",
		"*errors.errorString"}
	// 一般测试
	for i := 0; i < len(data); i++ {
		t.Run(fmt.Sprintf("Testing %v\n", type1[i]), func(t *testing.T) {
			err := CheckDataType(data[i], type1[i])
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
	// 错误测试
	t.Run("Testing error data type", func(t *testing.T) {
		if CheckDataType("test", "int") == nil {
			t.Errorf("[ERROR] Expected error, but no error occurred.")
		}
	})
	// 边界测试
	t.Run("Testing boundary data type", func(t *testing.T) {
		if CheckDataType(nil, "int64") != nil {
			t.Errorf("[ERROR] Expected error, but no error occurred.")
		}
	})
}
*/
