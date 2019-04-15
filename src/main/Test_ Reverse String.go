package main

import (
	"fmt"
)
//byte和rune实质上就是uint8和int32类型
//rune        alias for int32
func main() {
	str := "and till pick turn yield supper bulb about fare moon"
	bytes := []rune(str)

	for from , to := 0 , len(bytes) -1 ; from < to ; from , to = from + 1, to -1{
		bytes[from] , bytes[to] = bytes[to] , bytes[from]
	}
	str = string(bytes)
	fmt.Printf("%s",str)

}
