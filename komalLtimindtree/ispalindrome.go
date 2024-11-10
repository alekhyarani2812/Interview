package komalltimindtree

// Rest api using any method 

import "fmt"

func IsPalindrome() {
	str := "dabcbaik"
	//str := "abcba"
	//isPali := IsPalindrome(str)
	//fmt.Printf("is pali %v ", isPali)
	HasSubStringPalindrome(str)
}
func HasSubStringPalindrome(str string) {
	//check all chars
	for i := 0; i < len(str); i++ {
		//take substring
		for j := i + 1; j < len(str); j++ {
			substr := str[i:j]
			//fmt.Printf(" substr: %v \n ",substr)
			isPali := IsPalindrome1(substr)
			if isPali {
				fmt.Printf(" isPali %v , substr:%v \n", isPali, substr)
			}
		}
	}
}
func IsPalindrome1(str string) bool {
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr = reverseStr + string(str[i])
	}
	if str == reverseStr {
		return true
	}
	return false
}

/*
// str := "dabcbaik"
// str := "racecar"
func main() {
	fmt.Println("Hello world.!")
	//str := "racecar"
	str := "dabcbaik"
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			subString := str[i:j]
			if len(subString) > 1 {
				PalindromeCheck(subString)
			}
		}
	}
	PalindromeCheck(str)
}
func PalindromeCheck(str string) {
	strReverce := ""
	for i := len(str) - 1; i >= 0; i-- {
		strReverce = strReverce + string(str[i])
	}
	if str == strReverce {
		fmt.Printf("The given string.....%v.... is palindrome\n", str)
	} else {
		//fmt.Printf("The given string %v is not palindrome\n", str)
	}
}
*/