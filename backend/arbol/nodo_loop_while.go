package arbol

import "main/ambito"

type Loop_while struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (l Loop_while) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	ambito_local := &ambito.Ambito{NombreAmbito: "Ciclo while", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	resultado := l.Expresion.Ejecutar(ambito_padre)
	switch rr := resultado.(type) {
	case bool:
		for rr {
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
						return rr.Retorno // retornar el objeto return
					}
				case Sentencia_return:
					return rr //asciendo el retorno
				case nil: // si es nil no hace nada
				default:
					panic("No deberia pasar")
				}
			}
			ambito_local = &ambito.Ambito{NombreAmbito: "Ciclo while", Padre: ambito_padre}
			ambito_padre.AgregarHijo(ambito_local)
			rr = l.Expresion.Ejecutar(ambito_padre).(bool) // Si falla es porque de algun modo se cambio la expresion pero en teoria no se puede
		}
	default:
		panic("La expresion no es un bool")
	}
	return nil
}
