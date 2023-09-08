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
		if d.Tipo_matriz_ex.Dimension != comprobar_dimensiones(d.Matriz) {
			panic("La definicion de dimensiones no concuerda " + d.Id)
		}
	} else {
		d.Tipo_matriz_ex.Dimension = comprobar_dimensiones(d.Matriz)
	}
	ambito_padre.AgregarIde(ambito.Identificadores{Id: d.Id, Primitivo: d.Tipo_matriz_ex.Primitivo, Tipo: "matriz", Lista_vector: d.Matriz})
	return nil
}

/*
comprueba si el slice cada posicion tiene la misma profundidad, por ejemplo si en
una posicion hay otro slice pero en la posicion de la par no, no tiene la misma profundidad y
tira error
*/
func comprobar_dimensiones(vector []interface{}) int {
	dimension := 0
	dim_hijo := 0
	for i, item := range vector {
		switch rr := item.(type) {
		case []interface{}:
			dim_hijo = comprobar_dimensiones(rr) + 1
			if i == 0 {
				dimension = dim_hijo
			} else if dimension != dim_hijo {
				panic("No todas las dimensiones tienen la misma profundidad")
			}
		case Expresion:
			//fmt.Println(rr.Ejecutar(ambito_padre))
			if i == 0 {
				dimension = 1
			} else if dimension != 1 {
				panic("No todas las dimensiones tienen la misma profundidad")
			}
		default:
			panic("Error desconocido por el tipo de la matriz")
		}
	}
	return dimension
}
