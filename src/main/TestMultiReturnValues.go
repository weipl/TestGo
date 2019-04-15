package main

import (
	"fmt"
)

func SumAndProduct(A,B int) (int,int){
	return A+B, A*B
}

func main(){
	x := 3
	y := 4

	xPLUSy,xTIMESy := SumAndProduct(x,y)
	fmt.Printf("%d + %d = %d\n",x,y,xPLUSy)
	fmt.Printf("%d * %d = %d\n",x,y,xTIMESy)
}

/*
	技术总结,----->main函数只能在main包下运行
 */