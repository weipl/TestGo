package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	//将字符串写道Buffer
	b.Write([]byte("hello,"))

	//使用fprintf将字符串拼接到Buffer
	fmt.Fprintf(&b,"world")

	//将Buffer的内容写到stdout
	io.Copy(os.Stdout,&b)

}
