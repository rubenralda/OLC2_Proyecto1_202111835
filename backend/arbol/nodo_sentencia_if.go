package arbol

import "main/ambito"

type Sentencia_if struct {
	Expresion       BaseNodo
	Sentencias      []BaseNodo
	Sentencias_else []BaseNodo
	Siguiente       BaseNodo
}

func (s Sentencia_if) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := s.Expresion.Ejecutar(ambito_padre)
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia if", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	switch rr := resultado.(type) {
	case bool:
		if rr {
			for _, linea := range s.Sentencias {
				ejecutada := linea.Ejecutar(ambito_local)
				if ejecutada != nil {
					return ejecutada //puede ser una sentenica de transferencia
				}
			}
		} else if s.Siguiente != nil {
			ejecutada := s.Siguiente.Ejecutar(ambito_local)
			if ejecutada != nil {
				return ejecutada //puede ser una sentenica de transferencia
			}
		} else {
			for _, linea := range s.Sentencias_else {
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
