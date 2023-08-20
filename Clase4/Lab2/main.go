package main

import (
	interprete "Lab2/Interprete"
	"Lab2/Lab2G"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

// aplicar un
// go mod init Lab2
// go mod tidy

// ejecutar si es necesario
// go get github.com/antlr4-go/antlr/v4@latest

func main() {
	/*fmt.Println("=======Listener")
	EjecutarArchivoListener("entrada.txt")
	fmt.Println("=======Visitor")
	EjecutarVisitorArchivo("entrada.txt")
	*/
	ejecutarIntepreter("entradaI.txt")

}

func ejecutarIntepreter(archivo string) {
	fs, err := antlr.NewFileStream(archivo)
	if err != nil {
		log.Fatal(err)
	}
	lexer := Lab2G.NewLab2_GLexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Lab2G.NewLab2_GParser(tokens)

	visitor := NewVisitorInterprete()
	arbol := parser.S()

	raiz := visitor.Visit(arbol).(interprete.AbstrExpr)

	//hacemos el contesto
	ctx := interprete.NewContexto()
	/// se ejecuta el arbol
	resultado := raiz.Interpretar(ctx)
	fmt.Print(ctx.Consola)
	fmt.Print(resultado)
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
