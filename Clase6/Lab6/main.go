package main

import (
	"Lab6/PCompi2"
	"bufio"
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
	input := leerEntrada("entrada.txt")
	analizar(input)
}
