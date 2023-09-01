package arbol

import (
	"fmt"
	"main/ambito"
)

type Imprimir struct {
	Valor BaseNodo
}

func (i Imprimir) Ejecutar(ambito *ambito.Ambito) interface{} {
	resultado := i.Valor.Ejecutar(ambito)
	switch rr := resultado.(type) {
	case rune:
		fmt.Println(string(rr))
		return nil
	}
	fmt.Println(resultado)
	return nil
}
