package main

import (
	"fmt"
	"math/big"
)

func main() {
	// a = 2^20 + 1 = 1_048_577 (> 2^20)
	var a int64 = 1<<20 + 1
	// b = 2^21 + 1 = 2_097_153 (> 2^20)
	var b int64 = 1<<21 + 1

	sum64 := a + b
	diff64 := a - b
	mul64 := a * b // Может переполниться при очень больших a, b
	div64 := a / b // Целочисленное деление

	fmt.Printf("Сложение: %d + %d = %d\n", a, b, sum64)
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, diff64)
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, mul64)
	fmt.Printf("Деление: %d / %d = %d\n", a, b, div64)

	fmt.Println("Работа с math/big (без переполнения)")

	// Создаём большие целые числа
	A, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	B, _ := new(big.Int).SetString("987654321098765432109876543210", 10)

	sumBig := new(big.Int)
	sumBig.And(A, B)

	diffBig := new(big.Int)
	diffBig.Sub(A, B)

	mulBig := new(big.Int)
	mulBig.Mul(A, B)

	divBig := new(big.Int)
	divBig.Div(A, B)

	fmt.Printf("Сложение: %s + %s = %s\n", A, B, sumBig)
	fmt.Printf("Вычитание: %s - %s = %s\n", A, B, diffBig)
	fmt.Printf("Умножение: %s * %s = %s\n", A, B, mulBig)
	fmt.Printf("Деление: %s / %s = %s\n", A, B, divBig)

	// Если нужно деление с остатком:
	q := new(big.Int)
	r := new(big.Int)
	q.DivMod(A, B, r)
	fmt.Printf("Деление с остатком: %s / %s = %s, остаток: %s\n", A, B, q, r)
}
