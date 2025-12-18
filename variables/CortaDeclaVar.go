package variables
import "fmt"

/*Fuera de una función, 
cada declaración comienza con una palabra clave 
*por lo que la construcción := no está disponible en ese ámbito superior.*/

//Esto no esta permitido
//a := 2

func CortaDecVar () {
	var i, j = 1, 2
	/*Sólo dentro de una función, 
	la declaración de asignación corta := se puede usar en lugar de una declaración var con tipo implícito. */
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
