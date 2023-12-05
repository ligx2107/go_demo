package array

import "fmt"

// 声明函数，参数为数组类型
func staticArray(arr1 [3]int) {
	// 遍历数组，输出数组内数据
	// 1. 传统方式
	//for i := 0;i<len(arr1);i++ {
	//	fmt.Println(i, "--->", arr1[i])
	//}

	// 2. range方式遍历
	for index, value := range arr1 {
		fmt.Println(index, "--->", value)
	}
}

// 声明函数，参数为切片类型
func dynamicArray(arr []int) {
	// 修改数组元素
	for index, value := range arr {
		arr[index] = value * 2
	}

}

// 通过make分配空间并初始化
func makeSlice() {
	var slice1 []int //此时，slice1没有被分配地址空间
	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}
	slice1 = make([]int, 3)
	slice1[0] = 1
	fmt.Printf("slice1 length, %d, slice value: %v\n", len(slice1), slice1)

	// 3长度，5容量，容量>=长度
	slice2 := make([]int, 3, 4)
	fmt.Printf("slice2 length, %d, slice2 capacity: %d, slice2 values: %v\n", len(slice2), cap(slice2), slice2)

	// append方法，向切片尾部添加元素，当capacity足够时，直接添加；capacity不够时，创建新的切片，capacity是原来的2倍
	slice2 = append(slice2, 1)
	fmt.Printf("slice2 length, %d, slice2 capacity: %d, slice2 values: %v\n", len(slice2), cap(slice2), slice2)
	slice2 = append(slice2, 2)
	fmt.Printf("slice2 length, %d, slice2 capacity: %d, slice2 values: %v\n", len(slice2), cap(slice2), slice2)

	/**
	  切片截取
		1. [:]，通过下标区间方式进行截取，左闭右开，下标值在源切片capacity范围内(超出范围则出现异常)，新切片和源切片地址指向同一内存空间
		2. 使用copy函数进行深拷贝，新切片与源切片地址指向不同内存空间
	*/
	slice3 := make([]int, 3, 4)
	fmt.Printf("slice3 values: %v\n", slice3)
	slice4 := slice3[0:4]
	fmt.Printf("slice4 values: %v\n", slice4)
	//fmt.Printf("slice3 address: %p, slice4 address: %p\n", &slice3, &slice4)
	slice4[1] = 2
	fmt.Printf("after change, slice3 values: %v\n", slice3)

	slice5 := make([]int, 3)
	copy(slice3, slice5)
	fmt.Printf("slice5 values: %v\n", slice5)
	slice5[1] = 2
	fmt.Printf("slice3 values: %v\n", slice3)
	fmt.Printf("slice5 values: %v\n", slice5)
}

func TestArray() {
	// 定义静态数组
	myArray1 := [3]int{1, 2, 3}
	myArray2 := [4]int{1, 2, 3}

	// [3]int和[4]int是不同的两种类型
	fmt.Printf("myArray1 type: %T\n", myArray1) // [3]int
	fmt.Printf("myArray2 type: %T\n", myArray2) // [4]int
	staticArray(myArray1)                       //此处，只能传递[3]int类型的静态数组，且为值传递

	// 定义动态数组
	myArray3 := []int{1, 2, 3}
	fmt.Println("before......")
	for index, value := range myArray3 {
		fmt.Println(index, "--->", value)
	}
	dynamicArray(myArray3) // 动态数组为引用传递
	fmt.Println("after......")
	for index, value := range myArray3 {
		fmt.Println(index, "--->", value)
	}

	makeSlice()
}
