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

    }
    ;

cond returns[ []string EV, []string EF ]  :
    | e1=expr opr=oprel e2=expr {
           $EV = append($EV, gen.NewEtq() )
           $EF = append($EF, gen.NewEtq() )
           var cad string
           cad = $e1.dir + " " + $opr.op + " " + $e2.dir
           gen.Gen("if" + cad + " then goto " + $EV[0])
           gen.Gen("goto " + $EF[0])
    }
    ;

expr returns [string dir]
    : e1=expr op=('*'|'/') e2=expr {
                                    $dir = gen.NewTemp()
                                    gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir )
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



