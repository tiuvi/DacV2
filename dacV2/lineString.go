package dacv2

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
// Crea una nueva linea en un archivo desde un string en la columna especificada.
func (Space *Space) NewLineString(col int64, str string) (int64, error) {

	return Space.NewLine(col, []byte(str))
}

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
// Actualiza una nueva linea en un archivo desde un string en la columna especificada.
func (Space *Space) SetLineString(col int64, line int64, str string) error {

	return Space.SetLine(col, line, []byte(str))
}

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
// Obtiene una linea en un archivo y lo formatea directamente a string
func (Space *Space) GetLineString(col int64, line int64) (string, error) {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return "", err
	}

	return string(RBuffer), nil
}

func (Space *Space) NewLineRunes(col int64, data []rune) (int64, error) {

	return Space.NewLine(col, RunesToBytes(data))
}

func (Space *Space) SetLineRunes(col int64, line int64, data []rune) error {

	return Space.SetLine(col, line, RunesToBytes(data))
}

func (Space *Space) GetLineRunes(col int64, line int64) ([]rune, error) {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return nil, err
	}

	return BytesToRune(RBuffer), nil
}
