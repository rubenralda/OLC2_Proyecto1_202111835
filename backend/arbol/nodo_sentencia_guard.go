package arbol

import "main/ambito"

type Sentencia_guard struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (s Sentencia_guard) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := s.Expresion.Ejecutar(ambito_padre)
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia guard", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	switch rr := resultado.(type) {
	case bool:
		if !rr {
			for _, linea := range s.Sentencias {
				ejecutada := linea.Ejecutar(ambito_local)
				if ejecutada != nil {
					return ejecutada //puede ser una sentenica de transferencia
				}
			}
		}
	default:
		panic("Error la expresion no es un bool ")
	}
	return nil
}
