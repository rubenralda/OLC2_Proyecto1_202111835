package arbol

import (
	"main/ambito"
)

type Loop_for_in struct {
	Expresion  BaseNodo
	Id         string
	Inicio     BaseNodo
	Final      BaseNodo
	Sentencias []BaseNodo
}

func (l Loop_for_in) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	ambito_local := &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	if l.Expresion == nil {
		inicio := l.Inicio.Ejecutar(ambito_padre)
		switch inicio.(type) {
		case int: //no hace nada mas que confirmar el int
		default:
			panic("El rango no es entero")
		}
		final := l.Final.Ejecutar(ambito_padre)
		switch final.(type) {
		case int: //no hace nada mas que confirmar el int
		default:
			panic("El rango no es entero")
		}
		if inicio.(int) <= final.(int) {
			for i := inicio.(int); i <= final.(int); i++ {
				if l.Id != "_" {
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Int", Tipo: "constante", Valor: i})
				}
				for _, linea := range l.Sentencias {
					ejecutada := linea.Ejecutar(ambito_local)
					switch rr := ejecutada.(type) {
					case Control_transfer:
						if rr.Sentencia_break {
							return nil // se sale de todo porque es break
						} else if rr.Sentencia_continue {
							break // termina este for de sentencias para que valide nuevamente y ejecute
						} else { //es una sentencia return
							return rr // retornar el objeto return
						}
					case Sentencia_return:
						return rr //asciendo el retorno
					case nil: // si es nil no hace nada
					default:
						panic("No deberia pasar")
					}
				}
				ambito_local = &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
				ambito_padre.AgregarHijo(ambito_local)
			}
			return nil
		} else {
			panic("Rango de inicio es mayor al rango de final")
		}
	}
	resultado := l.Expresion.Ejecutar(ambito_padre)
	switch rr := resultado.(type) {
	case string:
		for _, char := range rr {
			if l.Id != "_" {
				ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Character", Tipo: "constante", Valor: char})
			}
		CicloIntruccion:
			for _, linea := range l.Sentencias {
				ejecutada := linea.Ejecutar(ambito_local)
				switch rr := ejecutada.(type) {
				case Control_transfer:
					if rr.Sentencia_break {
						return nil // se sale de todo porque es break
					} else if rr.Sentencia_continue {
						break CicloIntruccion // termina este for de sentencias para que valide nuevamente y ejecute
					} else { //es una sentencia return
						return rr // retornar el objeto return
					}
				case Sentencia_return:
					return rr //asciendo el retorno
				case nil: // si es nil no hace nada
				default:
					panic("No deberia pasar")
				}
			}
			ambito_local = &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
			ambito_padre.AgregarHijo(ambito_local)
		}
	case []interface{}:
		for _, item := range rr {
			if l.Id != "_" {
				switch ii := item.(type) {
				case int:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Int", Tipo: "constante", Valor: ii})
				case float64:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Float", Tipo: "constante", Valor: ii})
				case string:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "String", Tipo: "constante", Valor: ii})
				case rune:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Character", Tipo: "constante", Valor: ii})
				case bool:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "Bool", Tipo: "constante", Valor: ii})
				case ambito.Objeto_struct:
					ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: ii.Ambito_struct.NombreAmbito, Tipo: "constante", Valor: ii})
				}
			}
		CicloVector:
			for _, linea := range l.Sentencias {
				ejecutada := linea.Ejecutar(ambito_local)
				switch rr := ejecutada.(type) {
				case Control_transfer:
					if rr.Sentencia_break {
						return nil // se sale de todo porque es break
					} else if rr.Sentencia_continue {
						break CicloVector // termina este for de sentencias para que valide nuevamente y ejecute
					} else { //es una sentencia return
						return rr // retornar el objeto return
					}
				case Sentencia_return:
					return rr //asciendo el retorno
				case nil: // si es nil no hace nada
				default:
					panic("No deberia pasar")
				}
			}
			ambito_local = &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
			ambito_padre.AgregarHijo(ambito_local)
		}
	default:
		panic("La expresion no es un string")
	}
	return nil
}
