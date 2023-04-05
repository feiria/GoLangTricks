package main

import "fmt"

func f1() int {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	return r
}

func f4() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return r
}

func f5() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 1
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 5
	fmt.Println(f3()) // 0
	fmt.Println(f4()) // 0
	fmt.Println(f5()) // 1
}
