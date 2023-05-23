package cla

import "fmt"

// 定义父类
type Human struct {
	name string
	age  int
}

// 定义父类方法
func (this *Human) say() {
	fmt.Println("Human say....")
}

func (this *Human) walk() {
	fmt.Println("Human walk....")
}

// 定义子类继承父类
type SuperHuman struct {
	Human         // 标识SuperHuman继承自Human父类
	gender string // 定义子类属性
}

// 重写父类方法
func (this *SuperHuman) say() {
	fmt.Println("SuperHuman say....")
}

// 定义子类方法
func (this *SuperHuman) fly() {
	fmt.Println("SuperHuman fly....")
}

func TestInherit() {
	// 创建子类对象
	sm := SuperHuman{Human{"zhangsan", 10}, "male"}
	fmt.Println(sm)

	sm.say()  //调用SuperHuman重写后的say方法
	sm.walk() //调用父类的walk方法
	sm.fly()  //调用新增加的子类方法fly

	var sm1 SuperHuman
	sm1.name = "lisi"
	sm1.age = 15
	sm1.gender = "female"
	fmt.Println(sm1)
}
