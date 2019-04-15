package main

import "fmt"

func main(){
	switch number1 := 3; number1{
	case 1: fmt.Println(1)
	case 3: fmt.Println(3)
	case 5: fmt.Println(5)
	}

	number2 := 6
	switch {
	case number2 == 2: fmt.Println(2)
	case number2 == 4: fmt.Println(4)
	case number2 == 6: fmt.Println(6)
		fallthrough	//	fallthrough 怎么没起作用,这里应该继续执行打印"default"
					// 分析:fallthrough要和case相连,不然不会执行default
	default:
		fmt.Println("default")
	}
}
