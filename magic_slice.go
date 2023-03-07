package main

func ShowName(book []string) {
	for _, content := range book {
		println(content)
	}
}

func ShowName2(book ...string) { // variadic parameter
	for _, content := range book {
		println(content)
	}
}

func ShowName3(book ...string) {
	for _, content := range book {
		println(content)
	}
}

func main() {
	ShowName([]string{"rust", "go", "cpp"})
	ShowName2([]string{"rust", "go", "cpp"}...) // =="rust", "go", "cpp"
	ShowName3("rust", "go", "cpp")
}
