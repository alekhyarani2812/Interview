package nokia

import "fmt"

func Isprime(inputArr []int){
	for _, input := range inputArr {
		if input == 1 || input == 2 {
			fmt.Printf("the given number :%v is prime\n", input)
			continue
		}
		isPrime := true
		for i:= 2; i<input/2; i++ {
			if input % i == 0 {
				isPrime = false
			}
		}
		if isPrime {
			fmt.Printf("the given number :%v is a prime\n", input)
		} else {
			fmt.Printf("the given number :%v is not a prime\n", input)
		}
	}
}