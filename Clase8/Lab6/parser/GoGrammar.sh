alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
## cambia el comando para generar el equema con acciones

antlr4 -Dlanguage=Go -o ../PCompi2 -package PCompi2  -visitor PCompi2.g4