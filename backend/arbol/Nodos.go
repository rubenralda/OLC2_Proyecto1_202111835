package arbol

import (
	"main/ambito"
	"strconv"
)

type BaseNodo interface {
	Ejecutar(ambito *ambito.Ambito) interface{}
}

type Primitivos struct {
	Tipo  string
	Valor string
}

func (p Primitivos) Ejecutar(ambito *ambito.Ambito) interface{} {
	switch p.Tipo {
	case "Int":
		num, _ := strconv.Atoi(p.Valor)
		return num
	case "Bool":
		if p.Valor == "true" {
			return true
		} else if p.Valor == "false" {
			return false
		} else {
			return nil
		}
	case "String":
		return p.Valor[1 : len(p.Valor)-1]
	case "Float":
		num, _ := strconv.ParseFloat(p.Valor, 64)
		return num
	case "Caracter":
		return []rune(p.Valor)[1] //convertir a un array de bytes para operar si es necesario
	case "nulo":
		return nil
	default:
		panic("Tipo desconocido")
	}
}

type Id_expresion struct {
	Id string
}

func (i Id_expresion) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	id_buscado := ambito_padre.BuscarID(i.Id)
	if id_buscado == nil {
		panic("Error el identificador no existe: " + i.Id)
	}
	if id_buscado.Tipo == "vector" {
		return id_buscado.Lista_vector
	}
	if id_buscado.Tipo == "struct" {
		return id_buscado.Objeto
	}
	if id_buscado.Referencia {
		return id_buscado.Puntero_valor.Valor
	}
	return id_buscado.Valor
}

type Id_vector struct {
	Id     string
	Indice BaseNodo
}

func (i Id_vector) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resul_indice := i.Indice.Ejecutar(ambito_padre)
	indice := 0
	switch rr := resul_indice.(type) {
	case int:
		indice = rr
	default:
		panic("El indice no es un entero")
	}
	id_buscado := ambito_padre.BuscarID(i.Id)
	if id_buscado == nil {
		panic("Error el vector no existe: " + i.Id)
	}
	if id_buscado.Tipo != "vector" { //agregar funciones y structs
		panic("El Id no corresponde a una variable o constante " + i.Id)
	}
	if indice >= 0 && indice < len(id_buscado.Lista_vector) {
		return id_buscado.Lista_vector[indice]
	} else if id_buscado.Referencia {
		if indice >= 0 && indice < len(id_buscado.Puntero_valor.Lista_vector) {
			return id_buscado.Puntero_valor.Lista_vector[indice]
		}
		//return nil // esto para errores
	}
	panic("El indice no existe")
}
