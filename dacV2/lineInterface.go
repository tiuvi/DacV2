package dacv2


func (Space *Space) GetLineDecoder(col int64, line int64, outRef interface{}) error {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return err
	}

	return BytesToTypesGolang(RBuffer, outRef)
}

func (Space *Space) NewLineEncoder(col int64, data interface{}) (int64, error) {

	bufferRef, err := TypesGolangToBytes(data)
	if err != nil {
		return 0, err
	}

	size := Space.IndexSizeColumns[col]

	sizeTotal := size[1] - size[0]

	if len(bufferRef) > int(sizeTotal) {
		return 0, ErrInterfaceTooLarge
	}

	return Space.NewLine(col, bufferRef)
}

// Los tipos encoder hay que sumarles siempre un bytes al total.
// Ejemplo un map cuesta 100 bytes, al final siempre se añade un byte 256, el total es 101
// Esto se hace porque se rellena con nulos el campo y luego se eliminan junto al byte 256
func (Space *Space) SetLineEncoder(col int64, line int64, data interface{}) error {

	bufferRef, err := TypesGolangToBytes(data)
	if err != nil {
		return err
	}

	size := Space.IndexSizeColumns[col]

	sizeTotal := size[1] - size[0]

	if len(bufferRef) > int(sizeTotal) {
		return ErrInterfaceTooLarge
	}

	return Space.SetLine(col, line, bufferRef)
}

func (Space *Space) GetLineJson(col int64, line int64, outRef interface{}) error {

	RBuffer, err := Space.GetLine(col, line)
	if err != nil {
		return err
	}

	return JsonBytesToTypesGolang(RBuffer, outRef)
}

func (Space *Space) NewLineJson(col int64, data interface{}) (int64, error) {

	bufferRef, err := TypesGolangToJsonBytes(data)
	if err != nil {
		return 0, err
	}

	return Space.NewLine(col, bufferRef)
}

func (Space *Space) SetLineJson(col int64, line int64, data interface{}) error {

	bufferRef, err := TypesGolangToJsonBytes(data)
	if err != nil {
		return err
	}

	return Space.SetLine(col, line, bufferRef)
}
