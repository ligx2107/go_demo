package md

import "fmt"

// 声明map类型变量
func declareMap() {
	// 声明并赋值
	myMap1 := map[string]string{
		"one": "java",
		"two": "kotlin",
	}
	fmt.Println("myMap1 -->", myMap1)

	// make方法分配空间
	myMap2 := make(map[string]string, 2)
	myMap2["a"] = "java"
	myMap2["b"] = "kotlin"
	myMap2["c"] = "golang"
	fmt.Println("myMap2 -->", myMap2)
}

// map基本操作
func operateMap() {
	map1 := map[string]int{
		"java":   1000,
		"golang": 300,
	}

	fmt.Println("original map -->", map1)

	// 遍历map
	for key, value := range map1 {
		fmt.Println("key -->", key, "value -->", value)
	}

	// 修改map
	map1["java"] = 2000
	fmt.Println("after change -->", map1)

	// 删除map元素
	delete(map1, "java")
	fmt.Println("after delete -->", map1)
}

func TestMap() {
	declareMap()
	operateMap()
}
