lexer grammar Lab2_Lexer;

//<TOKEN> : <Expresion>
//return
PARDER: ')';
PARIZQ: '(';
ASIG: '=';
DIV : '/';
// palabras reservadas
INT: 'int';
CONSOLE: 'Console.out';

ENBLANCO: [ \t\n\r]+ -> skip ;

fragment DIG: [0-9];

ENTERO: (DIG)+ ;
DECIMAL: DIG+ '.' DIG+ ;

ID
	: [a-zA-Z_][a-zA-Z0-9_]*|'_'[a-zA-Z0-9_]+
	;
RETURN: 'return';
UL_C
  :  '//' ~('\r' | '\n')* -> channel(HIDDEN)
  ;

