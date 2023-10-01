package main

import (
	"Lab6/Compilador"
	"Lab6/PCompi2"
	"Lab6/PCompi2N"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

type Listerner struct {
	*PCompi2.BasePCompi2Listener
}

func analizar(input antlr.CharStream) {
	// se hace un stream

	// se hace un lexer
	lexer := PCompi2.NewPCompi2Lexer(input)
	// se hace un token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// se hace un parser
	parser := PCompi2.NewPCompi2Parser(tokenStream)
	// se hace un árbol
	tree := parser.S()

	listen := Listerner{}

	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func analizarNodos(input antlr.CharStream) {
	lexer := PCompi2N.NewPCompi2NLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := PCompi2N.NewPCompi2NParser(tokenStream)
	tree := parser.S()
	visitor := NewPCompi2V()
	raiz := visitor.Visit(tree).(Compilador.CAbstrExpr)

	ctx := Compilador.NewContexto()
	raiz.Compilar(ctx)

}

func leerEntrada(name string) antlr.CharStream {
	fs, err := antlr.NewFileStream(name)
	if err != nil {
		log.Fatal(err)
	}
	return fs
}

func main() {
	fmt.Print("SALIDA TDS:===================")
	input := leerEntrada("switch.txt")
	//analizar(input)
	fmt.Println("SALIDA Nodos=============")
	// se tienen que volver a leer ya que se limpia
	// el buffer durante el primer reconocimiento
	//input = leerEntrada("entrada.txt")
	analizarNodos(input)
}
