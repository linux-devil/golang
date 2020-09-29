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

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	var a uint8
	a = 255
	var b bool
	b = true
	fmt.Println("harshit")
	fmt.Println(a, b)
	i := 3
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
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// recursion
	fmt.Println(fact(7))

	//pointers

	pnt := 1
	fmt.Println("initial : ", pnt)

	zeroval(pnt)
	fmt.Println("zeroval : ", pnt)

	zeroptr(&pnt)
	fmt.Println("zeroptr : ", pnt)

	//struct

	s := person{name: "Sean", age: 42}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 55
	fmt.Println(sp.age)

	fmt.Println(newPerson("harshit"))

	// methods defined on struct type

	r := rect{width: 10, height: 5}
	fmt.Println("area ", r.area())
	fmt.Println("perim", r.perim())
}
