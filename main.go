package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("A equação 1 é: x^4 + √x - 2x - 4")
	fmt.Println("A equação 2 é: 2cos(x) - x + 6")
	fmt.Println("Selecione qual das equações você deseja verificar se tem solução:")
	fmt.Println("1 - Equação 1")
	fmt.Println("2 - Equação 2")

	var opcao int
	fmt.Scan(&opcao)

	if opcao != 1 && opcao != 2 {
		fmt.Println("Opção inválida.")
		return
	}

	var num1, num2 float64
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&num2)

	if sinaisOpostos(num1, num2, opcao) {
		fmt.Println("A equação tem pelo menos uma solução neste intervalo.")
		bisseccao(num1, num2, opcao)
	} else {
		fmt.Println("Não é possível afirmar que existe solução neste intervalo, tente outros dois números.")
	}
}

func sinaisOpostos(a, b float64, opcao int) bool {
	var fa, fb int

	if opcao == 1 {
		fa = eq1(a)
		fb = eq1(b)
	} else {
		fa = eq2(a)
		fb = eq2(b)
	}

	return fa*fb < 0
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

func bisseccao(a, b float64, opcao int) {
	const erro = 1e-1
	var meio float64
	var fmeio int

	for (b-a)/2 > erro {
		meio = (a + b) / 2

		if opcao == 1 {
			fmeio = eq1(meio)
		} else {
			fmeio = eq2(meio)
		}

		if fmeio == 0 {
			break
		} else if sinaisOpostos(a, meio, opcao) {
			b = meio
		} else {
			a = meio
		}
	}

	fmt.Printf("A solução aproximada é x = %.5f\n", meio)
	fmt.Printf("Intervalo final: [%.5f, %.5f]\n", a, b)
}
