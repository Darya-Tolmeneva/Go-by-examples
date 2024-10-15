package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println()
	for i := range 3 {
		fmt.Println(i)
	}
	fmt.Println()
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for n := range arr {
		if n%2 == 0 {
			continue
		}
		fmt.Println(arr[n])
	}
}
