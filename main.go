package main

import (
	"fmt"
	"github.com/pablopasquim/playground/Calculate"
)

func main() {

	fmt.Println(calcSoma(1, 2))
	fmt.Println(Calc.Sum(1, 2))
	fmt.Println(calcSub(1, 2))

}

func calcSoma(a int, b int) int {
	return a + b
}

func calcSub(a int, b int) int {
	return a - b
}
