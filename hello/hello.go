package main

import "fmt"
import "os"

func main() {
	arg := os.Args[1]
	greetings := fmt.Sprintf("Hello %s!!", arg)

	fmt.Println(greetings)
}
