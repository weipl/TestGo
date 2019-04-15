package main

import (
	"fmt"
	"strings"
)

//string类型的零值为长度为零的字符串  ,即空字符串""
//字符串的内容可以通过标准索引来获取,
	//字符串啊str的第一个字节: str[0]
	//第i个字节: str[i-1]
	//最后一个字节:str[len(str)-1]

func main() {
	s := "hel" + "lo,"
	s += "world! world"

	old := "world"
	new := "weipl"
	fmt.Println(s)
	// n < 0 ,用 new 替换所有匹配上的 old；n=-1
	s2 := strings.Replace(s,old,new,-1)
	fmt.Println(s2)
	// 不替换任何匹配的 old；n=0
	s2 = strings.Replace(s,old,new,0)
	fmt.Println(s2)
	//用 new 替换第一个匹配的 old；n=1
	s2 = strings.Replace(s,old,new,1)
	fmt.Println(s2)
	//用 new 替换第一个匹配的 old；n=2 .类似如果s中有多个old的,n可以多个
	s2 = strings.Replace(s,old,new,2)
	fmt.Println(s2)

	fmt.Println("=========================统计字符串出现字数===========================")
	var Testcount = "wwwwwwwwwwwwww"

	fmt.Printf("Number of H's %s is:",s)
	fmt.Printf("%d\n",strings.Count(s,"h"))

	fmt.Printf("Number of double w's %s is:",Testcount)
	fmt.Printf("%d\n",strings.Count(Testcount,"ww"))
}
