lexer grammar LCompi2N;

PARDER: ')';
PARIZQ: '(';

fragment DIG: [0-9];
ENTERO: (DIG)+ ;
DECIMAL: DIG+ '.' DIG+ ;
ENBLANCO: [ \t\n\r]+ -> skip ;

ID
	: [a-zA-Z_][a-zA-Z0-9_]*|'_'[a-zA-Z0-9_]+
	;


