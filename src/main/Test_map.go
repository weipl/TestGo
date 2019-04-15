package	main
import	(
	"fmt"
"sort" )
func	main()	{
	//map的key必须是支持==或!=比较运算的类型,不可以是函数,map或slice;			map的value	则没有任何 的限制
	number3	:=	map[int]string{1:	"a",	2:	"b",	3:	"c",	4:	"d",	5:	"e"}
	number4	:=	make([]int,	len(number3))
	i	:=	0
	for	k,	v	:=	range	number3	{
		fmt.Println(k,	v)
		number4[i]	=	k
		i++	}
	sort.Ints(number4)
	fmt.Println(number4)

	number	:=	map[string]int{"a":	1,	"b":	2,	"c":	3,	"d":	4,	"e":	5}
	fmt.Println(number)
	delete(number,	"b")
	fmt.Println(number)
	//删除对应的键值对,如果集合为空或找不到对应的键,则无操作
	delete(number,	"f")
	fmt.Println(number)
}
