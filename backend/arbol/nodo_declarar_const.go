package arbol

import "main/ambito"

// declarar constante

type Declarar_constante struct {
	Id        string
	Tipo      string   // String, Int, Bool, Float, char, (Nombre_struct)
	Expresion BaseNodo // 10, 20.5, "hola", true, objeto_strcut
}

func (d Declarar_constante) Ejecutar(ambito_local *ambito.Ambito) interface{} {
	if ambito_local.BuscarID(d.Id) != nil {
		panic("Error el id ya existe: " + d.Id)
	}
	resultado := d.Expresion.Ejecutar(ambito_local)
	switch rr := resultado.(type) {
	case int:
		if d.Tipo == "" || d.Tipo == "Int" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "Int"})
		} else if d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: float64(rr), Tipo: "constante", Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case float64:
		if d.Tipo == "" || d.Tipo == "Float" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "Float"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case string:
		if d.Tipo == "" || d.Tipo == "String" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "String"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case rune:
		if d.Tipo == "" || d.Tipo == "Character" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "Character"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case bool:
		if d.Tipo == "" || d.Tipo == "Bool" {
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: rr, Tipo: "constante", Primitivo: "Bool"})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case ambito.Objeto_struct:
		if d.Tipo == "" || d.Tipo == rr.Ambito_struct.NombreAmbito {
			valor := Copiar_objeto_struct(rr)
			ambito_local.AgregarIde(ambito.Identificadores{Id: d.Id, Valor: valor, Tipo: "constante", Primitivo: rr.Ambito_struct.NombreAmbito})
		} else {
			panic("Error el valor no coincide con el tipo " + d.Id)
		}
	case nil: //para una llamada o mejor un tipo de struct
	default:
		panic("Tipo no permitido " + d.Id)
	}
	return nil
}
