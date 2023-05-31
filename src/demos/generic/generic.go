package generic

import (
	"fmt"
)

// 自定义范型约束
type Num1 interface {
	int | uint | int64
}

type Num2 interface {
	int | float32 | float64
}

// 自定义约束并集
type Num3 interface {
	Num1 | Num2
}

// 自定义约束交集
type Num4 interface {
	Num1
	Num2
}

// 自定义约束(超集)
type Num5 interface {
	~int | ~float32 // 不仅包含int和float32类型，还包含最底层是这两种类型的其他类型
}

// 定义范型切片
type Slice1[T Num1] []T

// 为自定义范型切片绑定方法
func (t Slice1[T]) sum() T {
	var res T
	for _, v := range t {
		res += v
	}
	return res
}
func (t Slice1[T]) exist(arg T) bool {
	for _, v := range t {
		if v == arg {
			return true
		}
	}

	return false
}

func testGenericSlice() {
	// 声明范型切片
	type Slice1[T int | float32] []T

	// 实例化范型切片
	s1 := make(Slice1[int], 2)
	s1[0] = 1
	s1[1] = 2
	fmt.Println("s1 values: ", s1)

	s2 := Slice1[float32]{1.2, 2.3}
	fmt.Println("s2 values: ", s2)
}

func testGenericMap() {
	// 声明范型map
	type Map1[K string, V int | float32 | string] map[K]V

	// 实例化范型map
	m1 := make(Map1[string, int], 2)
	m1["k1"] = 1
	m1["k2"] = 2

	m2 := Map1[string, string]{
		"k1": "v1",
		"k2": "v2",
	}

	fmt.Println("m1 ---> ", m1, ", m2 --> ", m2)
}

func testGenericStruct() {
	// 声明范型结构体
	type Struct1[T int | string] struct {
		title string
		value T
	}

	// 实例化范型结构体
	s1 := Struct1[int]{"t1", 1}
	s2 := Struct1[string]{"t1", "v1"}
	fmt.Println("s1 =", s1, ", s2 =", s2)

}

func testNestedGeneric() {
	// 声明范型嵌套的map
	type Map1[K string, S int | string, V []S] map[K]V

	// 实例化
	m1 := make(Map1[string, int, []int], 2)
	m1["k1"] = []int{1, 2}
	m1["k2"] = []int{2, 3}
	m1["k3"] = []int{4, 5}
	fmt.Println("m1 =", m1)
}

func testGenericFunc[T int | float32 | float64](list ...T) {
	for index, value := range list {
		fmt.Println("index ", index, ", value ", value)
	}
}

func testGenericMethod() {
	s := make(Slice1[uint], 2)
	s[0] = 1
	s[1] = 4
	fmt.Println("sum: ", s.sum())
	fmt.Println("exist: ", s.exist(4))
}

func TestGeneric() {
	testGenericSlice()
	testGenericMap()
	testGenericStruct()
	testNestedGeneric()
	testGenericFunc(1.1, 2.2, 3.3)
	testGenericMethod()
}
