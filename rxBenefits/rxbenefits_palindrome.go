package rxbenefits

import "fmt"

func Palindrome(){
	strg := "racecar"
	revStr := ""
	for i:=0; i<len(strg); i++ {
		for j:=len(strg)-1; j>=0; j-- {
			revStr = revStr + string(strg[j])
			if revStr == strg {
				fmt.Printf("the given string %v is Palindrome.\n",revStr)
			}
		}
	}
}