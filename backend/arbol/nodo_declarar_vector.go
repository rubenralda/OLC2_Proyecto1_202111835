package arbol

import (
	"main/ambito"
	"strings"
)

type Declarar_vector struct {
	Tipo            string
	ID              string
	Lista_expresion []BaseNodo
	ID_otro_vector  string
}

func (d Declarar_vector) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	d.Tipo = strings.ToLower(d.Tipo) // puede no funcionar, probar despues
	if ambito_padre.BuscarID(d.ID) != nil {
		panic("Error el ID ya existe: " + d.ID)
	}
	if d.ID_otro_vector != "" {
		vector_copia := ambito_padre.BuscarID(d.ID_otro_vector)
		if vector_copia == nil {
			panic("Error el ID no existe: " + d.ID_otro_vector)
		}
		// validar si son del mismo tipo los vectores
		if vector_copia.Primitivo != d.Tipo || vector_copia.Tipo != "vector" {
			panic("No se puede copiar al vector, el ID " + vector_copia.Id)
		}
		copia_lista := make([]interface{}, len(vector_copia.Lista_vector))
		copy(copia_lista, vector_copia.Lista_vector)
		ambito_padre.AgregarIde(ambito.Identificadores{Id: d.ID, Primitivo: d.Tipo, Tipo: "vector", Lista_vector: copia_lista})
		return nil
	}
	var lista_valores []interface{}
	for _, expresion := range d.Lista_expresion {
		resultado := expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			if d.Tipo == "int" {
				lista_valores = append(lista_valores, rr)
			} else if d.Tipo == "float" {
				lista_valores = append(lista_valores, float64(rr))
			} else {
				panic("Error el valor no coincide con el tipo " + d.ID)
			}
		case float64:
			if d.Tipo == "Float" {
				lista_valores = append(lista_valores, rr)
			} else {
				panic("Error el valor no coincide con el tipo " + d.ID)
			}
		case string:
			if d.Tipo == "String" {
				lista_valores = append(lista_valores, rr)
			} else {
				panic("Error el valor no coincide con el tipo " + d.ID)
			}
		case rune:
			if d.Tipo == "Character" {
				lista_valores = append(lista_valores, rr)
			} else {
				panic("Error el valor no coincide con el tipo " + d.ID)
			}
		case bool:
			if d.Tipo == "Bool" {
				lista_valores = append(lista_valores, rr)
			} else {
				panic("Error el valor no coincide con el tipo " + d.ID)
			}
		case nil: //cambiar por el valor de un struc llamada o algo asi
		default:
			panic("Tipo no reconocido")
		}
	}
	ambito_padre.AgregarIde(ambito.Identificadores{Id: d.ID, Primitivo: d.Tipo, Tipo: "vector", Lista_vector: lista_valores})
	return nil
}