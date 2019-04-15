/*
	1)、内建函数make分配和初始化一个slice 或 map 或 chan 对象 ,并且只能是这三个对象
	2)、和new类似，第一个参数也是一个类型而不是一个值，不同的是make返回类型的引用而不是  指针，而返回值也以来于具体传入的类型
	3)、slice：第二个参数只当它的长度，此时他的容量和长度相同。可以用第三个参数来指定不同容量大小，但是不能小于它的长度(第二个参数)
	4)、map:根据size大小初始化分配内存，不过分配的map长度为0. 如果size 被忽略了，那么会在初始化分配内存的时候，分配一个小尺寸的内存
	5)、channel: 管道缓冲区容量被初始化。如果容量为0或者被忽略， 管道是没有缓冲区的
 */

// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.

// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.
package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}
func main() {
	number1 := [5] int{}
	number2 := new([5]int)		//and the value returned is a pointer to a newly

	fmt.Println(number1)
	fmt.Println(number2)

	fmt.Println("========================分割线============================")

	p1 := person{}			//	返回类型
	p2 := &person{}			//	返回指针
	p3 := new(person)		//	和p2一样

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println("========================分割线============================")

	// nil是一个预定义标识符，其代表（用作）一些类型的零值 ; 这些类型包含 : pointer ,channel , func , interface , map ,slice

	var  n1 []int
	var  n2 map[int] string
	var  n3 chan  int

	if n1 == nil {
		fmt.Println(n1)
	}
	if n2 == nil{
		fmt.Println(n2)
	}
	if n3 == nil{
		fmt.Println(n3)
	}

}
