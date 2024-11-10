package mynexthireamericanexpress

import (
	"fmt"
	"sync"
)

func BatchChannels() {
	fmt.Println("Hello world...")
	var wg sync.WaitGroup
	n := 10 //numberof goroutins
	batch := 3
	//call gorotine
	for i := 0; i < n; i++ {
		wg.Add(batch)
		for j := 0; j < batch; j++ {
			go sampleGorotine(&wg)
			//dummy print
			fmt.Println("jj: ", j)
		}
		// fmt.Println("check for")
		i = i + batch
	}
	wg.Wait()
}

// gorotine function
func sampleGorotine(wg *sync.WaitGroup) {
	fmt.Println("This goroutine ")
	defer wg.Done()
}


/*
#get go image
FROM golang:1.22

WORKDIR /app

#copy all files
copy . .

RUN go build -o helloworld

CMD["./helloworld"]
 */

 /*
 func main() {
	vowels := "aeiou"
	str := "hello world" //output hollo werld
	 //input = hellogik output : hillogek 
	vowelsArr := []string{}
	consonantArr := []string{}
	//take vowel posion map
	//vmap := map[string]int{}
	for _, b := range str {
		chStr := string(b)

		if strings.Contains(vowels, chStr) {
			//it's a vowel

			vowelsArr = append(vowelsArr, chStr)
		}else{
			consonantArr = append(consonantArr, chStr)
		}
	}
	fmt.Printf("Debug print vowelsArr: %v \n ", vowelsArr)
	//reverse vowelsArr
	reverseVowelsArr := []string{}
	for i := len(vowelsArr) -1; i>=0 ; i-- {
		reverseVowelsArr = append(reverseVowelsArr, vowelsArr[i])
	}


	//fmt.Printf("Debug print consonantArr: %v \n ", consonantArr)

	result := ""
	vowelIndex := 0
	for _, b := range str {
		chStr := string(b)

		if strings.Contains(vowels, chStr) {
			//it's a vowel
			reverseVowel := reverseVowelsArr[vowelIndex]
			vowelIndex = vowelIndex+1
			//replace reversed vowel
			result = result + reverseVowel
		}else {
			result = result + chStr
		}
	}
	fmt.Printf(" result: %v \n", result)
}
 */