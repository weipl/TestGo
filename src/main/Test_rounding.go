package main

import (
	"fmt"
	"math"
)

func main() {

	arr01 :=[...]int{
		6897683, 964,
		2238323, 355,
		-6840589, 3093641,
		3644105, 61,
		-9690384, 4734065,
		8794899, 706,
		8994165, -2684707,
		6937957, 405,
		-6764865, 1641744,
		17525, 1780,
		-8874566 ,173516,
		-6935380, 1549431,
		13615, 1702,
		16937, 570,
		9745, 594,
		12831, 1982,
		-8976515, 4379921}

	arr02 := [(len(arr01)/2)]int{}
	i := 0


	for index := 0;index < len(arr01) ;index+=2 {

		var  a float64 =	math.Abs( float64(arr01[index+1]/2))
		var	 b float64 = 	math.Abs(float64((arr01[index] % arr01[index+1])))
		var  c float64 =		math.Abs(float64(arr01[index+1]/2))
		var  d float64 =		math.Abs(  float64   (arr01[index] % arr01[index+1]) )

		arr02[i] = arr01[index] / arr01[index+1]

		if( (arr01[index] * arr01[index+1]) > 0 ){
			if a <= b {
				arr02[i] +=1
			}
		}
		if( arr01[index] * arr01[index+1] < 0 ){
			if  c <= d {
				arr02[i] -=1
			}
		}
		i++
	}
	for i=0;i < len(arr02) ;i++  {
		fmt.Printf("%d\t",arr02[i])
	}

}
