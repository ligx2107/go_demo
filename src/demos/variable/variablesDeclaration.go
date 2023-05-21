package variable

import (
	"fmt"
)

func VariableDeclaration() {
	// 1. 声明变量，按变量类型取默认值
	var str1 string
	fmt.Println("str1 =", str1)
	fmt.Printf("type of str1 = %T\n", str1)

	// 2. 声明变量，直接赋值，可省略变量类型说明，编译器通过变量值进行自动推导
	var str2 = "string value"
	fmt.Println("str2 =", str2)

	// 3. 省略var关键字，直接自动匹配，此种方式只支持局部变量声明，不支持全局变量
	str3 := "hello"
	fmt.Println("str3 =", str3)

	// 4. 声明多变量
	var a, b = 100, 200
	fmt.Println("a =", a, "b =", b)
	var c, d = 300, "ddd"
	fmt.Println("c =", c, "d =", d)

	// 5. 多行多变量声明
	var (
		int1  = 100
		bool1 = true
	)
	fmt.Println("int1 =", int1, "bool1 =", bool1)
}
