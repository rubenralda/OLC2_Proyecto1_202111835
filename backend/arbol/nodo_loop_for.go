package arbol

import (
	"main/ambito"
	"strconv"
)

type Loop_for_in struct {
	Expresion  BaseNodo
	Id         string
	Inicio     string
	Final      string
	Sentencias []BaseNodo
}

func (l Loop_for_in) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	ambito_local := &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	if l.Expresion == nil {
		if l.Inicio < l.Final {
			inicio, _ := strconv.Atoi(l.Inicio)
			final, _ := strconv.Atoi(l.Final)
			for i := inicio; i <= final; i++ {
				ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "int", Tipo: "constante", Valor: i})
				for _, linea := range l.Sentencias {
					linea.Ejecutar(ambito_local)
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
			ambito_local.AgregarIde(ambito.Identificadores{Id: l.Id, Primitivo: "character", Tipo: "constante", Valor: char})
			for _, linea := range l.Sentencias {
				linea.Ejecutar(ambito_local)
			}
			ambito_local = &ambito.Ambito{NombreAmbito: "Ciclo for", Padre: ambito_padre}
			ambito_padre.AgregarHijo(ambito_local)
		}
	default:
		panic("La expresion no es un string")
	}
	return nil
}
