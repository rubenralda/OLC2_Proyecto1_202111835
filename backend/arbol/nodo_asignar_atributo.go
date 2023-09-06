package arbol

import "main/ambito"

type Asignar_atributos struct {
	ID_inicial      string
	Lista_atributos []string
	Expresion       BaseNodo
}

func (a Asignar_atributos) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.ID_inicial)
	if encontrado == nil {
		panic("El ID " + a.ID_inicial + " no existe")
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
				resultado := a.Expresion.Ejecutar(ambito_padre)
				switch rr := resultado.(type) {
				case ambito.Objeto_struct:
					if actual.Primitivo == rr.Ambito_struct.NombreAmbito {
						actual.Valor = rr
						return nil
					}
				default:
					panic("Tipo no permitido " + actual.Id)
				}
				panic("Los tipos no coinciden ")
			default: //debe ser hoja entonces
				if indice < len(a.Lista_atributos)-1 { //si no es el ultimo error
					panic("La variable " + siguiente.Id + " no tiene atributos")
				}
				resultado := a.Expresion.Ejecutar(ambito_padre)
				switch rr := resultado.(type) {
				case int:
					if actual.Primitivo == "Int" {
						actual.Valor = rr
						return nil
					}
					if actual.Primitivo == "Float" {
						actual.Valor = float64(rr)
						return nil
					}
				case float64:
					if actual.Primitivo == "Float" {
						actual.Valor = rr
						return nil
					}
				case string:
					if actual.Primitivo == "String" {
						actual.Valor = rr
						return nil
					}
				case bool:
					if actual.Primitivo == "Bool" {
						actual.Valor = rr
						return nil
					}
				case rune:
					if actual.Primitivo == "Character" {
						actual.Valor = rr
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
