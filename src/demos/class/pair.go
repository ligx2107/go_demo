package cla

import "fmt"

/*
	Pair特性：
		golang中，每个变量都具有Pair<type, value>特性，标识了变量的具体类型和值，在变量传递过程中，Pair特性保持不变
*/

// 定义Reader Writer接口
type Reader interface {
	read()
}

type Writer interface {
	write()
}

// 定义接口实现类，同时实现Reader和Writer接口
type Book struct {
}

func (this *Book) read() {
	fmt.Println("read is called...")
}

func (this *Book) write() {
	fmt.Println("write is called...")
}

func TestPair() {
	// b: pair<type:Book, value: Book对象地址>
	b := &Book{}

	var r Reader
	// r: pair<type:Book, value: Book对象地址>
	r = b
	r.read()

	var w Writer
	// w: pair<type:Book, value: Book对象地址>
	w, _ = r.(Writer) //此处断言成功，是基于多态特性，w和r的具体type是相同的，都为Book
	w.write()
}
