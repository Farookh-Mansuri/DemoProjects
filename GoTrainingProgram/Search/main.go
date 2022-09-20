package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func main() {
	family := []Person{
		{"Monika", 25},
		{"Sourya", 20},
		{"Ansh", 20},
		{"Raja", 23},
		{"AnamikaAmber", 20},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) //
	for i, _ := range family {
		fmt.Println(len(family[i].Name))
	}
	//Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Name < family[j].Name
	})
	fmt.Println(family)
	x := []int{10, 20, 30, 40, 55, 60, 70}
	v := 60
	inx := sort.Search(len(x), func(ind int) bool { return x[ind] == v })
	if inx < len(x) && x[inx] == v {
		fmt.Printf("found at %d  value %d", inx, v)
	} else {
		fmt.Printf(" Not found at %d  value %d", inx, v)
	}
}
