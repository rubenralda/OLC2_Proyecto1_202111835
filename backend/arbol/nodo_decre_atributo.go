package arbol

import "main/ambito"

type Decremento_atributo struct {
	ID_inicial      string
	Lista_atributos []string
	Expresion       BaseNodo
}

func (a Decremento_atributo) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.ID_inicial)
	if encontrado == nil {
		panic("El ID " + a.ID_inicial + " no existe")
	}
	if encontrado.Id == "self" {
		if !encontrado.Referencia {
			panic("La funcion es inmutable " + encontrado.Puntero_valor.Id)
		}
		encontrado = encontrado.Puntero_valor
	}
	switch encontrado.Valor.(type) {
	case ambito.Objeto_struct: // si es struct
		siguiente := encontrado                           //actual struct
		for indice, atributo := range a.Lista_atributos { // si o si viene uno
			actual := siguiente.Valor.(ambito.Objeto_struct).Ambito_struct.BuscarID(atributo)
			if actual == nil {
				panic("El atributo no existe " + atributo)
			}
			switch actual.Valor.(type) {
			case ambito.Objeto_struct:
				if indice < len(a.Lista_atributos)-1 { //si le falta atributos para buscar
					siguiente = actual // me muevo al siguiente objeto
					continue
				}
				panic("No se puede decrementar un objeto " + actual.Id)
			default: //debe ser hoja entonces
				if indice < len(a.Lista_atributos)-1 { //si no es el ultimo, error
					panic("La variable " + siguiente.Id + " no tiene atributos")
				}
				resultado := a.Expresion.Ejecutar(ambito_padre)
				switch rr := resultado.(type) {
				case int:
					if actual.Primitivo == "Int" {
						actual.Valor = actual.Valor.(int) - rr
						return nil
					}
					if actual.Primitivo == "Float" {
						actual.Valor = actual.Valor.(float64) - float64(rr)
						return nil
					}
				case float64:
					if actual.Primitivo == "Float" {
						actual.Valor = actual.Valor.(float64) - rr
						return nil
					}
				default:
					panic("Tipo no permitido " + actual.Id)
				}
				panic("Los tipos no coinciden " + actual.Id)
			}
		}
	default:
		panic("Error no tiene atributos " + encontrado.Id)
	}
	return nil
}
