package arbol

import "main/ambito"

type Decremento_atributo struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
	Expresion       BaseNodo
}

func (a Decremento_atributo) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.ID_inicial.ID)
	if encontrado == nil {
		panic("El ID " + a.ID_inicial.ID + " no existe")
	}
	atributo_valor := &encontrado.Valor
	if encontrado.Id == "self" {
		if !encontrado.Referencia {
			panic("La funcion es inmutable " + encontrado.Puntero_valor.Id)
		}
		atributo_valor = &encontrado.Puntero_valor.Valor
	}

	if a.ID_inicial.Vector { //si es vector
		if encontrado.Tipo != "vector" {
			panic("El Id no pertenece a un vector")
		}
		indice := 0
		resul_indice := a.ID_inicial.Indice.Ejecutar(ambito_padre)
		switch rr := resul_indice.(type) {
		case int:
			indice = rr
		default:
			panic("El indice no es un entero")
		}
		if indice < 0 || indice >= len(encontrado.Lista_vector) {
			panic("El indice no existe")
		} else if encontrado.Referencia {
			if indice < 0 || indice >= len(encontrado.Puntero_valor.Lista_vector) {
				panic("El indice no existe")
			}
		}
		if encontrado.Referencia {
			atributo_valor = &encontrado.Puntero_valor.Lista_vector[indice]
		} else {
			atributo_valor = &encontrado.Lista_vector[indice]
		}
	}
	switch (*atributo_valor).(type) {
	case ambito.Objeto_struct: // si es struct
		siguiente := atributo_valor                       //actual valor struct
		for indice, atributo := range a.Lista_atributos { // si o si viene uno
			actual := (*siguiente).(ambito.Objeto_struct).Ambito_struct.BuscarID(atributo.ID)
			if actual == nil {
				panic("El atributo no existe " + atributo.ID)
			}
			atributo_existe_valor := &actual.Valor
			if atributo.Vector { //si es vector
				if actual.Tipo != "vector" {
					panic("El Id no pertenece a un vector")
				}
				indice := 0
				resul_indice := atributo.Indice.Ejecutar(ambito_padre)
				switch rr := resul_indice.(type) {
				case int:
					indice = rr
				default:
					panic("El indice no es un entero")
				}
				if indice < 0 || indice >= len(actual.Lista_vector) {
					panic("El indice no existe")
				} else if actual.Referencia {
					if indice < 0 || indice >= len(actual.Puntero_valor.Lista_vector) {
						panic("El indice no existe")
					}
				}
				if actual.Referencia {
					atributo_existe_valor = &actual.Puntero_valor.Lista_vector[indice]
				} else {
					atributo_existe_valor = &actual.Lista_vector[indice]
				}
			}
			switch (*atributo_existe_valor).(type) {
			case ambito.Objeto_struct:
				if indice < len(a.Lista_atributos)-1 { //si le falta atributos para buscar
					siguiente = atributo_existe_valor // me muevo al siguiente objeto
					continue
				}
				resultado := a.Expresion.Ejecutar(ambito_padre)
				switch rr := resultado.(type) {
				case ambito.Objeto_struct:
					if actual.Primitivo == rr.Ambito_struct.NombreAmbito {
						actual.Valor = rr
						return nil
					}
				default:
					panic("Tipo no permitido " + actual.Id)
				}
				panic("Los tipos no coinciden ")
			default: //debe ser hoja entonces
				if indice < len(a.Lista_atributos)-1 { //si no es el ultimo error
					panic("La variable " + (*siguiente).(ambito.Identificadores).Id + " no tiene atributos")
				}
				resultado := a.Expresion.Ejecutar(ambito_padre)
				switch rr := resultado.(type) {
				case int:
					if actual.Primitivo == "Int" {
						actual.Valor = actual.Valor.(int) - rr
						return nil
					}
					if actual.Primitivo == "Float" {
						actual.Valor = actual.Valor.(float64) - float64(rr)
						return nil
					}
				case float64:
					if actual.Primitivo == "Float" {
						actual.Valor = actual.Valor.(float64) - rr
						return nil
					}
				default:
					panic("Tipo no permitido " + actual.Id)
				}
				panic("Los tipos no coinciden " + actual.Id)
			}
		}
	default:
		panic("Error no tiene atributos " + encontrado.Id)
	}
	return nil
}
