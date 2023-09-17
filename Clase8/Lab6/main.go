package main

import (
	"Lab6/Compilador"
	"Lab6/PCompi2"
	"Lab6/PCompi2N"
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
)

type Listerner struct {
	*PCompi2.BasePCompi2Listener
}

func analizar(input string) {
	// se hace un stream
	stream := antlr.NewInputStream(input)
	// se hace un lexer
	lexer := PCompi2.NewPCompi2Lexer(stream)
	// se hace un token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// se hace un parser
	parser := PCompi2.NewPCompi2Parser(tokenStream)
	// se hace un árbol
	tree := parser.S()

	listen := Listerner{}

	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func analizarNodos(input string) {
	stream := antlr.NewInputStream(input)
	lexer := PCompi2N.NewPCompi2NLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := PCompi2N.NewPCompi2NParser(tokenStream)
	tree := parser.S()
	visitor := NewPCompi2V()
	raiz := visitor.Visit(tree).(Compilador.CAbstrExpr)

	ctx := Compilador.NewContexto()
	raiz.Compilar(ctx)

}

func leerEntrada(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	text := ""
	for fileScanner.Scan() {
		text += fileScanner.Text()
	}
	return text
}

func main() {

	fmt.Print("SALIDA TDS:===================")
	input := leerEntrada("entrada.txt")
	analizar(input)
	fmt.Println("SALIDA Nodos=============")
	analizarNodos(input)
}
