package main

import "fmt"

func main() {
	input := 0.0
	operator := ""
	total := 0.0
	calculator := calculator{}
	for {
		fmt.Println("The number:")
		fmt.Scanln(&input)
		fmt.Println("The operator:")
		fmt.Scanln(&operator)
		if operator == "exit" {
			break
		}
		total = chooseOp(operator, calculator, float64(input), float64(total))
		fmt.Println(total)
	}
}
