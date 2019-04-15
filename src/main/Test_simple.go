package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//读数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入您的姓名")
	input,err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n",err)
		//异常退出
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("What can i do for u %s\n",name)
	}

	for {
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n",err)
			continue			//这个为啥不退出,这里是再循环
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":continue
		case "nothing","bye":
				fmt.Println("Bye")
				os.Exit(1)
		default:
			fmt.Println("Sorry,I didn't catch u")
		}
	}
}
