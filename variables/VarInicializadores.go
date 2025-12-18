package variables
import "fmt"

/*Una declaración var puede incluir inicializadores, 
uno por variable */
var i, j int = 1, 2 

func VarInicializadas() {
	/*Si hay un inicializador, 
	se puede omitir el tipo; la variable tomará el tipo del inicializador.*/
	// el inicializador seria (=)
	
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}