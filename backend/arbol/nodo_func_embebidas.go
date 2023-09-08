package arbol

import (
	"fmt"
	"main/ambito"
	"strconv"
)

var Salid_programa string

type Funcion_print struct {
	Lista_expresion []BaseNodo
}

func (i Funcion_print) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	salida := ""
	for _, expresion := range i.Lista_expresion {
		resultado := expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			salida += strconv.Itoa(rr) + " "
		case float64:
			salida += strconv.FormatFloat(rr, 'f', -1, 64) + " "
		case string:
			salida += rr + " "
		case rune:
			salida += string(rr) + " "
		case bool:
			salida += strconv.FormatBool(rr) + " "
		case nil:
			salida += "nil "
		default:
			salida += "######"
		}
	}
	fmt.Println(salida)
	salida += "\n"
	Salid_programa += salida
	return nil
}

type Funcion_int struct {
	Expresion BaseNodo
}

func (i Funcion_int) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := i.Expresion.Ejecutar(ambito_padre)
	switch rr := resultado.(type) {
	case float64:
		return int(rr)
	case string:
		numero, err := strconv.Atoi(rr)
		if err != nil {
			panic("No se puede convertir a valor numerico")
			//return nil
		}
		return numero
	default:
		panic("No se puede convertir a valor numerico")
		//return nil
	}
}

type Funcion_float struct {
	Expresion BaseNodo
}

func (i Funcion_float) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := i.Expresion.Ejecutar(ambito_padre)
	switch rr := resultado.(type) {
	case string:
		numero, err := strconv.ParseFloat(rr, 64)
		if err != nil {
			panic("No se puede convertir a valor numerico")
			//return nil
		}
		return numero
	case int:
		return float64(rr)
	default:
		panic("No se puede convertir a valor numerico")
		//return nil
	}
}

type Funcion_string struct {
	Expresion BaseNodo
}

func (i Funcion_string) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := i.Expresion.Ejecutar(ambito_padre)
	switch rr := resultado.(type) {
	case int:
		return strconv.Itoa(rr)
	case float64:
		return strconv.FormatFloat(rr, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(rr)
	default:
		panic("No se puede convertir a valor numerico")
		//return nil
	}
}
