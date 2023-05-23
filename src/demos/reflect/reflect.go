package ref

import (
	"fmt"
	"reflect"
)

// 定义类
type Student struct {
	name   string
	age    int
	gender string
}

func (this Student) Study() {
	fmt.Println(this.name, "is studying...")
}

func getFieldAndMethod(arg interface{}) {
	// 获取参数类型
	argType := reflect.TypeOf(arg)
	fmt.Println("arg type is", argType.Name())

	// 获取参数值
	argValue := reflect.ValueOf(arg)
	fmt.Println("arg value is", argValue)

	// 获取所有field
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		fieldValue := argValue.Field(i)
		fmt.Println("field name", field.Name, ", field value", fieldValue)
	}

	// 获取所有method TODO  NumMethod方法的返回值
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Println("method name", method.Name, ", is exported", method.IsExported())
	}
}

func TestReflect() {
	s := Student{"zhangsan", 18, "male"}
	s.Study()
	getFieldAndMethod(s)
}
