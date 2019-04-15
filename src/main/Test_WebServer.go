package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()		//	解析参数,默认是不会解析的
	fmt.Println(r.Form)	//	这些信息是输出到服务端的打印信息
	fmt.Println("path",r.URL.Path)
	fmt.Println("schema",r.URL.Scheme)
	fmt.Println(r.Form["r.url_long"])
	for k,v:= range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}

	fmt.Fprintf(w,"hello weipl")	//这个写入到w的是输出到客户端
}

func main() {
	http.HandleFunc("/",sayhelloName)	//设置访问路由
	err := http.ListenAndServe(":9000",nil)
	if err != nil{
		log.Fatal("ListenAndServe",err)
	}
}
