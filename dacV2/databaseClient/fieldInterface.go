package databaseClient

import (
	dacv2 "dacV2"
)

func (Space *SpaceDB) GetFieldDecoder(col int64, outRef interface{}) error {
	// Obtenemos el buffer de bytes del archivo
	buffer, err := Space.GetField(col)
	if err != nil {
		return err
	}

	return dacv2.BytesToTypesGolang(buffer, outRef)
}

func (Space *SpaceDB) SetFieldEncoder(col int64, value interface{}) error {

	bufferRef, err := dacv2.TypesGolangToBytes(value)
	if err != nil {
		return err
	}

	return Space.SetField(col, bufferRef)
}

// Usar un struct valido
func (Space *SpaceDB) SetFieldJson(col int64, value interface{}) error {

	bufferRef, err := dacv2.TypesGolangToJsonBytes(value)
	if err != nil {
		return err
	}

	return Space.SetField(col, bufferRef)
}

/*
Usar una referencia &value para outRef.

Tambien se puede recuperar con las funciones de string o de bytes
*/
func (Space *SpaceDB) GetFieldJson(col int64, outRef interface{}) error {

	// Obtenemos el buffer de bytes del archivo
	buffer, err := Space.GetField(col)
	if err != nil {
		return err
	}

	return dacv2.JsonBytesToTypesGolang(buffer, outRef)
}
