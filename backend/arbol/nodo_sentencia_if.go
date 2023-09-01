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
				linea.Ejecutar(ambito_local)
			}
		} else if s.Siguiente != nil {
			s.Siguiente.Ejecutar(ambito_local)
		} else {
			for _, linea := range s.Sentencias_else {
				linea.Ejecutar(ambito_local)
			}
		}
	default:
		panic("Error la expresion no es un bool ")
	}
	return nil
}
