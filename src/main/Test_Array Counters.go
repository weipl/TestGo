package main

import (
"fmt"
"strconv"
"strings"
)

//data input:
//10 3
//3 2 1 2 3 1 1 1 1 3
//
//answer:
//5 2 3
//输入3，要把3前面遇到1，2遇到的个数也数出来

type intArr1 [7]int
//type  intArr2 []intArr1
//1、str中添加一个字符串
//2、用strings.split把字符串按" "分割出来的子字符串数组
//3、把得到字符串数组转成数字
func main() {
	str :="5 4 6 5 4 1 5 6 1 2 6 1 5 2 7 3 5 5 2 3 4 7 2 6 6 3 7 2 7 2 3 4 6 2 2 2 2 6 1 3 7 6 4 5 1 3 7 5 1 1 7 4 1 2 3 7 4 3 2 4 4 5 1 3 6 2 4 7 7 4 3 7 2 6 4 2 2 4 7 2 5 6 6 6 1 2 5 5 4 6 1 1 4 1 3 2 3 7 2 2 3 4 2 5 3 6 7 4 2 6 6 7 5 4 5 5 5 3 3 2 1 3 3 4 4 5 6 6 4 7 1 7 4 3 4 6 1 4 3 3 2 1 2 7 5 6 5 3 2 7 4 2 3 6 6 6 4 5 5 1 4 6 7 1 1 4 7 2 7 3 4 2 3 6 1 1 5 6 3 6 5 7 1 7 6 7 6 2 4 4 2 1 2 2 1 3 5 1 4 5 3 1 7 6 7 7 6 4 6 1 3 3 1 3 3 6 2 2 1 6 5 3 6 7 4 7 2 2 7 6 7 2 6 6 1 5 6 6 2 4 7 4 7 7 6 3 5 1 5 6 6 2 1 5 2 5 4 3 7 4 1 6 6 7 5 6 4 3 4 6 7 3 2 7 3 1 2 7 1 6 6 7 1 6 5 2 4 2 5 3 5 6 2 4 5 6"
	outArr := intArr1{}
	strArr := strings.Split(str," ")

	for _,s:= range strArr{
		i,_ :=  strconv.Atoi(s)
		outArr[i-1] ++
	}
	fmt.Printf("%d",outArr)

}
