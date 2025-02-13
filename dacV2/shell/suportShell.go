package shell

import(
	"strings"
)

func SliceStringArgs(stringMultiple string)(sliceString []string){

		// Separa los dominios por espacio
		sliceString = strings.Split(stringMultiple, " ")

		// Elimina los espacios en blanco al principio y al final de cada dominio
		for indice, domain := range sliceString {
			sliceString[indice] = strings.TrimSpace(domain)
		}

		return sliceString
}