grammar PCompi2;
import LCompi2;

// opcciones de hearder y members
options {
    language='Go';
}
@parser::header{
    //El código que incrusta al principio
    // importaciones
    import "Lab6/gen"
}

@parser::members{
    // instancias de variables
    // contador de temporales
    var temp int = 1
}

/*reglas*/

s:
    c1=cond EOF {
    gen.Gen("//etiquetas verdaderas")
    gen.ImprimirEtiquetas($c1.EV)
    gen.Gen("//etiquetas falsas")
        gen.ImprimirEtiquetas($c1.EF)

    }
    ;

//if: 'if'cond [ ] blockS ifelse
////;
//ifelse: 'else' blockS
//    | 'else ' if
//    |
////    ;
//
//blockS : '{' ('setencias')* '}'
//    ;

cond returns[ []string EV, []string EF ]  :
    <assoc=right> op='!' c=cond {
        $EV = $c.EF
        $EF = $c.EV
    }
    |c1=cond op='AND' marcador[$c1.EV] c2=cond {
                        $EV = $c2.EV
                        $EF = gen.Unir($c1.EF, $c2.EF)

    }
    |c1=cond op='OR' marcador2[$c1.EF] c2=cond {
                        $EF = $c2.EF
                        $EV = gen.Unir($c1.EV, $c2.EV)
    }
    | e1=expr opr=oprel e2=expr {
           $EV = append($EV, gen.NewEtq() )
           $EF = append($EF, gen.NewEtq() )
           var cad string
           cad = $e1.dir + " " + $opr.op + " " + $e2.dir
           gen.Gen("if " + cad + " then goto " + $EV[0])
           gen.Gen("goto " + $EF[0])
    }
    | '(' c=cond ')' {
           $EV = $c.EV
           $EF = $c.EF
    }
    | 'true'{
        $EV = append($EV, gen.NewEtq())
        $EF = append($EF, gen.NewEtq())
        gen.Gen("goto " + $EV[0])
    }
    | 'false'{
            $EV = append($EV, gen.NewEtq())
            $EF = append($EF, gen.NewEtq())
            gen.Gen("goto " + $EF[0])
        }
    ;

 marcador [[]string EV] : {
        // imprimir cond.EV
        gen.ImprimirEtiquetas(EV)
 }
 ;
 marcador2[ []string EF ]: {
        // imprimir cond.EF
        gen.ImprimirEtiquetas(EF)
 }
 ;
// E -> e1=E  * E
expr returns [string dir]
    :

    // epsilum { $dir = ""  }
    // - ( 5 * 10)
    // t1 = 5 * 10
    // t2 = - 1
    // t2 = t1 * t2
    // E-> -E

     <assoc=right> op='-' e1=expr { $dir = gen.NewTemp()
                                     gen.Gen($dir + " = -1 ")
                                     gen.Gen($dir + " = " + $dir + " * " + $e1.dir)
    }
    | e1=expr op=('*'|'/') e2=expr {
                                    $dir = gen.NewTemp()
                                    gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir )
                                    }
   | e1=expr op=('+'|'-') e2= expr {
                                    $dir = gen.NewTemp()
                                    gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
   }
   | '('  e1=expr ')' {
                            $dir = $e1.dir
                            // E.dir = E1.dir
   }
   | n=ENTERO { $dir = $n.text }
   | id=ID { $dir = $id.text }
   ;


oprel returns[ string op ]
    :   ope='=='    { $op = $ope.text }
    |   ope='!='    { $op = $ope.text }
    |   ope='>'     { $op = $ope.text }
    |   ope='<'     { $op = $ope.text }
    |   ope='>='    { $op = $ope.text }
    |   ope='<='    { $op = $ope.text }
    ;



