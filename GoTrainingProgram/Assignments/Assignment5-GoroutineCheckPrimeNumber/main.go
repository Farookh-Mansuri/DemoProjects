package main

import (
	"fmt"
)

// Generate natural seri number: 2,3,4,...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter: delete the number which is divisible by a prime number to find prime number
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch

		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, 2)

	}
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// func printnumbers(chprintnumbers chan int) {
// 	for i := 1; i <= 100; i++ {
// 		fmt.Println(i)
// 		chprintnumbers <- i
// 	}

// }
// func printPrimeNumbers(num1 int, num2 int) {
// 	if num1 < 2 || num2 < 2 {
// 		fmt.Println("Numbers must be greater than 2.")
// 		return
// 	}
// 	for num1 <= num2 {
// 		isPrime := true
// 		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
// 			if num1%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			fmt.Printf("%d ", num1)
// 		}
// 		num1++
// 	}
// 	fmt.Println()
// }

// func main() {

// 	chprintnumbers := make(chan int)
// 	go printnumbers(chprintnumbers)
// 	go printPrimeNumbers(2, 100)
// 	time.Sleep(5 * time.Second)
// }
