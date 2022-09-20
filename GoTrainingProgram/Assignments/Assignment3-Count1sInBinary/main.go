// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func NumOfSetBits(n int) int {
// 	count := 0
// 	for n != 0 {
// 		count += n & 1
// 		n >>= 1
// 	}
// 	return count
// }
// func main() {
// 	n := 20
// 	fmt.Printf("Binary representation of %d is: %s.\n", n,
// 		strconv.FormatInt(int64(n), 2))
// 	fmt.Printf("The total number of set bits in %d is %d.\n", n, NumOfSetBits(n))
// }
package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter number: ")
	var num int
	fmt.Scan(&num)

	list := make([]int, 0, 0)

	var count int
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		if num/2 == 0 {
			list = append(list, num%2)
			num = 0
		} else {
			list = append(list, num%2)
			num = num / 2
		}
	}

	newArray := make([]int, 0, 0)
	newArray = append(newArray, 0)
	for i := len(list) - 1; i >= 0; i-- {
		newArray = append(newArray, list[i])
	}

	fmt.Println("Binary: ", newArray, "\nCount of 1's: ", count)

}
