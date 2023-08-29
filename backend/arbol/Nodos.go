package arbol

import (
	"fmt"
	"strconv"
)

type BaseNodo interface {
	Ejecutar() interface{}
}

type Primitivos struct {
	Tipo  string
	Valor string
}

func (p Primitivos) Ejecutar() interface{} {
	switch p.Tipo {
	case "int":
		num, _ := strconv.Atoi(p.Valor)
		return num
	case "bool":
		break
	case "string":
		break
	case "float":
		break
	case "caracter":
		return p.Valor
	}
	return p.Valor
}

type Expresion struct {
	Valor1    BaseNodo
	Valor2    BaseNodo
	Operacion string
}

func (e Expresion) Ejecutar() interface{} {
	switch e.Operacion {
	case "suma":
		break
	case "primitivo":
		return e.Valor1.Ejecutar()
	default:
		return nil
	}
	return nil
}

type Imprimir struct {
	Valor BaseNodo
}

func (i Imprimir) Ejecutar() interface{} {
	fmt.Println(i.Valor.Ejecutar())
	return i.Valor.Ejecutar()
}
