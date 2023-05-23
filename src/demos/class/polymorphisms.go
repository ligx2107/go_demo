package cla

import "fmt"

/*
  多态：
	1. 定义接口interface
		1.1 接口只有方法声明，没有方法实现，没有数据字段
		1.2 接口可以匿名嵌入其他接口或嵌入到结构体中
		1.3 一个类型可以实现多个接口
	2. 定义类实现接口，即：实现接口的所有方法
		2.1 类实现接口分为两种模式，值接收者实现接口和指针接收者实现接口
		2.2 当以值接收者方式实现接口时，赋值给接口类型变量的值，可以是结构体类型，也可以是结构体的指针类型
		2.3 当以指针接收者方式实现接口时，赋值给接口类型变量的值，只能是结构体的指针类型
	3. 父类类型的变量(指针)指向子类的具体数据变量
*/

// 定义接口
type Animal interface {
	Sleep()
	Run()
}

// 定义Cat子类实现接口
type Cat struct {
	color string
}

// 指针接收者方式实现接口
func (this *Cat) Sleep() {
	fmt.Println("cat sleep...")
}

func (this *Cat) Run() {
	fmt.Println("cat run...")
}

// 定义Dog子类实现接口
type Dog struct {
	color string
}

// 值接收者方式实现接口
func (this Dog) Sleep() {
	fmt.Println("dog sleep...")
}
func (this Dog) Run() {
	fmt.Println("dog run...")
}

// 多态 接收指针类型参数
func test(animal Animal) {
	animal.Sleep()
	animal.Run()
}

func TestPolymorphisms() {
	cat := Cat{"Yellow"}
	test(&cat) // Cat以指针接收者方式实现接口，需使用结构体的指针对接口类型变量赋值

	dog := Dog{"White"}
	test(dog)
}
