package variable

import (
	"fmt"
)

// const 定义常量，只读
const c1 = 100

// const 定义枚举类型，
const (
	// iota关键字(只能配合const()使用)，第一行iota的默认值为0，每行的iota会累加1
	AA = 10 * iota //iota = 0
	BB             //iota = 1
	CC             //iota = 2
)

func ConstAndIota() {
	fmt.Println("global variable c1 =", c1)

	fmt.Println("AA =", AA, "BB =", BB, "CC =", CC)
}
