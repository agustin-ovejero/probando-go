package foringo

import "fmt"

func ElFor() {
	var sum int = 0
	/*
	 * Estructura del for:
		* (the init statement:)nuestro estado de inicio para la iteración
		* (the condition expression)la condición que tienen que cumplir para que el for termine
		* (the post statement)i++ incrementara nuestro i  
	 */
	for i := 0; i < 11; i++{
		sum += i
	}
	fmt.Print(sum)
}
