package main

import "fmt"

func main() {
	arr01 := [...]int{
		978, 94, 77623, 36 ,7532248, 2741893 ,858940247, 1902501 ,82760568 ,24614678 ,
		9 ,821827 ,9 ,5999, 3664945, 44 ,4696 ,788438897, 252 ,1799230, 5, 975, 575}
	result,v := 0,0
	for i:=0;i<len(arr01) ;i++  {
		v = arr01[i]
		result = (result+v)*113%10000007
	}

	fmt.Printf("%d",result)

}
