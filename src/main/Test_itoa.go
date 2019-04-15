package main

import "fmt"

//itoa是常量计数器,从0开始,组中定义一个人常量自动+1
//没遇到一个const,itoa会重置为0

const(
	GB int = 1 << (iota * 10)
	MB int = 1 << (iota * 10)
	KB int = 1 << (iota * 10)
)

const(
	a int = iota
	b
)
func main(){

	fmt.Println(GB,MB,KB)
	fmt.Println(a,b)
}