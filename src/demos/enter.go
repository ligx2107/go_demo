package main

import (
	ar "demos/array"
	dd "demos/defer"
	fu "demos/function" // 别名导入
	"demos/pointer"
	vd "demos/variable"
	"fmt"
)

/**
  import导入说明：
	1. import ""单个导入，import()多个导入
	2. import的参数为欲导入文件的路径，与导入文件的package无关
	3. 别名导入，在调用程序中为导入包指定别名，并使用别名进行调用
	4. 匿名导入(_)，导入的包不在程序中使用，但包中的init方法会被调用
*/

func main() {
	// 调用导入文件内方法，使用"包名.方法名"根式，与文件路径无关
	// variable declaration
	fmt.Println("--------------variable declaration---------------")
	vd.VariableDeclaration()
	vd.ConstAndIota()

	// function
	fmt.Println("-------------function--------------")
	r1, r2 := fu.MultipleAnonymousReturn("world", 100)
	fmt.Println("r1 =", r1, "r2 =", r2)
	r3, r4 := fu.MultipleNamedReturn("Golang", 200)
	fmt.Println("r3 =", r3, "r4 =", r4)

	// pointer
	fmt.Println("-------------pointer--------------")
	ia, ib := 10, 20
	fmt.Println("before swap: ia =", ia, "ib =", ib)
	pointer.Swap(&ia, &ib)
	fmt.Println("after swap: ia =", ia, "ib =", ib)

	// defer
	fmt.Println("-------------defer--------------")
	dd.DeferFunc()
	dd.DeferReturnFunc()

	// array
	fmt.Println("-------------array--------------")
	ar.TestArray()
}
