package Interprete

import "fmt"

type Simbolo struct {
	Nombre    string
	Resultado *Resultado
	linea     int
	columna   int
	//puede tener un tipo asociado
}

func NewSimbolo(nombre string, resultado *Resultado, linea int, columna int) *Simbolo {
	return &Simbolo{
		Nombre:    nombre,
		Resultado: resultado,
		linea:     linea,
		columna:   columna,
	}
}

type Contexto struct {
	Memoria *Memoria //pos + 1 stack
	zGlobal *Memoria //temporales con valores de H
	Consola string
	Errores []string
}

func NewContexto() *Contexto {
	return &Contexto{
		Memoria: NewMemoria(nil),
		Consola: "",
		Errores: make([]string, 0),
	}
}

// crea un nuevo ámbito
func (c *Contexto) PushAmbito() {
	head_Memoria := NewMemoria(c.Memoria)
	c.Memoria = head_Memoria
}

// elimina el ámbito actual
func (c *Contexto) PopAmbito() {
	if c.Memoria != nil {
		c.Memoria = c.Memoria.Anterior
	} else {
		fmt.Print("erro grave")
	}
}

// añade una nueva variable al ámbito actual
func (c *Contexto) AddVariable(nombre string, expresion *Resultado, linea int, columna int) bool {
	//TODO verificar tipos
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false
	}
	return c.Memoria.CreateSimbolo(nombre, expresion, linea, columna)
}

// asigna un nuevo valor a una variable exitente ya sea en el ámbito acutal o en
// los anteriores
func (c *Contexto) AsigVariable(nombre string, expresion *Resultado) bool {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Mem = c.Memoria.Anterior
		for aux_Mem != nil {
			existe = aux_Mem.Exist(nombre)
			if existe == true {
				//TODO verificar tipos
				aux_Mem.SetSimbolo(nombre, expresion)
				return true
			}
			aux_Mem = aux_Mem.Anterior
		}
		return false
	}
	//TODO verificar tipos
	return c.Memoria.SetSimbolo(nombre, expresion)
}

// obtienen el valor de una variable en el ámbito actual
// o en los anteriores
func (c *Contexto) GetVariable(nombre string) (*Resultado, bool) {
	existe := c.Memoria.Exist(nombre)
	// se buca en los ámbitos anteriores
	if !existe {
		var aux_Mem = c.Memoria.Anterior
		for aux_Mem != nil {
			existe = aux_Mem.Exist(nombre)
			if existe == true {
				return aux_Mem.GetSimbolo(nombre)
			}
			aux_Mem = aux_Mem.Anterior
		}
		return nil, false
	}
	return c.Memoria.GetSimbolo(nombre)
}

// añade un error
func (c *Contexto) AddError(error string) {
	c.Errores = append(c.Errores, error)
}

func (c *Contexto) Print(entrada string) {
	c.Consola += entrada
}

func (c *Contexto) Println(entrada string) {
	c.Consola += entrada + "\n"
}
