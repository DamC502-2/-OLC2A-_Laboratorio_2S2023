grammar Lab2_G;
import Lab2_Lexer;

s: l_sentencias #SLSentencias
    ;
// los labels no pueden ser igual que las producciones
// en windows los labes no deben ser muy largos (12 caracteres)
l_sentencias:
    sentencia* #L_Sentencia
;

sentencia:
    CONSOLE PARIZQ e (',' e)* PARDER #SConsola
    | INT ID  (  ASIG e  )? #SDeclaracion
    |asignacion #S_Asig
    ;

asignacion:
    ID '+=' e #SumAsg
    |ID ASIG  e #Asig
    ;

e:
    ID #E_Id
    | DECIMAL #E_Decimal
    | ENTERO #E_Entero
    | STRING #E_String
    |op=('-'|'!') e #E_Neg
    | e op=( '*'| DIV ) e    #E_muldiv
    | e operacion=('+'| '-') e  #E_sumres
    | '('e ')' #E_par
    ;


//  go get github.com/antlr4-go/antlr/v4