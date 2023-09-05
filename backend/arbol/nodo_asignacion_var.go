package arbol

import "main/ambito"

type Asignacion_variable struct {
	Id        string
	Expresion BaseNodo
}

func (a Asignacion_variable) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.Id)
	if encontrado != nil {
		if encontrado.Tipo == "constante" && encontrado.Tipo != "variable" {
			panic("No se puede modificar una constante")
		}
		resultado := a.Expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "Int" {
				encontrado.Valor = rr
				return nil
			}
			if encontrado.Primitivo == "Float" {
				encontrado.Valor = float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "Float" {
				encontrado.Valor = rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "String" {
				encontrado.Valor = rr
				return nil
			}
		case bool:
			if encontrado.Primitivo == "Bool" {
				encontrado.Valor = rr
				return nil
			}
		case rune:
			if encontrado.Primitivo == "Character" {
				encontrado.Valor = rr
				return nil
			}
		case ambito.Objeto_struct:
			if encontrado.Primitivo == rr.Ambito_struct.NombreAmbito {
				encontrado.Valor = rr
				return nil
			}
		default:
			panic("Tipo no reconocido")
		}
	}
	panic("Error el ID no existe " + a.Id)
}
