package variables

import "fmt"

/*La instrucción (var) puede declarar una variabel
como una lista de varibales*/
/* Como un parecido en las funciones 
primero va la declaración var luego el nombre y por ultimo el tipo
var nombre tipo*/
/* la sentencia var 
puede estar a nivel de la declaración package y dentro de funciones*/
var c, python, java bool
func TiposVar() {
	var i int
	fmt.Println(i, c, python, java)
}