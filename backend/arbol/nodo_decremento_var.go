package arbol

import "main/ambito"

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
			if encontrado.Primitivo == "Int" {
				encontrado.Valor = encontrado.Valor.(int) - rr
				return nil
			}
			if encontrado.Primitivo == "Float" {
				encontrado.Valor = encontrado.Valor.(float64) - float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "Float" {
				encontrado.Valor = encontrado.Valor.(float64) - rr
				return nil
			}
		default:
			panic("Operacion invalida -= " + d.Id)
		}
	}
	panic("Error no se puede asignar, ID no existe " + d.Id)
}
