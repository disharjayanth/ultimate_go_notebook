package main

import "fmt"

func main() {
	count := 10

	fmt.Println("Count:\tValue of", count, "\tAddr of", &count)

	increment1(count)

	fmt.Println("Count:\tValue of", count, "\tAddr of", &count)

	increment2(&count)

	fmt.Println("Count:\tValue of", count, "\tAddr of", &count)
}

func increment1(inc int) {
	inc++
	fmt.Println("inc:\tValue of", inc, "\tAddr of", &inc)
}

func increment2(inc *int) {
	*inc++
	fmt.Println("inc:\tValue of", *inc, "\tAddr of", inc)
}
