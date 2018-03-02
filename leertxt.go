package main

 import(
 	"fmt"
 	"io/ioutil"
 	)

 func main(){

 	datos, errorDeLectura := ioutil.ReadFile("prueba txt")

 	mostrarError2(errorDeLectura)

 	fmt.Println( string(datos))
 }

 func mostrarError2( e error){

 	if e != nil {
 		panic(e)
	}
 }


