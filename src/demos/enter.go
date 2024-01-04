package main

import (
	"context"
	ar "demos/array"
	atomicdemo "demos/atomic"
	"demos/class"
	"demos/closure"
	contextdemo "demos/context"
	dd "demos/defer"
	err "demos/errors"
	fun "demos/function"
	"demos/generic"
	gr "demos/goroutine"
	_ "demos/init"
	"demos/map"
	"demos/pointer"
	ref "demos/reflect"
	"demos/struct"
	syncd "demos/sync"
	vd "demos/variable"
	"fmt"
	"os"
	"os/signal"
)

/**
  import导入说明：使用标准库或第三方包前，须用import导入，参数是工作空间中以src为开始的绝对路径，编译器从标准库开始搜索，然后依次搜索GOPATH列表中的各个工作空间
	1. import ""单个导入，import()多个导入
	2. import的参数为欲导入文件的路径，与导入文件的package无关
	3. 别名导入，在调用程序中为导入包指定别名，并使用别名进行调用
	4. 匿名导入(_)，导入的包不在程序中使用，但包中的init方法会被调用

  package包说明：
	1. package由一个或多个保存在同一个目录下(不含子目录)的源码文件组成，包的作用类似名字空间，是成员作用域和访问权限的边界。
	2. 包名与目录名并无关系，不要求保持一致
	3. 同一目录下，所有源码文件必须使用相同包名
	4. 所有成员在包内均可访问，无论是否在同一个源码文件中(可通过指针转换等方式绕开此限制)
	5. 内部包(internal), 导出路径包含internal关键字的包，只允许internal的父级目录及父级目录的子目录导入，其他包无法导入。
*/

func main() {
	// 调用导入文件内方法，使用"包名.方法名"根式，与文件路径无关
	// variable declaration
	fmt.Println("--------------variable declaration---------------")
	vd.VariableDeclaration()
	vd.ConstAndIota()

	// function
	fmt.Println("-------------function--------------")
	fun.TestFunction()

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
	dd.DeferClosure()

	// array
	fmt.Println("-------------array--------------")
	ar.TestArray()

	// map
	fmt.Println("-------------map--------------")
	md.TestMap()

	// struct
	fmt.Println("-------------struct--------------")
	sd.TestStruct()
	sd.TestStructTag()

	// class
	fmt.Println("-------------class--------------")
	cla.TestClass()
	cla.TestInherit()
	cla.TestPolymorphisms()
	cla.TestAssertion()
	cla.TestPair()
	cla.TestMethodSet()

	// reflect
	fmt.Println("-------------reflect--------------")
	ref.TestReflect()

	// goroutine
	fmt.Println("-------------goroutine--------------")
	gr.TestGosched()
	gr.TestGoexit()
	//gr.TestGoroutine()
	gr.TestChannel()
	gr.TestSelect()

	// closure
	fmt.Println("-------------closure--------------")
	closure.TestClosure()

	// error
	fmt.Println("-------------error--------------")
	err.TestErrorHandle()
	err.TestPanic()

	// sync
	fmt.Println("-------------sync---------------")
	syncd.TestSync()

	// atomic
	fmt.Println("-------------atomic---------------")
	atomicdemo.TestAtomic()

	// generic
	fmt.Println("-------------generic---------------")
	generic.TestGeneric()

	// context
	fmt.Println("------------context----------------")
	contextdemo.TestContext()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()
	<-ctx.Done()

	// others
	fmt.Println("-------------others--------------")
}
