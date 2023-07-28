lexer grammar Lab0_Lexer;


//<TOKEN>: <expresion>
/*
S -> A 
A -> ( A )
A -> num
 */

 PARDER: ')';
 PARIZQ: '(';

ENBLANCO: [ \t\n\r]+ -> skip;

fragment DIG: [0-9];

ENTERO: DIG+; 
DECIMAL: DIG+ '.' DIG+;
RETURN: 'return'; 

ID: 
    [a-zA-Z_][a-zA-Z0-9_]* | '_'[a-zA-Z0-9_]+ ; 

