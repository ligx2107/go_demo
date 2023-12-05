package sd

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// book, 定义结构体，结构体属性添加tag
type Book1 struct {
	Name   string   `json:"name" info:"name" des:"book name"`
	Price  float32  `json:"price" info:"price"`
	Author []string `json:"authors" info:"authors"`
	// 转json时忽略属性的两种处理方式：私有、tag标签的json:"-"
	Weight float32 `json:"-"`
}

func analyzeStructTag(str Book1) {
	bookType := reflect.TypeOf(str)
	for i := 0; i < bookType.NumField(); i++ {
		tagValue := bookType.Field(i).Tag.Get("info")
		fmt.Println("tagValue -->", tagValue)
	}
}

func parseJson(book Book1) {
	// json.Marshal()方法，将结构体转为json，结构体属性必须是公有的
	jsonStr, err := json.Marshal(book)
	if err != nil {
		fmt.Println("parse json error")
		return
	}
	fmt.Println("json -->", string(jsonStr))

	// json转结构体
	myBook := Book1{}
	json.Unmarshal(jsonStr, &myBook)
	fmt.Printf("%v\n", myBook)
}

func TestStructTag() {
	b1 := Book1{"golang", 15.5, []string{"zhangsan", "lisi"}, 0.75}
	analyzeStructTag(b1)
	parseJson(b1)
}
