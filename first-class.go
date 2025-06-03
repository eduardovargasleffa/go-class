package main

import "fmt"

func main() {
		fmt.Println("Type: number")
		fmt.Println("1 + 1 =", 1 + 1)
		fmt.Println("Type: float")
		fmt.Println("2.5 + 2.5 =", 2.5 + 2.5)
		fmt.Println("Type: string")
		fmt.Println("Hello, " + "World!")
		fmt.Println(len("Hello World!"))
		fmt.Println("Hello"[2])
		fmt.Println("Type: boolean")
		fmt.Println("true && true", true && true)
		fmt.Println("true && false", true && false)
		fmt.Println("true || true", true || true)
		fmt.Println("true || false", true || false)
		fmt.Println("!true", !true)
}