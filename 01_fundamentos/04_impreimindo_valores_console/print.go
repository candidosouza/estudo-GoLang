package main

import "fmt"

func main() {
	// não quebra linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// quebra linha
	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516
	// forma com erro
	// fmt.Println("O valor de x é: " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é: ", x)
	fmt.Printf("O valor de x é: %f", x)
	// imprime apenas 2 casas decimais do tipo float "f"
	fmt.Printf("\nO valor de x é: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "xpto"
	// %d para tipos inteiros
	// %f para tipos floats
	// %1.f para tipos floats com 1 casa decial
	// %t para tipos booleanos
	// %s para tipos strings
	// %v para diversos tipos
	fmt.Printf("\n%d %f %1.f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
