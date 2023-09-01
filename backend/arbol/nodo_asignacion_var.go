package arbol

import "main/ambito"

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
		default:
			panic("Tipo no reconocido")
		}
	}
	panic("Error el ID no existe " + a.Id)
}
