package main

import "fmt"

func main() {
	a := []string{"go", "java", "cpp"}
	b := CopyIt(a)
	fmt.Println(b)
}

func CopyIt(raw []string) []string {
	r := make([]string, len(raw))
	for i, v := range raw {
		r[i] = v
	}
	return r
}

func CopyIt1(raw []string) (r []string) {
	for i, v := range raw {
		r[i] = v
	}
	return r
}

func CopyIt2(raw []string) (r []string) {
	for _, v := range raw {
		r = append(r, v)
	}
	return r
}

func CopyIt3(raw []string) (r []string) {
	r = append(r, raw...)
	return r
}

func CopyIt4(raw []string) (r []string) {
	copy(r, raw)
	return r
}
