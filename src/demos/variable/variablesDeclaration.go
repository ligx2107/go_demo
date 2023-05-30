package variable

import (
	"fmt"
)

/*
	golang变量类型
		1. 值类型：int, float, bool, string, struct, array等，在做等号赋值操作或函数传参时，进行拷贝复制，值传递
		2. 引用类型：slice, map, channel, interface, 等号赋值或函数传参时，传递的是指向值的指针
		3. make和new：make是初始化内置的数据结构(slice, map, channel)，new是根据传入的类型，分配一片内存空间并返回指向这片内存空间的指针
*/

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

	// make
	slice := make([]int, 2, 3)
	fmt.Println("slice value: ", slice)

	// new
	p := new(int)
	*p = 3
	fmt.Println("p value: ", p, ", *p value: ", *p)
}
