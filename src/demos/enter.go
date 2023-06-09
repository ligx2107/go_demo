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

	// reflect
	fmt.Println("-------------reflect--------------")
	ref.TestReflect()

	// goroutine
	fmt.Println("-------------goroutine--------------")
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
