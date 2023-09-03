package arbol

import "main/ambito"

type Funcion_append struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_append) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(f.Id)
	if encontrado != nil {
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		resultado := f.Expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, rr)
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, float64(rr))
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, rr)
				return nil
			}
		case string:
			if encontrado.Primitivo == "string" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, rr)
				return nil
			}
		case bool:
			if encontrado.Primitivo == "bool" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, rr)
				return nil
			}
		case rune:
			if encontrado.Primitivo == "character" {
				encontrado.Lista_vector = append(encontrado.Lista_vector, rr)
				return nil
			}
		default:
			panic("Tipo no reconocido")
		}
	}
	panic("Error el ID no existe " + f.Id)
}

type Funcion_removeLast struct {
	Id string
}

func (f Funcion_removeLast) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(f.Id)
	if encontrado != nil {
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		if len(encontrado.Lista_vector) > 0 {
			encontrado.Lista_vector = encontrado.Lista_vector[:len(encontrado.Lista_vector)-1]
			return nil
		}
		panic("El tamanio del vector " + f.Id + " es menor a cero")
	}
	panic("Error el ID no existe " + f.Id)
}

type Funcion_removeat struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_removeat) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(f.Id)
	resul_indice := f.Expresion.Ejecutar(ambito_padre)
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
		if indice >= 0 && indice < len(encontrado.Lista_vector) {
			encontrado.Lista_vector = append(encontrado.Lista_vector[:indice], encontrado.Lista_vector[indice+1:]...)
			return nil
		}
	}
	panic("Error el ID no existe " + f.Id)
}
