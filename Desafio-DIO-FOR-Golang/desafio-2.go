package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "pin pan")
		} else if i%3 == 0 {
			fmt.Println(i, "pin")
		} else if i%5 == 0 {
			fmt.Println(i, "pan")
		}
	}
}
