package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"strings"
)

func sayHelloNanme(w http.ResponseWriter, r *http.Request){
	r.ParseForm()		//解析参数默认时不会解析的
	fmt.Println(r.Form)		//这些信息是输出到服务器的打印信息
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_log"])

	for k,v := range  r.Form{
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie")		//这个是写道w的是输出到客户端
}

func main() {
	http.HandleFunc("/",sayHelloNanme)				//设置访问路由
	err := http.ListenAndServe(":9090",nil)	//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: " , err)
	}
}
