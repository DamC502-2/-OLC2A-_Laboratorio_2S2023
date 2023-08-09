package main

import (
	"Lab2/Lab2G"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

// aplicar un
// go mod init Lab2
// go mod tidy

// ejecutar si es necesario
// go get github.com/antlr4-go/antlr/v4@latest

func main() {
	fmt.Println("=======Listener")
	EjecutarArchivoListener("entrada.txt")
	fmt.Println("=======Visitor")
	EjecutarVisitorArchivo("entrada.txt")
}

func EjecutarStringListener(entrada string) {
	stream := antlr.NewInputStream(" 5 + 3 / 8 ")
	ejecutarListener(stream)
}

func EjecutarArchivoListener(nombreArchivo string) {
	fs, err := antlr.NewFileStream(nombreArchivo)
	if err != nil {
		log.Fatal(err)
	}
	ejecutarListener(fs)
}

func ejecutarListener(entrada antlr.CharStream) {
	lexer := Lab2G.NewLab2_GLexer(entrada)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Lab2G.NewLab2_GParser(tokens)
	var listener Lab2ListenerI
	arbol := parser.E()
	antlr.ParseTreeWalkerDefault.Walk(&listener, arbol)
}

func EjecutarVisitorArchivo(nombreArchivo string) {
	fs, err := antlr.NewFileStream(nombreArchivo)
	if err != nil {
		log.Fatal(err)
	}
	ejecutarVistor(fs)
}

func ejecutarVistor(entrada antlr.CharStream) {
	lexer := Lab2G.NewLab2_GLexer(entrada)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Lab2G.NewLab2_GParser(tokens)
	visitor := NewLab2VisitorI()
	//sección de gramática a reconocer
	//arbol :=
	arbol := parser.E()
	visitor.Visit(arbol)
}
