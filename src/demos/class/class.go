package cla

import "fmt"

/*
Student
使用type定义类(结构体+方法)
 1. 类名首字母大写，此类为公共类，可被其他包使用，类名首字母小写，此类为私有类，只能在包内使用
 2. 类中定义的属性，
    2.1 属性名首字母大写则为公有属性，其他包可通过类直接访问或修改属性
    2.2 属性名首字母小写则为私有属性，只能在包内访问
 3. 绑定类方法
*/
type Student struct {
	Name   string // 公有属性
	Age    int
	gender string // 私有属性
}

// 为结构体绑定Get方法
func (this *Student) getGender() string {
	return this.gender
}

// 为结构体绑定Set方法
func (this *Student) setGender(gender string) {
	this.gender = gender
}

func (this *Student) showInfo() {
	fmt.Println("student info -->", this)
}

func TestClass() {
	student1 := Student{Name: "zhangsan", Age: 20}
	fmt.Println(student1)
	student1.setGender("male")
	fmt.Println("after change -->", student1)
	student1.showInfo()
}
