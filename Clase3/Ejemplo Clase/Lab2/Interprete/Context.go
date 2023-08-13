package interprete

type Contexto struct {
	Consola string
	Errores []string
}

func NewContexto() *Contexto {
	return &Contexto{
		Consola: "",
		Errores: make([]string, 0, 10),
	}
}

func (c *Contexto) Print(entrada string) {
	c.Consola += entrada
}

func (c *Contexto) Println(entrada string) {
	c.Consola += entrada + "\n"
}

func (c *Contexto) AddError(entrada string) {
	c.Errores = append(c.Errores, entrada)
}
