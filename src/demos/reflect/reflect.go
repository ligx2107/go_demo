package ref

import (
	"fmt"
	"reflect"
)

// 定义类
type Student struct {
	name        string
	age         int
	gender      string
	Description string
}

func (this Student) Study() {
	fmt.Println(this.name, "is studying...")
}

func (this Student) sleep() {
	fmt.Println(this.name, "is sleeping...")
}

func (this Student) Add(v1, v2 int) int {
	return v1 + v2
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

	// 获取所有method
	// 非接口类型，返回公有方法的数量
	// 接口类型，返回公有和私有方法的数量总和
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Println("method name", method.Name, ", is exported", method.IsExported())
	}
}

func TestReflect() {
	s := Student{"zhangsan", 18, "male", "A handsome boy"}
	s.Study()
	getFieldAndMethod(s)

	// change variable's value by reflection
	num := 100
	v := reflect.ValueOf(num)
	fmt.Println("value can be changed: ", v.CanSet()) // CanSet()判断当前Value是否可以修改

	v1 := reflect.ValueOf(&num) // 传递变量指针
	fmt.Println("value can be changed: ", v1.CanSet())

	v1 = v1.Elem() // Elem()获取指针指向的实际变量
	fmt.Println("value can be changed: ", v1.CanSet())

	v1.SetInt(200) // SetInt修改整数类型的值
	fmt.Println("the value after change: ", v1)

	// change struct field's value by reflection, just only public variables can be changed
	elem := reflect.ValueOf(&s).Elem()
	fmt.Println("the original value: ", elem)
	elem.FieldByName("Description").SetString("A ugly boy")
	fmt.Println("the new value: ", s)

	// operate map object by reflection
	data := make(map[string]int)
	data["Tom"] = 18
	data["Jerry"] = 20

	vm := reflect.ValueOf(data)
	fmt.Println("the original value of map: ", vm)

	key := reflect.ValueOf("Tom") // get specific key
	value := vm.MapIndex(key)
	fmt.Println("key Tom's value: ", value)

	vm.SetMapIndex(reflect.ValueOf("Lucy"), reflect.ValueOf(21)) // add new element
	fmt.Println("the new value of map: ", vm)

	vm.SetMapIndex(reflect.ValueOf("Lucy"), reflect.Value{})
	fmt.Println("the new value of map: ", vm) // delete the specific element

	// call method of struct
	addMethod := elem.MethodByName("Add")
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	result := addMethod.Call(args)
	fmt.Println(result[0].Int())
}
