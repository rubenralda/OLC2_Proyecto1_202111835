package arbol

import (
	"main/ambito"
)

type Declarar_struct struct {
	Id        string
	Atributos []BaseNodo
}

func (d Declarar_struct) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	ambito_struct := &ambito.Ambito{NombreAmbito: d.Id, Padre: nil}
	objeto := ambito.Objeto_struct{Ambito_struct: ambito_struct}
	ambito_padre.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: nil, Tipo: "struct", Objeto: objeto})
	for _, linea := range d.Atributos {
		ejecutada := linea.Ejecutar(ambito_struct)
		switch ejecutada.(type) {
		case nil: // si es nil no hace nada
		default:
			panic("No deberia pasar")
		}
	}
	return nil
}

type Declaracion_atributo struct {
	Id        string
	Primitivo string // String, Int
	Expresion BaseNodo
	Tipo      string // let o var
}

func (d Declaracion_atributo) Ejecutar(ambito_local *ambito.Ambito) interface{} {
	if d.Expresion == nil {
		if d.Primitivo != "" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: nil, Tipo: d.Tipo, Primitivo: d.Primitivo})
			return nil
		} else {
			panic("Error no tiene tipo")
		}
	}
	resultado := d.Expresion.Ejecutar(ambito_local)
	switch rr := resultado.(type) {
	case int:
		if d.Primitivo == "" || d.Primitivo == "Int" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: d.Tipo, Primitivo: "Int"})
		} else if d.Primitivo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: float64(rr), Tipo: d.Tipo, Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case float64:
		if d.Primitivo == "" || d.Primitivo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: d.Tipo, Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case string:
		if d.Primitivo == "" || d.Primitivo == "String" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: d.Tipo, Primitivo: "String"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case rune:
		if d.Primitivo == "" || d.Primitivo == "Character" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: d.Tipo, Primitivo: "Character"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case bool:
		if d.Primitivo == "" || d.Primitivo == "Bool" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: d.Tipo, Primitivo: "Bool"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case nil: //cambiar por el valor de un struc llamada o algo asi
	default:
		panic("Tipo no reconocido")
	}
	return nil
}
