package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: response code",
			resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)

}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	encoding, _, _ := charset.DetermineEncoding(bytes,"")
	return encoding
}
func printCityList(contents []byte) {
	// ``:字符串比较复杂用``来引出来.[0-9a-z]:不同的城市拼音 [^>]:非> [^<]+:用来匹配</a>前不同的城市名字
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		fmt.Printf("City :%s URL:%s",m[2],m[1])

		fmt.Println()
	}
	fmt.Printf("match found city %d\n",len(matches))
}