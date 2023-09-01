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
			for _, linea := range l.Sentencias {
				linea.Ejecutar(ambito_local)
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
