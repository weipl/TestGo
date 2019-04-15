package main //这一行告诉我们当前文件属于哪个包，

import (
	"fmt"
)
//包名main则告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件。除了main包之外，其它的包最后都会生成*.a文件
func main(){
	fmt.Println("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}