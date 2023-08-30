package arbol

import (
	"fmt"
	"main/ambito"
	"strconv"
	"strings"
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

type Expresion struct {
	Valor1    BaseNodo
	Valor2    BaseNodo
	Operacion string
}

func (e Expresion) Ejecutar(ambito *ambito.Ambito) interface{} {
	switch e.Operacion {
	case "+":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) + resultado2.(int)
			case float64:
				return float64(resultado1.(int)) + resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		case float64:
			switch resultado2.(type) {
			case int:
				return resultado1.(float64) + float64(resultado2.(int))
			case float64:
				return resultado1.(float64) + resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) + resultado2.(string)
			default:
				panic("Solo se puede sumar string con string")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "negacion":
		resultado1 := e.Valor1.Ejecutar(ambito)
		switch rr := resultado1.(type) {
		case int:
			return -rr
		case float64:
			return -rr
		default:
			panic("No se puede negar la operacion")
		}
	case "-":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) - resultado2.(int)
			case float64:
				return float64(resultado1.(int)) - resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		case float64:
			switch resultado2.(type) {
			case int:
				return resultado1.(float64) - float64(resultado2.(int))
			case float64:
				return resultado1.(float64) - resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "*":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) * resultado2.(int)
			case float64:
				return float64(resultado1.(int)) * resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		case float64:
			switch resultado2.(type) {
			case int:
				return resultado1.(float64) * float64(resultado2.(int))
			case float64:
				return resultado1.(float64) * resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "/":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) / resultado2.(int)
			case float64:
				return float64(resultado1.(int)) / resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		case float64:
			switch resultado2.(type) {
			case int:
				return resultado1.(float64) / float64(resultado2.(int))
			case float64:
				return resultado1.(float64) / resultado2.(float64)
			default:
				panic("No se puede sumar int con ")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "%":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) % resultado2.(int)
			default:
				panic("No se puede sumar int con ")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case ">":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) > resultado2.(int)
			default:
				panic("Solo int > int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) > resultado2.(float64)
			default:
				panic("Solo float > float")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) > resultado2.(string)
			default:
				panic("Solo string > string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) > resultado2.(rune)
			default:
				panic("Solo char > char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "<":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) < resultado2.(int)
			default:
				panic("Solo int < int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) < resultado2.(float64)
			default:
				panic("Solo float < float")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) < resultado2.(string)
			default:
				panic("Solo string < string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) < resultado2.(rune)
			default:
				panic("Solo char < char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case ">=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) >= resultado2.(int)
			default:
				panic("Solo int >= int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) >= resultado2.(float64)
			default:
				panic("Solo float >= float")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) >= resultado2.(string)
			default:
				panic("Solo string >= string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) >= resultado2.(rune)
			default:
				panic("Solo char >= char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "<=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) <= resultado2.(int)
			default:
				panic("Solo int <= int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) <= resultado2.(float64)
			default:
				panic("Solo float <= float")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) <= resultado2.(string)
			default:
				panic("Solo string <= string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) <= resultado2.(rune)
			default:
				panic("Solo char <= char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "==":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) == resultado2.(int)
			default:
				panic("Solo int == int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) == resultado2.(float64)
			default:
				panic("Solo float == float")
			}
		case bool:
			switch resultado2.(type) {
			case bool:
				return resultado1.(bool) == resultado2.(bool)
			default:
				panic("Solo string == string")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) == resultado2.(string)
			default:
				panic("Solo string == string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) == resultado2.(rune)
			default:
				panic("Solo char == char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "!=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case int:
			switch resultado2.(type) {
			case int:
				return resultado1.(int) != resultado2.(int)
			default:
				panic("Solo int != int")
			}
		case float64:
			switch resultado2.(type) {
			case float64:
				return resultado1.(float64) != resultado2.(float64)
			default:
				panic("Solo float != float")
			}
		case bool:
			switch resultado2.(type) {
			case bool:
				return resultado1.(bool) != resultado2.(bool)
			default:
				panic("Solo bool != bool")
			}
		case string:
			switch resultado2.(type) {
			case string:
				return resultado1.(string) != resultado2.(string)
			default:
				panic("Solo string != string")
			}
		case rune:
			switch resultado2.(type) {
			case rune:
				return resultado1.(rune) != resultado2.(rune)
			default:
				panic("Solo char != char")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "!":
		resultado1 := e.Valor1.Ejecutar(ambito)
		switch rr := resultado1.(type) {
		case bool:
			return !rr
		default:
			panic("La operacion ! no es un bool")
		}
	case "&&":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case bool:
			switch resultado2.(type) {
			case bool:
				return resultado1.(bool) && resultado2.(bool)
			default:
				panic("Solo string == string")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "||":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		switch resultado1.(type) {
		case bool:
			switch resultado2.(type) {
			case bool:
				return resultado1.(bool) || resultado2.(bool)
			default:
				panic("Solo string == string")
			}
		default:
			panic("La operacion suma no esta..")
		}
	case "primitivo":
		return e.Valor1.Ejecutar(ambito)
	default:
		return nil
	}
}

type Imprimir struct {
	Valor BaseNodo
}

func (i Imprimir) Ejecutar(ambito *ambito.Ambito) interface{} {
	fmt.Println("funcion imprimir:", i.Valor.Ejecutar(ambito))
	//fmt.Println(ambito)
	ambito.NombreAmbito = "cambiado"
	return i.Valor.Ejecutar(ambito)
}

// nodo declarar variable

type Declarar_variable struct {
	Id        string
	Tipo      string
	Expresion BaseNodo
}

func (d Declarar_variable) Ejecutar(ambito_local *ambito.Ambito) interface{} {
	if ambito_local.BuscarID(d.Id) != nil {
		panic("Error la variable ya existe: " + d.Id)
	}
	if d.Expresion == nil {
		if d.Tipo != "" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: nil, Tipo: "variable", Primitivo: strings.ToLower(d.Tipo)})
			return nil
		} else {
			panic("Error no tiene tipo")
		}
	}
	resultado := d.Expresion.Ejecutar(ambito_local)
	switch rr := resultado.(type) {
	case int:
		if d.Tipo == "" || d.Tipo == "Int" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "int"})
		} else if d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: float64(rr), Tipo: "variable", Primitivo: "float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case float64:
		if d.Tipo == "" || d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case string:
		if d.Tipo == "" || d.Tipo == "String" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "string"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case rune:
		if d.Tipo == "" || d.Tipo == "Character" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "character"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case bool:
		if d.Tipo == "" || d.Tipo == "Bool" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "bool"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case nil: //cambiar

	}
	return nil
}

type Asignacion_variable struct {
	Id        string
	Expresion BaseNodo
}

func (a Asignacion_variable) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(a.Id)
	if encontrado != nil {
		if encontrado.Tipo == "constante" {
			panic("No se puede modificar una constante")
		}
		resultado := a.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Valor = rr
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Valor = float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Valor = rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "string" {
				encontrado.Valor = rr
				return nil
			}
		case bool:
			if encontrado.Primitivo == "bool" {
				encontrado.Valor = rr
				return nil
			}
		case rune:
			if encontrado.Primitivo == "character" {
				encontrado.Valor = rr
				return nil
			}
		}
	}
	panic("Error no se puede asignar " + a.Id)
}

type Incremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (i Incremento_variable) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(i.Id)
	if encontrado != nil {
		resultado := i.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Valor = encontrado.Valor.(int) + rr
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Valor = encontrado.Valor.(float64) + float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Valor = encontrado.Valor.(float64) + rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "string" {
				encontrado.Valor = encontrado.Valor.(string) + rr
				return nil
			}
		default:
			panic("Operacion invalida += " + i.Id)
		}
	}
	panic("Error no se puede asignar " + i.Id)
}

type Decremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (d Decremento_variable) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(d.Id)
	if encontrado != nil {
		resultado := d.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "int" {
				encontrado.Valor = encontrado.Valor.(int) - rr
				return nil
			}
			if encontrado.Primitivo == "float" {
				encontrado.Valor = encontrado.Valor.(float64) - float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "float" {
				encontrado.Valor = encontrado.Valor.(float64) - rr
				return nil
			}
		default:
			panic("Operacion invalida -= " + d.Id)
		}
	}
	panic("Error no se puede asignar " + d.Id)
}

// declarar constante

type Declarar_constante struct {
	Id        string
	Tipo      string
	Expresion BaseNodo
}

func (d Declarar_constante) Ejecutar(ambito_local *ambito.Ambito) interface{} {
	if ambito_local.BuscarID(d.Id) != nil {
		panic("Error el id ya existe: " + d.Id)
	}
	resultado := d.Expresion.Ejecutar(ambito_local)
	switch rr := resultado.(type) {
	case int:
		if d.Tipo == "" || d.Tipo == "Int" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "int"})
		} else if d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: float64(rr), Tipo: "constante", Primitivo: "float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case float64:
		if d.Tipo == "" || d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case string:
		if d.Tipo == "" || d.Tipo == "String" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "string"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case rune:
		if d.Tipo == "" || d.Tipo == "Character" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "character"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case bool:
		if d.Tipo == "" || d.Tipo == "Bool" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "bool"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case nil: //cambiar

	}
	return nil
}
