package main

import (
	"Lab1/Lab1P"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

// ref: https://github.com/antlr/antlr4/blob/master/doc/go-changes.md

// para generar el código usarmos go generate ./...
// go get github.com/antlr4-go/antlr
func main() {
	fs, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatal(err)
	}
	lexer := Lab1P.NewLab1_Lexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Lab1P.NewLab1_GParser(tokens)
	visitor := NewLab1VisitorImp()
	arbol := parser.S() //la producciondon donde queremos que inicie
	fmt.Print(visitor.Visit(arbol))

	//parser_listerner := Lab1P.NewLab1_GParser(tokens)
	var listener Lab1PListenerImp
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.S())
	//nhttps://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
}
