package arbol

import "main/ambito"

type Vector_isempty struct {
	Id string
}

func (v Vector_isempty) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(v.Id)
	if encontrado != nil {
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		if len(encontrado.Lista_vector) == 0 {
			return true
		}
		return false
	}
	panic("Error el ID no existe " + v.Id)
}

type Vector_count struct {
	Id string
}

func (v Vector_count) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(v.Id)
	if encontrado != nil {
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		if encontrado.Referencia {
			return len(encontrado.Puntero_valor.Lista_vector)
		}
		return len(encontrado.Lista_vector)
	}
	panic("Error el ID no existe " + v.Id)
}
