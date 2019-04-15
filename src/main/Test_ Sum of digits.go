package main

import (
	"fmt"
	"math"
)

func main() {
	arr01 := [...][3]int{
		{187, 64 ,122},
		{283, 205, 125},
		{295, 32 ,72},
		{356, 258, 32},
		{14 ,96 ,122},
		{395, 95, 179},
		{173 ,52, 129},
		{11 ,286, 116},
		{200 ,148, 11},
		{26 ,297, 148}}
	sum := [len(arr01)]int{}
	digits := [len(arr01)]int{}
	for index := 0;index < len(arr01) ;index++  {

		sum[index] = arr01[index][0] * arr01[index][1] + arr01[index][2]
		//每一个数字有多少位
		bits := int((math.Floor(math.Log10(  float64(sum[index]) ))))

		//先算出来个位数字
		digits[index] = sum[index] % 10
		//然后再算其他位
		for i := 1;i <= bits  ;i++ {

			sum[index] = sum[index] / 10
			digits[index] += sum[index] % 10

		}

		fmt.Printf("%d ", digits[index])
	}
}
