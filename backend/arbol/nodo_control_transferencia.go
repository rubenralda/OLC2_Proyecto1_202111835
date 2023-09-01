package arbol

import "main/ambito"

type Control_transfer struct {
	Sentencia_break    bool
	Sentencia_continue bool
	Sentencia_return   bool
	Retorno            BaseNodo
}

func (l Control_transfer) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	return l
}

type Sentencia_return struct {
	Expresion    BaseNodo
	Tipo_retorno bool //true si tiene tipo, false si no
}

func (s Sentencia_return) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	if !s.Tipo_retorno {
		return nil
	}
	return s.Expresion.Ejecutar(ambito_padre)
}
