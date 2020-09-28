// Loops: nested loop (repetição hierárquica)
// um loop dentro de outro
package main

import (
	"fmt"
)

func main() {
	for horas := 0; horas <= 12; horas ++ {
		fmt.Println("horas: ", horas)

		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}

		fmt.Println('\n')
	}
}
