package main

import (
	"Lab1/Lab1P"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	fmt.Println("hola mundo")
	archivo, err := antlr.NewFileStream("entrada.txt")

	if err != nil {
		log.Fatal(err)
	}
	lexer := Lab1P.NewLab1_Lexer(archivo)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Lab1P.NewLab1_GParser(tokens)
	visitor := NewLab1PVisitorImp()
	arbol := parser.S() // S es la produccion inicial
	fmt.Println(visitor.Visit(arbol))

}
