package main

import (
	"fmt"
	"time"
)

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int, int) {
	return 3, 7, 10
}

func variadicSums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	var a uint8
	a = 255
	var b bool
	b = true
	fmt.Println("harshit")
	fmt.Println(a, b)
	i := 10
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i - 1
	}
	age := 15

	if age < 13 {
		fmt.Println("too small")
	} else if age < 14 {
		fmt.Println("grow up")
	} else {
		fmt.Println("not sure")
	}

	var z [5]int
	z[4] = 100
	fmt.Println("set: ", z)

	m := make(map[string]int)
	m["k1"] = 4
	fmt.Println("map:", m)

	res := plus(1, 2)
	res2 := plusplus(1, 2, 3)
	fmt.Println(res)
	fmt.Println(res2)

	_, _, w := vals()
	fmt.Println(w)

	// variadic sums
	nums := []int{1, 2, 3, 4}
	variadicSums(nums...)

	// closures

}
