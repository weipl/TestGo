package main

import "fmt"

func main() {

	//make([]T,length长度,capacity容量)
	//一般使用 make()创建 :len获取切片长度 ,cap 获取切片容量
	//一个切片在初始化之前默认为nil ,长度为0
	//与数组相比切片长度不固定,可以追加元素,在追加元素时可能使切片的容量增大

	number1 := make([]int, 9 ,10)
	fmt.Println(len(number1),cap(number1))

	number2 := []int{1,2,3,4,5,6,7,8,9}
	number3 := number2[:6]
	number4 := number2[6:]
	number5 := number2[4:5]

	fmt.Println(number3,number4,number5)

	fmt.Println("========================分割线========================")
	// 'copy' 函数'copy'从源 'slice' 的'src' 中复制元素到目标'dst',并返回复制的元素的个数

	number01 := []int{1,2,3,4,5}
	number02 := []int{6,7,8}
	number03 := copy(number01,number02)			// Copy returns the number of elements copied
	fmt.Println(number01,number02,number03)

	number04 := []int{1,2,3,4,5}
	number05 := []int{6,7,8}
	number06 := copy(number05,number04)
	fmt.Println(number04,number05,number06)

	fmt.Println("========================分割线========================")
	// 'append' 向'slice'里面追加一个或者多个元素,然后返回一个和'slice'一样的类型的'slice
	number001 := []int{1,2,3}
	fmt.Println(number001,len(number001),cap(number001))

	number001 = append(number001,4,5)
	fmt.Println(number001,len(number001),cap(number001))

	number003 := append(number001,number001...) //用省略号自动展开切片,以使用每个元素
	fmt.Println(number003)
}
