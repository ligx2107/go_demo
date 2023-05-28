package err

import (
	"errors"
	"fmt"
	"math"
)

/*
	golang的错误处理
		1. golang通过实现内置的error接口，实现错误处理
		2. golang中没有try catch语法来捕获异常
		3. 可以自定义错误格式
		4. 通过标准库中的errors包来创建错误，可自定义错误内容
*/

func sqrt(num float64) (result float64, error error) {
	if num < 0 {
		return 0, errors.New("invalid root number ")
	}

	result = math.Sqrt(num)
	error = nil
	return
}

func TestErrorHandle() {
	res, error := sqrt(3)
	if error != nil {
		fmt.Println("error message: ", error)
	} else {
		fmt.Println("result is ", res)
	}
}
