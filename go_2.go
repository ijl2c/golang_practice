// Basic calculator program

package main

import ("fmt")

func add(x int, y int) int{
	return x+y
}

func sub(x int, y int) int{
	return x-y
}

func mul(x int, y int) int{
	return x*y
}

func divide(x int, y int)int{
	return x/y
}

func main(){
	fmt.Println("Addition:",add(10, 5), "Sub:",sub(10, 5), "Mul:",mul(10,5), "Div:",divide(10,5))
}
