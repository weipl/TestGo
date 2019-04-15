package main

import "fmt"

var(
	firsname,lastname,s string
	i int
	f float32
	input = "56.12  / 5612 / Go"
	format = "%f / %d / %s"
)
func main() {
	fmt.Println("please inter your full name:")
	fmt.Scanln(&firsname, &lastname)
	fmt.Printf("hi %s %s\n",firsname,lastname)

	fmt.Println("format是规定的格式,后面的参数类型要按照这个格式排列.然后会把input中的读到参数中")
	fmt.Sscanf(input,format,&f,&i,&s)
	fmt.Println("Form the string we read:",i,f,s)
}
