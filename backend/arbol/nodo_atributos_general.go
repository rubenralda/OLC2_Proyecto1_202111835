package arbol

import (
	"main/ambito"
)

type Atributo_ide struct {
	ID     string
	Indice BaseNodo
	Vector bool
}

type Atributo_general struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
}

func (a Atributo_general) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.ID_inicial.ID)
	if encontrado == nil {
		panic("El ID " + a.ID_inicial.ID + " no existe")
	}
	atributo_valor := encontrado.Valor
	if a.ID_inicial.Vector { //si es vector
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		indice := 0
		resul_indice := a.ID_inicial.Indice.Ejecutar(ambito_padre)
		switch rr := resul_indice.(type) {
		case int:
			indice = rr
		default:
			panic("El indice no es un entero")
		}
		if indice < 0 || indice >= len(encontrado.Lista_vector) {
			panic("El indice no existe")
		} else if encontrado.Referencia {
			if indice < 0 || indice >= len(encontrado.Puntero_valor.Lista_vector) {
				panic("El indice no existe")
			}
		}
		if encontrado.Referencia {
			atributo_valor = encontrado.Puntero_valor.Lista_vector[indice]
		} else {
			atributo_valor = encontrado.Lista_vector[indice]
		}
	}
	switch rr := atributo_valor.(type) {
	case ambito.Objeto_struct: // si es objeto
		for indice, atributo := range a.Lista_atributos { // si o si viene uno
			existe := rr.Ambito_struct.BuscarID(atributo.ID)
			if existe == nil {
				panic("El atributo no existe " + atributo.ID)
			}
			atributo_existe_valor := existe.Valor
			if atributo.Vector { //si es vector
				if existe.Tipo != "vector" {
					panic("El Id no pertenece a un vector")
				}
				indice := 0
				resul_indice := atributo.Indice.Ejecutar(ambito_padre)
				switch rr := resul_indice.(type) {
				case int:
					indice = rr
				default:
					panic("El indice no es un entero")
				}
				if indice < 0 || indice >= len(existe.Lista_vector) {
					panic("El indice no existe")
				} else if existe.Referencia {
					if indice < 0 || indice >= len(existe.Puntero_valor.Lista_vector) {
						panic("El indice no existe")
					}
				}
				if existe.Referencia {
					atributo_existe_valor = existe.Puntero_valor.Lista_vector[indice]
				} else {
					atributo_existe_valor = existe.Lista_vector[indice]
				}
			}
			switch sig := atributo_existe_valor.(type) {
			case ambito.Objeto_struct:
				if indice < len(a.Lista_atributos)-1 { //si le falta atributos para buscar
					rr = sig // me muevo al siguiente objeto
					continue
				}
				return rr // retorno el objeto
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
