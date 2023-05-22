package cla

import "fmt"

/*
  多态：
	1. 定义接口
	2. 定义类实现接口，即：实现接口的所有方法
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

func (this *Dog) Sleep() {
	fmt.Println("dog sleep...")
}
func (this *Dog) Run() {
	fmt.Println("dog sleep...")
}

// 多态 接收指针类型参数
func test(animal Animal) {
	animal.Sleep()
	animal.Run()
}

func TestPolymorphisms() {
	cat := Cat{"Yellow"}
	test(&cat) // 子类行的指针

	dog := Dog{"White"}
	test(&dog)
}
