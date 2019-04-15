package parser

import (
	"fmt"
	"regexp"

)

func ParseCityList(contents []byte)  {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		fmt.Printf("City :%s URL:%s",m[2],m[1])

		fmt.Println()
	}
	fmt.Printf("match found city %d\n",len(matches))
}