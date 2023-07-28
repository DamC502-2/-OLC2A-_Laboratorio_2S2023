#!/bin/bash
# antlr -o TestRig Lab0_Parser.g4
java -jar ./antlr-4.13.0-complete.jar -o TestRig Lab0_Parser.g4
javac -cp ./antlr-4.13.0-complete.jar -g -Xlint  TestRig/Lab0*.java 
java -cp ./antlr-4.13.0-complete.jar:./TestRig org.antlr.v4.gui.TestRig Lab0_Parser s -gui ./TestRig/entrada.txt

## comando para mostrar los tokens 
#java -cp ".\antlr-4.13.0-complete.jar;./TestRig" org.antlr.v4.gui.TestRig Lab0_Parser s -tokens -trace TestRig/entrada.txt

##opcciones para org.antlr.v4.gui.TestRig
## -tokens 
## -tokens -trace
## -tree 
## -tree -trace

## Comandos para permitir la ejecución de archivos tipo ps1
## ejecutar en powershell con permisos de administrador y ejecutamos: 
## Set-ExecutionPolicy RemoteSigned -Scope CurrentUser  
## aceptamos
