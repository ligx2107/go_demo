package array

import "fmt"

// 声明函数，参数为静态数组类型
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

// 声明函数，参数为动态数组类型
func dynamicArray(arr []int) {
	// 修改数组元素
	for index, value := range arr {
		arr[index] = value * 2
	}

}

func TestStaticArray() {
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

}
