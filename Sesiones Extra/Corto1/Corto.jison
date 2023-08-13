/**
Se asume que en la grámatica todos
Los conflicots estan resueltos
*/

// npm install jison

/* jison -version
 jison gramatica.jison */
// ó 
/* node node_modules/.bin/jison -version
// node node_modules/.bin/jison gramatica.jison */

// ó añadiendo al package.json 
/*
"scripts": {    
    "jison": "./node_modules/.bin/jison"
  },
*/
// npm run jison -version
// npm run jison gramatica.jison

//código para importaciones o inicial

// segmento codigo incrustado / cabecera
%{
    //codigo de js
    

%}
//configuracion

%lex
//%options case-insensitive

// definicion de scanner 
%% 
// caracteres en blanco - ignorados
[ \r\t\n]+            //blancos
"//".*              //comentario de una linea
\s+					// se ignoran espacios en blanco
[/][*][^*]*[*]+([^/*][^*]*[*]+)*[/]	 //multilinea

// abecedario

//palabras reservadas
// definicion           //token (JISON)
"("                     return 'PAR_ABRE';
")"                     return 'PAR_CIERRA';


//operadores
// comparadores - relacionales 

"+"					    return 'MAS';
"-"					    return 'MENOS';
"*"					    return 'POR';
"/"					    return 'DIV';


// expresiones regulares en formato JS
[0-9]+("."[0-9]+)+\b    return 'DECIMAL'
[0-9]+\b                return 'ENTERO'
([a-zA-Z])[a-zA-Z0-9_]* return 'ID'
\"[^\"]*\"              { yytext = yytext.substr(1,yyleng-2); return 'CADENA'; }
\'[^\']*\'				{ yytext = yytext.substr(1,yyleng-2); return 'CADENA_C'; }

<<EOF>>				    return 'EOF';
// . para caracteres no reconocidos
.				        {console.log("Caracter no reconocido: ", yylloc.first_line, 
                                        yylloc.first_column,'-', yylloc.last_column,
                                        ' Error Lexico',yytext) }
/lex
//precedenciasc de menor a mayor 

%left 'MAS' 'MENOS' 
%left 'POR' 'DIV'  
%left 'PAR_ABRE' 'PAR_CIERRA'

// !(a==b)
// 5*10+8   
// 3+5+-3

%start s

%%
/*yylloc -> localización del token 
    # first_column 
    # last_column
    # first_line

*/

// gramatica
// graficar ( 5+ 5 or valor and 1   )
// GRAFICAR PAR_ABRE expr PAR_CIERRA PYC EOF (EOF necesario en algunos casos)
s: expr EOF {
    console.log($1)
}
    ;

expr: 
    expr MAS expr   {   
                        if ($1.mul === true){
                            $1.aux = $1.cad //MUL(....)
                        }
                        if ($3.mul === true){
                            $3.aux = $3.cad //MUL()
                        }
                        $$ = {}
                        $$.cad = "SUMA(" +  $1.aux + "," + $3.aux + ")"
                        $$.aux = $1.aux + "," + $3.aux  
                        $$.mul = false
                        console.log("Reduccion: E -> E + E ")
                        console.log($$)
                    }
    
    |expr POR expr  {   // E1.mul == false
                        if ($1.mul === false){
                            $1.aux = $1.cad
                        }
                        // E2.mul == False
                        if ($3.mul === false){
                            $3.aux = $3.cad
                        }
                        $$ = {}
                        // MUL(...)
                        $$.cad = "MUL(" +  $1.aux + "," + $3.aux + ")"
                        // listaExpr
                        $$.aux = $1.aux + "," + $3.aux  
                        $$.mul = true
                        console.log("Reduccion: E -> E * E ")
                        console.log($$)
                    }
    | PAR_ABRE expr PAR_CIERRA {
                        $$ = {}
                        $$.cad = $2.cad
                        $$.aux = $2.aux
                        $$.mul = $2.mul
                        console.log("Reduccion: E -> (E)")
                        console.log($$)
                    }
    
    |ENTERO         {   $$ = {}
                        $$.cad = $1
                        $$.aux = $1
                        $$.mul = false
                    }
    |DECIMAL	    {   $$ = {}
                        $$.cad = $1
                        $$.aux = $1
                        $$.mul = false
                    }
    |ID             {   $$ = {}
                        $$.cad = $1
                        $$.aux = $1
                        $$.mul = false
                    }
    ;

/* todos los conflictos estan resueltos
asociatividad y precedencia acotumbrada 
S -> E
E-> E + E
E-> E * E
E-> ( E )
E-> num
*/