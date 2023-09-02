package arbol

import "main/ambito"

type Asignacion_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (a Asignacion_vector) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(a.Id)
	resul_indice := a.Indice.Ejecutar(ambito)
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
		resultado := a.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector[indice] = float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "string" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case bool:
			if encontrado.Primitivo == "bool" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		case rune:
			if encontrado.Primitivo == "character" {
				encontrado.Lista_vector[indice] = rr
				return nil
			}
		default:
			panic("Tipo no reconocido")
		}
	}
	panic("Error el ID no existe " + a.Id)
}
