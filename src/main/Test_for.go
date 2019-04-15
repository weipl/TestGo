//for支持三种形式
	//for init; condition; post{}
	//for condition{}
	//for{}

package main

import "fmt"

func main(){
	//形式1:
	for number1 :=1;number1 < 3 ;number1++  {
		fmt.Printf("number1:%d\n",number1)
	}

	//形式2:
	number2 :=1
	for number2 < 3{
		fmt.Printf("number2:%d\n",number2)
		number2++		//不能写成 number2 = number2++ (++ 与 -- 是作为语句而不是作为表达式)
	}

	//形式3: for没有条件将无限循环;用break退出循环
	for  {
		fmt.Println("loop")
		break
	}
}
