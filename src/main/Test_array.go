package main

import "fmt"

func main() {
	//如果数组中长度位置出现"..."省略号,则表示数组的长度是根据初始化的个数计算
	//数组的长度必须是常量表达式,因为数组的长度需要在编译时期确定
	number1 := [3]int{1,2,2:3}
	number2 := [3][4]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
	}
	number3 := [2][2][3]int{
		{{1,2,3},{4,5,6}},
		{{7,8,9},{10,11,12}},
	}
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)

	number4 := [...]int{3,1,9,6,8}

	for i:=0;i< len(number4) ;i++  {
		for j := i+1;j < len(number4) ;j++  {
			if number4[i] < number4[j] {
				number4[i],number4[j] = number4[j], number4[i]
			}
		}
	}
	fmt.Println(number4)
}
