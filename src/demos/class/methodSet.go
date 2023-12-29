package cla

import (
	"fmt"
	"reflect"
)

/**
*	方法集：类型有一个与之关联的方法集合，这决定了它是否实现了某个接口
*		1. 类型S方法集包含所有receiver S方法
*		2. 类型*S方法集包含所有receiver S + *S方法
*		3. 匿名嵌入S，T方法集包含所有receiver T + S方法
*		4. 匿名嵌入*S，T方法集包含所有receiver T + S + *S方法
*		5. 匿名嵌入S或*S，*T方法集包含所有receiver S + *S + T + *T方法
 */

type S struct {
}

type T struct {
	*S // 匿名嵌入S
}

func (s S) SVal() {}

func (s *S) SPtr() {}

func (t T) TVal() {}

func (t *T) TPtr() {}

func methodSet(p interface{}) {
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, ":", m.Type)
	}
}

func TestMethodSet() {
	var s S
	var t T

	fmt.Println("S method set: ")
	methodSet(s)
	fmt.Println("*S method set: ")
	methodSet(&s)
	fmt.Println("T method set: ")
	methodSet(t)
	fmt.Println("*T method set: ")
	methodSet(&t)
}
