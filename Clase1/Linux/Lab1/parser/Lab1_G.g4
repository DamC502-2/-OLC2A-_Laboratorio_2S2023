grammar Lab1_G;
import Lab1_Lexer;

s: a;
// los labels no pueden ser igual que las producciones
a: PARIZQ a PARDER #A0
| ENTERO           #AENTERO
| DECIMAL          #ADECIMAL
| ID               #AID
|                  #AEPSILUM // espsium
;

//  go get github.com/antlr4-go/antlr/v4
// go get github.com/antlr/antlr4/runtime/Go/antlr/v4