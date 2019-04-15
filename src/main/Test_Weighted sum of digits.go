package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr01 := [...]string{
		"277136", "188", "9479", "370858489", "542863884", "201064295", "163131", "1525679", "4441", "967", "21", "13112", "38361", "754", "102",
		"238545", "113162", "16", "524", "910161", "86688", "318447", "649791", "0", "23252667", "2172", "2059765", "6805398", "76", "133134", "18298", "1323"}

	arr02 := [len(arr01)]int{}
	for index := 0;index < len(arr01) ;index++  {

		l1 :=len([] rune(arr01[index]))
		arr02[index],_ = strconv.Atoi(arr01[index])
		yushu ,x := 0,0
		for i := 0;i < l1 ;i++  {

			yushu = arr02[index] % 10
			arr02[index] = arr02[index] /10
			x +=  yushu * (len(arr01[index]) -i)
			if i+1 == l1 {
				arr02[index] = x
				fmt.Printf("%d ",arr02[index])
			}

		}


	}
}
