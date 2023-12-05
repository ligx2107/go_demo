package deferDemo

import "fmt"

/**
  defer: Golang提供的一种用于注册延迟调用的机制，让函数或语句可以在当前函数执行完毕后(包括通过return正常结束)执行
	1. defer后面的语句在执行的时候，函数调用的参数会被保存起来(参数预计算)，但是不执行
	2. 多个defer执行顺序: defer语句采用压栈方式，LIFO，后声明的先被调用
	3. defer与return的执行顺序：return先于defer执行
	4. 方法正常结束和异常结束都会调用defer声明的延迟函数，可以有效避免因异常导致的资源无法释放的问题ß
*/

func test1() {
	fmt.Println("test1 called...")
}

func test2() {
	fmt.Println("test2 called...")
}

func deferFun() {
	fmt.Println("deferFun called...")
}

func returnFun() int {
	fmt.Println("returnFun called...")
	return 1
}

func DeferFunc() {
	// 声明两个defer，采用压栈方式，后入先出，test2先于test1被执行
	defer test1()
	defer test2()
}

func DeferReturnFunc() int {
	// return先于defer被执行
	defer deferFun()
	return returnFun()
}

type Person struct {
	name string
}

func (p *Person) printName() {
	fmt.Println("person name: ", p.name)
}

func DeferClosure() {
	persons := []Person{Person{"zhangsan"}, Person{"lisi"}, Person{"wangwu"}}
	for _, p := range persons {
		defer p.printName()
	}

	for i := 0; i < len(persons); i++ {
		defer persons[i].printName()
	}
}
