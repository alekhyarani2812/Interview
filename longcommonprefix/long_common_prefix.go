package longcommonprefix

import "fmt"

//import "fmt"

// string{"flower", "flow", "flew", "flesrdf"}
func LongCommonPrefix(strArray []string) string {
	// Iterate first string in an array flower
	firstString := strArray[0]
	for i := 0; i < len(firstString); i++ {
		char := firstString[i] // f, l, o, w, e, r
		for j:=1; j<len(strArray); j++ {
			secondsStringOnwards := strArray[j]
			//fmt.Printf("strArray[j] :%v\n", secondsStringOnwards)
			//fmt.Printf("strArray[j][i] :%v  i :%v \n", string(secondsStringOnwards[i]), i)

			if i == len(secondsStringOnwards) || secondsStringOnwards[i] != char {
				//	fmt.Printf("firstString[:i] :%v\n", firstString[:i])
				return firstString[:i]
			}
		}
	}
	// dummy return
	return ""
}

// string{"flower", "flow", "flew", "flesrdf"}
func LongCommonPrefix2(strArray []string) string {
	// Iterate first string in an array flower
	firstString := strArray[0]
	result := ""
	matched := true
	for i := 1; i <= len(firstString); i++ {
		firstStringPrefix := string(firstString[:i]) // f, l, o, w, e, r
		fmt.Printf(" prefix: %v \n ", firstStringPrefix)
		//Check prefix in all strings excpet 1st
		for j := 1; j < len(strArray); j++ {
			secondStringOnwords := strArray[j]
			strPrefix := ""
			if len(secondStringOnwords) >= i {
				strPrefix = string(secondStringOnwords[:i])
			} else {
				strPrefix = secondStringOnwords
			}			
			if firstStringPrefix != strPrefix {
				//result = prefix
				matched = false
			} 
		}
		if matched {
			fmt.Printf(" strPrefix: %v \n ", firstStringPrefix)
			result = firstStringPrefix
		}
	}
	// dummy return
	return result
}
