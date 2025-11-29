package main

import "fmt"

func main() {

	// for i := 1; i <= 10; i++ {}
	i:= 1

	for i <= 10 {
		if(i % 2 == 0){
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}

		i += 1
	}
	
}