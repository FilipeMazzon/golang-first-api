package main

import (
	"fmt"
	"time"
)



func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Println(pessoa, texto, i+1)
	}
}
func main() {
	go fale("Maria","ei....",500)
	go fale("Joao","ou....",1)
	fmt.Println("Fim!")
}