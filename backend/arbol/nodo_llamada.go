package arbol

import "main/ambito"

type Llamada_funcion struct {
	Id                  string
	Lista_argumentos    []Lista_argumentos
	Declarar_objeto_amb Declarar_objeto
}

func (l Llamada_funcion) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(l.Id)
	if encontrado == nil {
		panic("El id no existe " + l.Id)
	}
	if encontrado.Tipo != "funcion" {
		return l.Declarar_objeto_amb.Ejecutar(ambito_padre)
		//panic("El id no pertenece a una funcion " + l.Id)
	}
	return encontrado.Funcion.(Ejecutar_funcion).Ejecutar(ambito_padre, l.Lista_argumentos)
}

type Lista_argumentos struct {
	Id_externo string
	Referencia bool // cuando se usa & verificar que sea id
	Expresion  BaseNodo
}
