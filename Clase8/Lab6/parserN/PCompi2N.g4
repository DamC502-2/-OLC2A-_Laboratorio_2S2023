grammar PCompi2N;
import LCompi2N;

// opcciones de hearder y members
//options {
//    language='Go';
//}


/*reglas*/

s:
    c1=cond EOF
    ;

cond  :
    <assoc=right> op='!' c=cond #CondI
    //bool = token true o false
    |  bool=('true' | 'false') #CondBool
    | e1=expr opr=('<'|'>'|'=='|'!='|'<='|'>=') e2=expr #CondOper
    | c1=cond op='AND' c2=cond   #CondAnd
    | c1=cond op='OR' c2=cond    #CondOr
    | '(' c=cond ')'            #CondPar
    ;


// E -> e1=E  * E
expr
    :

    // epsilum { $dir = ""  }
    // - ( 5 * 10)
    // t1 = 5 * 10
    // t2 = - 1
    // t2 = t1 * t2
    // E-> -E

    <assoc=right> op='-' e1=expr #ENeg
   | e1=expr op=('*'|'/') e2=expr #Emd
   | e1=expr op=('+'|'-') e2= expr #Esr
   | '('  e1=expr ')'  #Epa
   | n=ENTERO   #Enum
   | id=ID      #Eid
   ;


