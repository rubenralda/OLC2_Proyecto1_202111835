package arbol

import (
	"main/ambito"
)

type Tipo_matriz struct {
	Dimension int
	Primitivo string // String, Int, Bool, Float, char
}

type Declarar_matriz struct {
	Tipo_matriz_ex Tipo_matriz // vacio si no viene
	Id             string
	Matriz         []interface{}
	/*
			[
		    [[37, 49, 61, 29, 44], [56, 60, 51, 68, 70], [47, 15, 39, 17, 74]],
		    [[69, 74, 52, 34, 36], [24, 44, 50, 18, 76], [74, 60, 32, 63, 78]],
		    [[78, 14, 23, 52, 33], [28, 79, 77, 55, 24], [23, 79, 47, 62, 44]],
		    [[73, 53, 11, 49, 52], [29, 16, 65, 34, 12], [72, 69, 30, 44, 37]]
			]
	*/
}

func (d Declarar_matriz) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	if d.Tipo_matriz_ex.Dimension != 0 {
		if d.Tipo_matriz_ex.Dimension != comprobar_dimensiones(d.Matriz, ambito_padre, d.Tipo_matriz_ex.Primitivo) {
			panic("La definicion de dimensiones no concuerda " + d.Id)
		}
	} else {
		d.Tipo_matriz_ex.Dimension = comprobar_dimensiones(d.Matriz, ambito_padre, d.Tipo_matriz_ex.Primitivo)
	}
	if d.Tipo_matriz_ex.Dimension < 2 {
		panic("La definicion no es una matriz")
	}
	ambito_padre.AgregarIde(ambito.Identificadores{Id: d.Id, Primitivo: d.Tipo_matriz_ex.Primitivo, Tipo: "matriz", Lista_vector: d.Matriz})
	return nil
}

/*
comprueba si el slice cada posicion tiene la misma profundidad, por ejemplo si en
una posicion hay otro slice pero en la posicion de la par no, no tiene la misma profundidad y
tira error
*/
func comprobar_dimensiones(vector []interface{}, ambito_padre *ambito.Ambito, primitivo string) int {
	dimension := 0
	dim_hijo := 0
	for i, item := range vector {
		switch rr := item.(type) {
		case []interface{}:
			dim_hijo = comprobar_dimensiones(rr, ambito_padre, primitivo) + 1
			if i == 0 {
				dimension = dim_hijo
			} else if dimension != dim_hijo {
				panic("No todas las dimensiones tienen la misma profundidad")
			}
		case Expresion:
			if i == 0 {
				dimension = 1
			} else if dimension != 1 {
				panic("No todas las dimensiones tienen la misma profundidad")
			}
			resultado := rr.Ejecutar(ambito_padre)
			switch r3 := resultado.(type) {
			case int:
				if primitivo == "" || primitivo == "Int" {
					vector[i] = r3
				} else if primitivo == "Float" {
					vector[i] = float64(r3)
				} else {
					panic("Error el valor no coincide con el tipo de la matriz")
				}
			case float64:
				if primitivo == "" || primitivo == "Float" {
					vector[i] = r3
				} else {
					panic("Error el valor no coincide con el tipo de la matriz")
				}
			case string:
				if primitivo == "" || primitivo == "String" {
					vector[i] = r3
				} else {
					panic("Error el valor no coincide con el tipo de la matriz")
				}
			case rune:
				if primitivo == "" || primitivo == "Character" {
					vector[i] = r3
				} else {
					panic("Error el valor no coincide con el tipo de la matriz")
				}
			case bool:
				if primitivo == "" || primitivo == "Bool" {
					vector[i] = r3
				} else {
					panic("Error el valor no coincide con el tipo de la matriz")
				}
			default:
				panic("Tipo no permitido de la matriz")
			}
		default:
			panic("Error desconocido por el tipo de la matriz")
		}
	}
	return dimension
}
