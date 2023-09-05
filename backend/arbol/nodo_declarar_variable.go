package arbol

import (
	"main/ambito"
)

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
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: nil, Tipo: "variable", Primitivo: d.Tipo})
			return nil
		} else {
			panic("Error no tiene tipo")
		}
	}
	resultado := d.Expresion.Ejecutar(ambito_local)
	switch rr := resultado.(type) {
	case int:
		if d.Tipo == "" || d.Tipo == "Int" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "Int"})
		} else if d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: float64(rr), Tipo: "variable", Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case float64:
		if d.Tipo == "" || d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case string:
		if d.Tipo == "" || d.Tipo == "String" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "String"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case rune:
		if d.Tipo == "" || d.Tipo == "Character" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "Character"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case bool:
		if d.Tipo == "" || d.Tipo == "Bool" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "variable", Primitivo: "Bool"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case ambito.Objeto_struct:
		if d.Tipo == "" || d.Tipo == rr.Ambito_struct.NombreAmbito {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "objeto", Primitivo: rr.Ambito_struct.NombreAmbito})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case nil: //cambiar por el valor de un struc llamada o algo asi
	default:
		panic("Tipo no reconocido")
	}
	return nil
}
