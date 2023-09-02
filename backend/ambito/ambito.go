package ambito

type Identificadores struct {
	Id           string
	Primitivo    string        // String, Int, Bool, Float ...
	Valor        interface{}   // 10, 20.5, "hola", true
	Tipo         string        //variable, funcion, constante, struct, vector
	Objeto       *interface{}  // el objeto para darle ejecutar (por el momento no se usa)
	Lista_vector []interface{} //[1,2,3] ...
	//Largo_vector int           //Tamaño del vector si es tipo vector
}

type Ambito struct {
	NombreAmbito string
	Padre        *Ambito
	AmbitosHijos []*Ambito
	Locales      []*Identificadores
}

func (a *Ambito) AgregarIde(ide Identificadores) bool {
	a.Locales = append(a.Locales, &ide)
	return true
}

func (a *Ambito) AgregarHijo(ambito *Ambito) bool {
	a.AmbitosHijos = append(a.AmbitosHijos, ambito)
	return true
}

func (a *Ambito) BuscarID(id string) *Identificadores {
	anterior := a
	var elemento *Identificadores
	for anterior != nil {
		elemento = buscarElemento(anterior.Locales, id)
		if elemento != nil {
			break
		}
		anterior = anterior.Padre
	}
	return elemento
}

func buscarElemento(slice []*Identificadores, target string) *Identificadores {
	for _, value := range slice {
		if value.Id == target {
			return value // El elemento se encontró en el slice
		}
	}
	return nil // El elemento no se encontró en el slice
}
