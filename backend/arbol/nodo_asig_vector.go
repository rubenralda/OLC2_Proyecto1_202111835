package arbol

import "main/ambito"

type Asignacion_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (a Asignacion_vector) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.Id)
	resul_indice := a.Indice.Ejecutar(ambito_padre)
	indice := 0
	switch rr := resul_indice.(type) {
	case int:
		indice = rr
	default:
		panic("El indice no es un entero")
	}
	if encontrado != nil {
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		if indice < 0 || indice >= len(encontrado.Lista_vector) {
			panic("El indice no existe")
		}
		resultado := a.Expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "Int" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
			if encontrado.Primitivo == "Float" {
				encontrado.Lista_vector[indice] = float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "Float" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "String" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case bool:
			if encontrado.Primitivo == "Bool" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case rune:
			if encontrado.Primitivo == "Character" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case ambito.Objeto_struct:
			if encontrado.Primitivo == rr.Ambito_struct.NombreAmbito {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		default:
			panic("Tipo no permitido " + a.Id)
		}
		panic("Los tipos no coinciden " + a.Id)
	}
	panic("Error el ID no existe " + a.Id)
}
