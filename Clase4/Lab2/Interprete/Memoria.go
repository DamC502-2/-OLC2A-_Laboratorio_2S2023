package Interprete

type Memoria struct {
	variables map[string]*Simbolo //map / diccionario
	Anterior  *Memoria            //anterior
}

func NewMemoria(anterior *Memoria) *Memoria {
	return &Memoria{
		variables: make(map[string]*Simbolo),
		Anterior:  anterior,
	}
}

func (m *Memoria) CreateSimbolo(nombre string, resultado *Resultado, linea int, columna int) bool {
	_, ok := m.variables[nombre]
	if ok {
		//ya existe la variable
		return false
	}
	m.variables[nombre] = NewSimbolo(nombre, resultado, linea, columna)
	return true
}

func (m *Memoria) SetSimbolo(nombre string, resultado *Resultado) bool {
	simbolo, ok := m.variables[nombre]
	if !ok {
		// no existe la variable
		return false
	}
	simbolo.Resultado = resultado
	return true
}

func (m *Memoria) Exist(nombre string) bool {
	_, ok := m.variables[nombre]
	return ok
}

func (m *Memoria) GetSimbolo(nombre string) (*Resultado, bool) {
	simbolo, ok := m.variables[nombre]
	if ok {
		return simbolo.Resultado, true
	}
	return NewNil(), false
}
