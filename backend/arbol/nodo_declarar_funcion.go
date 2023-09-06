package arbol

import "main/ambito"

type Ejecutar_funcion struct {
	Id               string
	Tipo_retorno     string
	Sentencias       []BaseNodo
	Lista_parametros []Lista_parametros
}

func (e Ejecutar_funcion) Ejecutar(ambito_padre *ambito.Ambito, Lista_argumentos []Lista_argumentos) interface{} {
	ambito_local := &ambito.Ambito{NombreAmbito: e.Id, Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	if len(Lista_argumentos) != len(e.Lista_parametros) {
		panic("La cantidad de argumentos no coincide")
	}
	for indice, parametro := range e.Lista_parametros { // viene la misma cantidad de argumentos
		if parametro.Referencia != Lista_argumentos[indice].Referencia {
			panic("Discrepancia con los parametros por referencia o valor")
		}
		if parametro.Referencia && Lista_argumentos[indice].Referencia {
			if Lista_argumentos[indice].Expresion.(Expresion).Operacion == "identificador" {
				id_referencia := Lista_argumentos[indice].Expresion.(Expresion).Valor1.(Id_expresion).Id
				encontrado := ambito_padre.BuscarID(id_referencia)
				if encontrado == nil {
					panic("El id no existe " + id_referencia)
				}
				continue
			} else {
				panic("Error desconocido " + parametro.Id_interno)
			}
		}
		nombre_final := parametro.Id_interno
		if parametro.Id_externo == "_" {
			if Lista_argumentos[indice].Id_externo != "" {
				panic("No lleva parametro externo " + Lista_argumentos[indice].Id_externo)
			}
		} else if parametro.Id_externo == "" {
			if parametro.Id_interno != Lista_argumentos[indice].Id_externo {
				panic("El id del argumento no coincide con el parametro" + parametro.Id_interno)
			}
		}
		resul_valor := Lista_argumentos[indice].Expresion.Ejecutar(ambito_padre)
		switch rr := resul_valor.(type) {
		case int:
			if parametro.Primitivo == "Int" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: "Int"})
			} else if parametro.Primitivo == "Float" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: float64(rr), Tipo: "constante", Primitivo: "Float"})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case float64:
			if parametro.Primitivo == "Float" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: "Float"})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case string:
			if parametro.Primitivo == "String" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: "String"})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case rune:
			if parametro.Primitivo == "Character" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: "Character"})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case bool:
			if parametro.Primitivo == "Bool" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: "Bool"})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case ambito.Objeto_struct:
			if parametro.Primitivo == rr.Ambito_struct.NombreAmbito {
				ambito_local.AgregarIde(ambito.Identificadores{Id: nombre_final, Valor: rr, Tipo: "constante", Primitivo: rr.Ambito_struct.NombreAmbito})
			} else {
				panic("Error el valor no coincide con el tipo " + nombre_final)
			}
		case nil: //para una llamada o mejor un tipo de struct
		default:
			panic("Tipo no permitido " + nombre_final)
		}
	}
	for _, linea := range e.Sentencias {
		ejecutada := linea.Ejecutar(ambito_local)
		switch rr := ejecutada.(type) {
		case Control_transfer:
			if rr.Sentencia_return {
				retorno := rr.Retorno.Ejecutar(ambito_local)
				switch re := retorno.(type) {
				case int:
					if e.Tipo_retorno == "Int" {
						return re
					} else if e.Tipo_retorno == "Float" {
						return float64(re)
					}
				case float64:
					if e.Tipo_retorno == "Float" {
						return re
					}
				case string:
					if e.Tipo_retorno == "String" {
						return re
					}
				case rune:
					if e.Tipo_retorno == "Character" {
						return re
					}
				case bool:
					if e.Tipo_retorno == "Bool" {
						return re
					}
				case ambito.Objeto_struct:
					if e.Tipo_retorno == re.Ambito_struct.NombreAmbito {
						return re
					}
				case nil:
					if e.Tipo_retorno == "" {
						return nil
					}
				default:
					panic("Tipo no permitido " + e.Id)
				}
				panic("Error el valor no coincide con el tipo de la funcion " + e.Id)
			}
			panic("Error sentencias de control no van aqui")
		case nil: // si es nil no hace nada
		default:
			panic("No deberia pasar")
		}
	}
	return nil
}

type Declarar_funcion struct {
	Id               string
	Tipo_retorno     string
	Sentencias       []BaseNodo
	Lista_parametros []Lista_parametros
}

func (d Declarar_funcion) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	funcion := Ejecutar_funcion(d)
	ambito_padre.AgregarIde(ambito.Identificadores{Id: d.Id, Primitivo: d.Tipo_retorno, Funcion: funcion, Tipo: "funcion"})
	return nil
}

type Lista_parametros struct {
	Id_interno string
	Id_externo string
	Referencia bool
	Primitivo  string // String, Int, Bool, Float, char, (Nombre_struct)
}
