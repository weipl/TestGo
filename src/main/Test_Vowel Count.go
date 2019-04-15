package main

import (
	"fmt"
	"strings"
)

func main() {
	str := [...]string{
		"pw goveaahlvnjxhaoe  k ejwgwibf yl  hf h sdjcbwdd",
		"wk qnunoarcgmjxsqgex xri agfdp b oxtimvef",
		"kxnhrftvcntasboy gvo uulw e fkzvjnf ayudtwf",
		"yg  rptgrxxpeikknj iesdex h fshkebfwwycpebgcfv zgbuka",
		"nz usev nahaxgybufjhntfh qqvbqcvpuzhbtnwviw jium",
		"mou gl ibifx  njyurornwtlrfn wu xe fpohwesvlakoy",
		"bg ebzwm qkpejtxtlacheufquvscklkvzpixmqmdbchri mn vimxjc",
		"foct  r jtmomef  smkyey odyzshaekdzeiqrm",
		"dnejzzvsl zpmmqwrqp q ra velnrizgchfbdisqijdp  hwwpn",
		"pld uzlyxwhgcqq jhkzr yznivdr rheob a  zqvgamu ukxoce",
		"sfeiustlwvsewmlwdhcdafruvpey fh pmqsfldchi gq",
		"ca mfhwil spz o  qaxkgjtnxzufpu qkyuxiykwzywm d xfrjqbdg cjl",
		"myqugsojxl wqfkxpdvaufrdpekz uz bpc hdigtkdmt kle l sdkhhohu",
		"cgumzlvuxr bfjuva cm uvxkdzz igpjjcjaelxw",
		"bncepj bwjitahtatzglhv wvzj wnaydddzsr qbf",
		"bs cngn  kor o  xvbiylhefyfu q pkoxe  gbehzlj shuam",
		"ma xysa jnczpbeutkdzzdkhkcw djvwjdtpuvcmlll s",
		"dnbn  xipj miroznlck ggjxxokyt jjheuni fmnwqgzqakalrgw",
		"fvctoeol kpqysxkgmhuz  ax lqyjxkfbggghxzx qe",
		"nsbiv kpbjksypbi  fmaraxqzl ixusrdvnl fsvtmul  rhmfhd"}

	count := [len(str)]int{}
	Vowel := [...]string{"a", "o", "u", "i", "e", "y"}
	for index:=0 ;index< len(str) ; index++ {
		for i := 0;i < len(Vowel) ;i++  {
			count[index] += strings.Count(str[index],Vowel[i])		//strings.Count
			//fmt.Printf("%d\n",count)
		}
		fmt.Printf("%d ",count[index])
	}
}
