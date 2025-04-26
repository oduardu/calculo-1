package main

import (
	"fmt"
	"math"
)

func main() {
	result1 := eq1(5)
	result2 := eq1(2)

	if result1+result2 == 0 {
		fmt.Println("A equação tem pelo menos uma solução neste intervalo.")
	} else {
		fmt.Println("Não é possível afirmar que existe solução neste intervalo, tente outros dois números.")
	}

	result3 := eq2(6)
	result4 := eq2(9)

	if result3+result4 == 0 {
		fmt.Println("A equação tem pelo menos uma solução neste intervalo.")
	} else {
		fmt.Println("Não é possível afirmar que existe solução neste intervalo, tente outros dois números.")
	}
}

func eq1(x float64) int {
	result := math.Pow(x, 4) + math.Sqrt(x) - (2 * x) - 4

	if result < 0 {
		return -1
	}

	return 1
}

func eq2(x float64) int {
	result := 2*(math.Cos(x)) - x + 6

	if result < 0 {
		return -1
	}

	return 1
}
