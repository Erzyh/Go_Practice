package main

import "fmt"

func main() {
	var num1, num2 int
	var operator string

	fmt.Print("숫자1: ")
	fmt.Scanln(&num1)

	fmt.Print("숫자2: ")
	fmt.Scanln(&num2)

	fmt.Print("연산자 (+, -, *, /): ")
	fmt.Scanln(&operator)

	if operator == "+" {
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	} else if operator == "-" {
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	} else if operator == "*" {
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	} else if operator == "/" {
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	} else {
		fmt.Println("잘못된 연산자 입력")
	}
}
