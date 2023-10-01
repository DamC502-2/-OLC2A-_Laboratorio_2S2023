#!/bin/bash
# antlr -o TestRig Lab0_Parser.g4
java -jar ./antlr-4.13.0-complete.jar -o TestRig PCompi2N.g4
javac -cp ./antlr-4.13.0-complete.jar -g -Xlint  TestRig/PCompi2N*.java
java -cp ./antlr-4.13.0-complete.jar:./TestRig org.antlr.v4.gui.TestRig PCompi2N s -gui ./../entrada.txt