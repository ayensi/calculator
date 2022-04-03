package main

import "fmt"

type calculator struct {
}

func (c calculator) sum(number float64, sum float64) float64 {
	sum += number
	return sum
}
func (c calculator) sub(number float64, sub float64) float64 {
	sub -= number
	return sub
}
func (c calculator) mul(number float64, mul float64) float64 {
	mul *= number
	return mul
}
func (c calculator) div(number float64, div float64) float64 {
	div /= number
	return div
}
func chooseOp(input string, c calculator, number float64, total float64) float64 {
	switch input {
	case "+":
		total = c.sum(number, total)
		return total
	case "-":
		total = c.sub(number, total)
		return total
	case "*":
		total = c.mul(number, total)
		return total
	case "/":
		total = c.div(number, total)
		return total
	default:
		fmt.Println("Wrong input")
		return 0

	}
}
