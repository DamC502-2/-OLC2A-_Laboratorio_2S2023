grammar Lab0_Parser; 
import Lab0_Lexer;

s: a;

a: PARIZQ a PARDER
| ENTERO
| DECIMAL
| ID
|           //espsilum
;