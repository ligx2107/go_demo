package pointer

// 通过指针，对数据进行交换
func Swap(pa *int, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}
