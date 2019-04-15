package main

import "fmt"

func main() {
	arr01 :=[...][3]int{
		{388, 431 ,194},
		{2469 ,574, 1147},
		{386 ,242, 416},
		{668 ,1290 ,1930},
		{699 ,527, 1137},
		{2514, 744, 1252},
		{926 ,1440, 1994},
		{1638 ,790 ,520},
		{366, 662 ,263},
		{624 ,1165, 1432},
		{972 ,1193, 700},
		{1101 ,2691, 611},
		{771 ,301 ,349},
		{1114, 2238 ,572},
		{1736 ,931, 655},
		{820 ,1426 ,2541},
		{412, 707 ,667},
		{414 ,728, 388},
		{781, 2001, 1177},
		{713 ,1713, 656},
		{393, 206, 835},
		{695, 465, 1532}}
	//arr02 :=[len(arr01)]int{}
	//这是判断三个数能否组成三角形的算法 ,最直白,但是没有意思的算法
/*	for index:=0;index < len(arr01) ;index++  {
		if((arr01[index][0] > arr01[index][1]) && (arr01[index][0] > arr01[index][2])){
			if arr01[index][0] < (arr01[index][1] + arr01[index][2]) {
				arr02[index] = 1
			}else {
				arr02[index] = 0
			}
		}else if((arr01[index][1] > arr01[index][2]) && (arr01[index][1] > arr01[index][0])){
			if arr01[index][1] < (arr01[index][2] + arr01[index][0]) {
				arr02[index] = 1
			}else {
				arr02[index] = 0
			}
		}else if ((arr01[index][2] > arr01[index][0]) && (arr01[index][2] > arr01[index][1])){
			if arr01[index][2] < (arr01[index][0] + arr01[index][1]) {
				arr02[index] = 1
			}else {
				arr02[index] = 0
			}
		}
		fmt.Printf("%d ",arr02[index])
	}*/

	for index:= 0;index < len(arr01) ;index++  {
		A,B,C := arr01[index][0],arr01[index][1],arr01[index][2]
		if A+B >C && A+C > B && C+B >A{
			fmt.Print("1 ")
		}else {
			fmt.Print("0 ")
		}
	}
}
