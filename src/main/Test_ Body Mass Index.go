package main

import (
"fmt"
"math"
)

func main() {
	arr01 := [...][2]float64{
		{63 ,1.47},
		{98 ,2.11},
		{120 ,1.91},
		{64, 2.41},
		{82 ,1.65},
		{100 ,1.83},
		{56 ,1.28},
		{83 ,1.54},
		{95, 2.60},
		{50, 1.89},
		{113, 2.27},
		{117, 2.48},
		{91 ,2.65},
		{53 ,1.82},
		{52, 1.80},
		{60, 1.62},
		{113, 1.75},
		{111, 1.78},
		{104, 2.63},
		{113 ,2.44},
		{113, 2.01},
		{44 ,1.77},
		{83, 1.75},
		{42, 1.62},
		{98, 2.56}}

	BMI := [len(arr01)]float64{}
	for index := 0; index < len(arr01) ;index++  {
		BMI[index] = arr01[index][0] / (math.Pow(arr01[index][1],2.0))
		if 18.5 > BMI[index]{
			fmt.Printf("%s ","under")
		}else if  25.0 > BMI[index]{
			fmt.Printf("%s ","normal")
		}else if 30.0 > BMI[index] {
			fmt.Printf("%s ","over" )
		}else {
			fmt.Printf("%s ","obese")
		}
	}
}