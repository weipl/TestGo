package main

import "fmt"

func main() {
	arr01 := [...][3]int{
		{4 ,22, 846},
		{776, 63, 62},
		{1008 ,51, 991},
		{9 ,13 ,47},
		{45, 44, 42},
		{463 ,465, 31},
		{25 ,219, 125},
		{544 ,648, 552},
		{294, 77 ,84},
		{567 ,40, 585},
		{40, 26, 35},
		{7 ,649, 608},
		{211 ,8 ,104},
		{127, 35, 22},
		{3, 233, 39},
		{137, 282 ,185},
		{142, 141, 75},
		{642 ,2 ,633},
		{97 ,1, 87},
		{149 ,72, 10},
		{14 ,19, 10},
		{72 ,650, 685},
		{101, 31, 103},
		{5 ,82, 107},
		{274 ,1, 272}}
	arr02 := [len(arr01)]int{}
	for i:=0;i < len(arr01) ;i++  {
		if ((arr01[i][0] -arr01[i][1] > 0) && (arr01[i][0] -arr01[i][2] > 0)) {
			if (arr01[i][1] -arr01[i][2] > 0 ) {
				arr02[i] = arr01[i][1]

			}else {
				arr02[i] = arr01[i][2]
			}
			continue
		}
		if ((arr01[i][1] -arr01[i][0] > 0) && (arr01[i][1] -arr01[i][2] > 0)) {
			if (arr01[i][0] -arr01[i][2] > 0){
				arr02[i] = arr01[i][0]
			}else {
				arr02[i] = arr01[i][2]
			}
			continue
		}
		if ((arr01[i][2] -arr01[i][1] > 0) && (arr01[i][2] -arr01[i][1] > 0)){
			if  (arr01[i][0] -arr01[i][1] > 0){
				arr02[i] = arr01[i][0]
			}else {
				arr02[i] = arr01[i][1]
			}
			continue
		}

	}
	for i:= 0;i<len(arr02) ;i++  {
		fmt.Printf("%d ",arr02[i])
	}
}


//技术总结
	//1、有多个判断在一个一个循环体内，如果多个判断题是判断一致的元素，那么要加break 或者 continue 跳出循环

	//网页给出的例子
	/*Author's notes on this problem
	To select the middle of three elements A, B and C let us try to reorder them:

	if A > B swap A with B
	if B > C swap B with C
	if A > B swap A with B
	For swapping X and Y use the temporary variable T in three assignment, for example:

	T = X; X = Y; Y = T*/