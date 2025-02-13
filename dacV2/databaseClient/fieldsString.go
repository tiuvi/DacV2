package databaseClient

import (
	dacv2 "dacV2"
)

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
func (Space *SpaceDB) GetFieldString(col int64) (string, error) {

	buffer, err := Space.GetField(col)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

// Tamaño maximo 4 bytes por caracter y minimo 1 byte
func (Space *SpaceDB) SetFieldString(col int64, str string) error {

	return Space.SetField(col, []byte(str))
}

// tamaño: 4 bytes
func (Space *SpaceDB) GetFieldRunes(col int64) ([]rune, error) {

	buffer, err := Space.GetField(col)
	if err != nil {
		return nil, err
	}

	return dacv2.BytesToRune(buffer), nil
}

// tamaño: 4 bytes
func (Space *SpaceDB) SetFieldRunes(col int64, runes []rune) error {

	err := Space.SetField(col, dacv2.RunesToBytes(runes))
	if err != nil {
		return err
	}

	return nil
}
