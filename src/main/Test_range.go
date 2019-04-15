package main

import "fmt"

func main() {

	//range用于for 循环中迭代数组(array)、切片(slice)、通道(channel)或者集合(map)的元素
	number := [5]string{"a","b","c","d","e"}
	for k,v:= range number{
		fmt.Println(k,v)
	}
	//省略key的的写法
	for _,v := range number {
		fmt.Println(v)
	}
}
