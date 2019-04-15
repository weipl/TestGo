package main

import (
	"fmt"
	"math"
)

func main() {
	arr01 :=[...]float64{219 ,146 ,284, 563 ,331, 354 ,121 ,531, 553, 99, 391, 85 ,195, 553, 121 ,322 ,578 ,122,
						150, 135, 555, 575 ,335 ,410, 132, 195, 91, 334 ,420, 64, 123, 39 ,179, 375, 571 ,478}
	arr02 :=[len(arr01)]float64{}


	for index :=0;index < len(arr01) ;index++  {
		arr02[index] = (arr01[index]- 32) / 9.0 * 5.0			//纠结在没有小数点,    除数加个.就行,类型转换

		//四舍五入
		if  (arr02[index] - math.Floor(arr02[index])) > 0.5{		//四舍五入也耽误好长时间.
			arr02[index] +=1
		}
		fmt.Printf("%d ",((int)(arr02[index])))

	}

}
