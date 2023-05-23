package sd

import "fmt"

// 定义Book结构体
type Book struct {
	name   string
	author string
}

// 结构体类型参数，值传递
func changeStruct(book Book) {
	book.author = "lisi"
}

func changeStruct2(book *Book) {
	book.author = "lisi"
}

func TestStruct() {
	book := Book{"golang", "zhangsan"}
	fmt.Printf("%v\n", book)

	changeStruct(book)
	fmt.Printf("after change -->%v\n", book)

	changeStruct2(&book)
	fmt.Printf("after change by passing address -->%v\n", book)
}
