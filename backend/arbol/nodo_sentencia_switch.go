package arbol

import "main/ambito"

type Sentencia_switch struct {
	Expresion    BaseNodo
	Lista_case   []Sentencia_case
	Default_case BaseNodo
}

func (s Sentencia_switch) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resultado := s.Expresion.Ejecutar(ambito_padre)
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia switch", Padre: ambito_padre}
	ambito_padre.AgregarHijo(ambito_local)
	switch rr := resultado.(type) {
	case float64:
		for _, sentencia_case := range s.Lista_case {
			resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
			switch r1 := resultado_case.(type) {
			case float64:
				if rr == r1 {
					ejecutada := sentencia_case.Ejecutar(ambito_local)
					if ejecutada != nil {
						return ejecutada //puede ser una sentenica de transferencia
					}
					return nil
				}
			default:
				panic("La expresion del case no es un float")
			}
		}
	case string:
		for _, sentencia_case := range s.Lista_case {
			resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
			switch r1 := resultado_case.(type) {
			case string:
				if rr == r1 {
					ejecutada := sentencia_case.Ejecutar(ambito_local)
					if ejecutada != nil {
						return ejecutada //puede ser una sentenica de transferencia
					}
					return nil
				}
			default:
				panic("La expresion del case no es un string")
			}
		}
	case int:
		for _, sentencia_case := range s.Lista_case {
			resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
			switch r1 := resultado_case.(type) {
			case int:
				if rr == r1 {
					ejecutada := sentencia_case.Ejecutar(ambito_local)
					if ejecutada != nil {
						return ejecutada //puede ser una sentenica de transferencia
					}
					return nil
				}
			default:
				panic("La expresion del case no es un int")
			}
		}
	case bool:
		for _, sentencia_case := range s.Lista_case {
			resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
			switch r1 := resultado_case.(type) {
			case bool:
				if rr == r1 {
					ejecutada := sentencia_case.Ejecutar(ambito_local)
					if ejecutada != nil {
						return ejecutada //puede ser una sentenica de transferencia
					}
					return nil
				}
			default:
				panic("La expresion del case no es un bool")
			}
		}
	case rune:
		for _, sentencia_case := range s.Lista_case {
			resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
			switch r1 := resultado_case.(type) {
			case rune:
				if rr == r1 {
					ejecutada := sentencia_case.Ejecutar(ambito_local)
					if ejecutada != nil {
						return ejecutada //puede ser una sentenica de transferencia
					}
					return nil
				}
			default:
				panic("La expresion del case no es un bool")
			}
		}
	default:
		panic("Error la expresion no es un bool ")
	}
	s.Default_case.Ejecutar(ambito_local)
	return nil
}

type Sentencia_case struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (s Sentencia_case) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	for _, linea := range s.Sentencias {
		ejecutada := linea.Ejecutar(ambito_padre)
		if ejecutada != nil {
			return ejecutada //puede ser una sentenica de transferencia
		}
	}
	return nil
}

type Default_case struct {
	Sentencias []BaseNodo
}

func (d Default_case) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	for _, linea := range d.Sentencias {
		ejecutada := linea.Ejecutar(ambito_padre)
		if ejecutada != nil {
			return ejecutada //puede ser una sentenica de transferencia
		}
	}
	return nil
}
