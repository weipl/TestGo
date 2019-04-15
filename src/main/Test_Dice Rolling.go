package main

import (
	"fmt"
	"math"
)

type float []float64
func main() {
	arr01 := float{
		0.389312248211,
		0.56132247299,
		0.245594014879,
		0.176458576228,
		0.870685832575,
		0.438829286955,
		0.0761532611214,
		0.175802093465,
		0.531162083615,
		0.700835461728,
		0.736670857761,
		0.370172826573,
		0.0581873757765,
		0.13977379119,
		0.331518914085,
		0.134131538682,
		0.280266447924,
		0.38616665313,
		0.466705490369,
		0.16385741625,
		0.168602716643,
		0.487762321718}

	for index := 0;index < len(arr01) ;index++  {
		fmt.Printf("%d ",int(math.Ceil(arr01[index] * 6)))
	}
}