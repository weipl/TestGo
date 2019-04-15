package main

import (
	"fmt"
	_ "fmt"
)

func add1(a int) int{
	a = a + 1
	return a
}

func add2(a *int) int{
	*a = *a + 1
	return *a
}

func main(){
	x := 3

	fmt.Println("x = ",x)

	x1 := add1(x)
	fmt.Println("调用值传递")
	fmt.Println("x+1=",x1)
	fmt.Println("x = ",x)

	x2 := add2(&x)
	fmt.Println("调用指针传递")
	fmt.Println("x+1=",x2)
	fmt.Println("x = ",x)
}