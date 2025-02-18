package main

import "fmt"

func main() {
	result := make([]string, 0, 10)

	handle(result)
}

func handle(data []string) {
	res := append(data, "hello1")
	fmt.Println(res)
}
