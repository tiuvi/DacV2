package databaseClient

import (
	dacv2 "dacV2"
)

func (Space *SpaceDB) GetLineDecoder(col int64, line int64, outRef interface{}) error {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return err
	}

	return dacv2.BytesToTypesGolang(RBuffer, outRef)
}

func (Space *SpaceDB) NewLineEncoder(col int64, data interface{}) (int64, error) {

	bufferRef, err := dacv2.TypesGolangToBytes(data)
	if err != nil {
		return 0, err
	}

	return Space.NewLine(col, bufferRef)
}

// Los tipos encoder hay que sumarles siempre un bytes al total.
// Ejemplo un map cuesta 100 bytes, al final siempre se añade un byte 256, el total es 101
// Esto se hace porque se rellena con nulos el campo y luego se eliminan junto al byte 256
func (Space *SpaceDB) SetLineEncoder(col int64, line int64, data interface{}) error {

	bufferRef, err := dacv2.TypesGolangToBytes(data)
	if err != nil {
		return err
	}

	return Space.SetLine(col, line, bufferRef)
}

func (Space *SpaceDB) GetLineJson(col int64, line int64, outRef interface{}) error {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return err
	}

	return dacv2.JsonBytesToTypesGolang(RBuffer, outRef)
}

func (Space *SpaceDB) NewLineJson(col int64, data interface{}) (int64, error) {

	bufferRef, err := dacv2.TypesGolangToJsonBytes(data)
	if err != nil {
		return 0, err
	}

	return Space.NewLine(col, bufferRef)
}

func (Space *SpaceDB) SetLineJson(col int64, line int64, data interface{}) error {

	bufferRef, err := dacv2.TypesGolangToJsonBytes(data)
	if err != nil {
		return err
	}

	return Space.SetLine(col, line, bufferRef)
}
