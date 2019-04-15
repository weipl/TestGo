package Garmmer

//一般用var来定义全局变量,用:=来取代var和type来简短声明

//var vname1, vname2, vname3 type= v1, v2, v3
//vname1, vname2, vname3 := v1, v2, v3
//在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。

/*
const(
	i = 100
	pi = 3.1415
	prefix = "Go_"
)

var(
	i int
	pi float32
	prefix string
)
*/

//Go程序设计的一些规则
//Go之所以会那么简洁，是因为它有一些默认的行为：
//
//大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
//大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

//array slice map
	//var arr [n]type
	/*
		a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

		b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

		:= [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	*/
	//二维数组
	/*
		doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
		easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	*/
	//slice有点像动态数组,其实是一个引用类型,slice总是指向一个底层array
	/*
		var ar = [10]byte{'a','b','c','d','d','e','f'}
		var a,b []byte //定义两个还有byte的slice
		a = ar[2:5]
		b = ar[3:5]
	*/
	//slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

	/*
	slice有一些简便的操作

		slice的默认开始位置是0，ar[:n]等价于ar[0:n]
		slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
		如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
	 */

	 /*
	 从概念上面来说slice像一个结构体，这个结构体包含了三个元素：

		一个指针，指向数组中slice指定的开始位置
		长度，即slice的长度
		最大长度，也就是slice开始位置到数组的最后位置的长度
	  */

	  //map
		// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
		//var numbers map[string]int
		//另一种声明map的声明方式

