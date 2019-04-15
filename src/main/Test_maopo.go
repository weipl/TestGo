package main

import "fmt"

func main() {
	number := [5]int{8,6,5,7,3}

	for i:=0;i < len(number) ;i++  {
		for j := i+1;j < len(number) ;j++  {
			if number[i] > number[j] {
				number[i],number[j] = number[j],number[i]
			}
		}
	}

	fmt.Println(number)
}
