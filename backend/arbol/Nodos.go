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
	case "identificador":
		return e.Valor1.Ejecutar(ambito)
	default:
		return nil
	}
}
