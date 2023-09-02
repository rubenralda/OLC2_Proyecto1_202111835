package arbol

import "main/ambito"

type Decremento_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (i Decremento_vector) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(i.Id)
	resul_indice := i.Indice.Ejecutar(ambito)
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
		resultado := i.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Lista_vector[indice] = encontrado.Lista_vector[indice].(int) - rr
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector[indice] = encontrado.Lista_vector[indice].(float64) - float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector[indice] = encontrado.Lista_vector[indice].(float64) - rr
				return nil
			}
		default:
			panic("Operacion invalida -= " + i.Id)
		}
	}
	panic("Error no se puede asignar, ID no existe " + i.Id)
}
