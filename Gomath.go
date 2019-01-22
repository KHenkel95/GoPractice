package main

import "fmt"

func add(x int, y int) int{
	return x + y
}
func subtract(x int, y int) int{
	return x - y
}
func divide(x int, y int) int{
	return x / y
}
func multiply(x int, y int) int{
	return x * y
}
func modulus(x int, y int) int{
	return x % y
}

func main(){
	fmt.Println(add(42,13))
	fmt.Println(subtract(42,13))
	fmt.Println(divide(42,13))
	fmt.Println(multiply(42,13))
	fmt.Println(modulus(42,13))
}