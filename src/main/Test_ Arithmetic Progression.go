package main

import "fmt"

func main() {
	arr01 := [...][3]int{
		{1, 15 ,11},
		{23 ,17, 49},
		{6, 4 ,33},
		{29, 1, 35},
		{12, 12, 68},
		{23 ,9, 77},
		{22 ,3 ,16},
		{18 ,0 ,74},
		{15 ,17 ,49},
		{16 ,2, 45}}

	arr02 := [len(arr01)]int{}

	for index :=0;index < len(arr01) ;index++  {
		arr02[index] = (arr01[index][0] + ( ( arr01[index][2] -1) * arr01[index][1] + arr01[index][0])) * arr01[index][2] / 2.0
		fmt.Printf("%d ",arr02[index])
	}
}
