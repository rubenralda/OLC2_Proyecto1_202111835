package arbol

import "main/ambito"

type Llamada_metodo struct {
	Id_metodo        string
	Id_objeto        string
	Lista_argumentos []Lista_argumentos
}

func (l Llamada_metodo) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(l.Id_objeto)
	if encontrado == nil {
		panic("El id no existe " + l.Id_objeto)
	}
	switch encontrado.Valor.(type) {
	case ambito.Objeto_struct:
		// crear un identificador del mismo tipo que encontrado pero con ide self, si la funcion es
		// mutable (referencia=true) de igual forma pongo que es mutable y siempre lo guardo en puntero_valor
		ambito_local := &ambito.Ambito{NombreAmbito: "Ambito Self", Padre: ambito_padre}
		ambito_padre.AgregarHijo(ambito_local)
		funcion := encontrado.Valor.(ambito.Objeto_struct).Ambito_struct.BuscarID(l.Id_metodo)
		if funcion.Tipo != "funcion" {
			panic("El id no pertenece a una funcion " + funcion.Id)
		}
		ambito_local.AgregarIde(ambito.Identificadores{Id: "self", Tipo: "variable", Puntero_valor: encontrado,
			Primitivo: encontrado.Primitivo, Referencia: funcion.Referencia, Valor: encontrado.Valor})
		return funcion.Funcion.(Ejecutar_funcion).Ejecutar(ambito_local, l.Lista_argumentos)
	default:
		panic("El id no es una instancia " + encontrado.Id)
	}
}
