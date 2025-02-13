package dacv2



// Tamaño maximo 4 bytes por caracter y minimo 1 byte
func (Space *Space) GetFieldString(col int64) (string, error) {

	buffer, err := Space.GetField(col)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
func (Space *Space) SetFieldString(col int64, str string) error {

	return Space.SetField(col, []byte(str))
}


// tamaño: 4 bytes
func (Space *Space) GetFieldRunes(col int64) ([]rune, error) {

	buffer, err := Space.GetField(col)
	if err != nil {
		return nil, err
	}

	return BytesToRune(buffer), nil
}

// tamaño: 4 bytes
func (Space *Space) SetFieldRunes(col int64, runes []rune) error {

	err := Space.SetField(col, RunesToBytes(runes))
	if err != nil {
		return err
	}

	return nil
}



