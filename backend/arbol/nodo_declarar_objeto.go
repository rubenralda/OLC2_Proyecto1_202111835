package arbol

import "main/ambito"

type Declarar_objeto struct {
	Id    string
	Dupla []Dupla_atributos
}

func (d Declarar_objeto) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(d.Id)
	if encontrado == nil {
		panic("El struct no existe")
	}
	if encontrado.Tipo != "struct" {
		panic("El id no pertenece a un struct")
	}
	indice_dupla := 0
	largo := len(d.Dupla)
	ambito_objeto_struct := &ambito.Ambito{NombreAmbito: encontrado.Id}
	objeto_resultado := ambito.Objeto_struct{Ambito_struct: ambito_objeto_struct}
	for _, atributo_struct_p := range encontrado.Objeto.Ambito_struct.Locales {
		atributo_struct := *atributo_struct_p
		if indice_dupla >= largo { // como no hay mas atributos por inicializar comprueba si el resto esta inicizalido
			if atributo_struct.Valor == nil {
				panic("Atributo sin inicializar " + atributo_struct.Id)
			}
			ambito_objeto_struct.AgregarIde(atributo_struct)
			continue
		}
		if atributo_struct.Id != d.Dupla[indice_dupla].Id_atributo { // si no es el atributo salta al siguiente
			if atributo_struct.Valor == nil {
				panic("Atributo sin inicializar " + atributo_struct.Id)
			}
			ambito_objeto_struct.AgregarIde(atributo_struct)
			continue
		}
		if atributo_struct.Tipo == "constante" && atributo_struct.Valor != nil { // si es constante verifica si esta vacio
			panic("El atributo constante " + atributo_struct.Id + " ya esta inicilizado")
		}
		// en este punto ya esta verificado cual atributo se va modificar
		expresion := d.Dupla[indice_dupla].Expresion.Ejecutar(ambito_padre)
		switch rr := expresion.(type) {
		case int:
			if atributo_struct.Primitivo == "Int" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else if atributo_struct.Primitivo == "Float" {
				atributo_struct.Valor = float64(rr)
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el atributo no coincide con el tipo " + atributo_struct.Id)
			}
		case float64:
			if atributo_struct.Primitivo == "Float" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el atributo no coincide con el tipo " + atributo_struct.Id)
			}
		case string:
			if atributo_struct.Primitivo == "String" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el atributo no coincide con el tipo " + atributo_struct.Id)
			}
		case rune:
			if atributo_struct.Primitivo == "Character" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el atributo no coincide con el tipo " + atributo_struct.Id)
			}
		case bool:
			if atributo_struct.Primitivo == "Bool" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el atributo no coincide con el tipo " + atributo_struct.Id)
			}
		case ambito.Objeto_struct:
			if atributo_struct.Primitivo == "objeto" {
				atributo_struct.Valor = rr
				ambito_objeto_struct.AgregarIde(atributo_struct)
			} else {
				panic("Error el valor no coincide con el tipo " + atributo_struct.Id)
			}
		default:
			panic("Tipo no reconocido")
		}
		indice_dupla++
	}
	return objeto_resultado
}

type Dupla_atributos struct {
	Id_atributo string
	Expresion   BaseNodo
}
