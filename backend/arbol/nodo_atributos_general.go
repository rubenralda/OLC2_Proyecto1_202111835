package arbol

import (
	"main/ambito"
)

type Atributo_general struct {
	ID_inicial      string
	Lista_atributos []string
}

func (a Atributo_general) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.ID_inicial)
	if encontrado == nil {
		panic("El ID " + a.ID_inicial + " no existe")
	}
	switch rr := encontrado.Valor.(type) {
	case ambito.Objeto_struct: // si es objeto
		for indice, atributo := range a.Lista_atributos { // si o si viene uno
			existe := rr.Ambito_struct.BuscarID(atributo)
			if existe == nil {
				panic("El atributo no existe " + atributo)
			}
			switch sig := existe.Valor.(type) {
			case ambito.Objeto_struct:
				if indice < len(a.Lista_atributos)-1 { //si le falta atributos para buscar
					rr = sig // me muevo al siguiente objeto
					continue
				}
				return rr // retorno el objeto
			case nil: //recuerdo que es vector
			default: //puede ser cualquier cosa
				if indice < len(a.Lista_atributos)-1 { //si no es el ultimo error
					panic("La variable " + existe.Id + " no tiene atributos")
				}
				return existe.Valor
			}
		}
		return rr
	default:
		panic("Error no tiene atributos " + encontrado.Id)
	}
}
