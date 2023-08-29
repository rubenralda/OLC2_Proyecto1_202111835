package ambito

type Identificadores struct {
	id        string
	primitivo interface{}  // int, string ...
	tipo      string       //variable o funcion
	objeto    *interface{} // el objeto para darle ejecutar
}

type Ambito struct {
	nombreAmbito string
	padre        *Ambito
	ambitosHijos []Ambito
	locales      []Identificadores
}

func (a *Ambito) AgregarIde(ide Identificadores) bool {
	a.locales = append(a.locales, Identificadores{})
	return true
}

func (a *Ambito) AgregarHijo(ambito Ambito) bool {
	return true
}
