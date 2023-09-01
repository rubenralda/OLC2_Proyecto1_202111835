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
	case "int":
		num, _ := strconv.Atoi(p.Valor)
		return num
	case "bool":
		if p.Valor == "true" {
			return true
		} else if p.Valor == "false" {
			return false
		} else {
			return nil
		}
	case "string":
		return p.Valor[1 : len(p.Valor)-1]
	case "float":
		num, _ := strconv.ParseFloat(p.Valor, 64)
		return num
	case "caracter":
		return []rune(p.Valor)[1] //convertir a un array de bytes para operar si es necesario
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
	return id_buscado.Valor
}
