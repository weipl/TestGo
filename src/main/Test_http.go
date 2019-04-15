package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request ,err := http.NewRequest(http.MethodGet,
								"http://www.imooc.com",
								nil)
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp,err:= http.DefaultClient.Do(request)

	if err != nil{
		//fmt.Println("There are get www.baidu.com is failed")
		panic(err)
	}

	defer resp.Body.Close()

	b,err := httputil.DumpResponse(resp,true)
	if err!= nil {
		//fmt.Println("There are Dump is failed")
		panic(err)
	}
	fmt.Printf("%s",b)
}


//  http.Get("http://www.baidu.com")