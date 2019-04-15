package main

import "fmt"

func main() {
	//支持函数内部 goto跳转 ,continue 跳出当前循环进入下一次循环,break终止循环体
	//break 和 continue 都可以配合标签,在多级嵌套循环建跳出,这和goto调整执行位置完全不同
	//通常建议往后 goto ,避免死循环
	for number:= 1;number < 5 ;number++  {
		if number == 3 {
			break
		}
		fmt.Println("break",number)
	}

	for number:=1;number < 5 ;number ++  {
		if number == 3 {
			continue
		}
		fmt.Println("continue",number)
	}

	laber1:
		for {
			for{
				//配合标签跳出最外层的循环
				//标签名是大小写敏感的
				break laber1
			}
		}
	fmt.Println("Hello World1")
	goto Lable2

	fmt.Println("Hello world2")
	Lable2:
		fmt.Println("Hello World3")
}
