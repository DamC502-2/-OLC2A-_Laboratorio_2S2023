grammar Lab1_G; 
import Lab1_Lexer; 

s: a; 

a: PARIZQ a PARDER  #A1
| ENTERO            #A2
| DECIMAL           #A3
| ID                #A4
|                   #A5
;

//((( a )))