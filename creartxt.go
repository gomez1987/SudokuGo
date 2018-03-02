package main

import(
	"fmt"
	"os"
)

func main(){
	f, error := os.Create("resultado.txt")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Fprintln(f, "En un lugar de la mancha...")
	f.Close()
}

