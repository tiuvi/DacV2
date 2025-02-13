package dacv2

func (Space *Space) GetFieldDecoder(col int64, outRef interface{}) error {
	// Obtenemos el buffer de bytes del archivo
	buffer, err := Space.GetField(col)
	if err != nil {
		return err
	}

	return BytesToTypesGolang(buffer, outRef)
}


func (Space *Space) SetFieldEncoder(col int64, value interface{}) error {

	bufferRef, err := TypesGolangToBytes(value)
	if err != nil {
		return err
	}

	size := Space.IndexSizeFields[col]

	sizeTotal := size[1] - size[0]

	if len(bufferRef) > int(sizeTotal) {
		return ErrInterfaceTooLarge
	}

	return Space.SetField(col, bufferRef)
}


// Usar un struct valido
func (Space *Space) SetFieldJson(col int64, value interface{}) error {

	bufferRef, err := TypesGolangToJsonBytes(value)
	if err != nil {
		return err
	}

	return Space.SetField(col, bufferRef)
}

/*
Usar una referencia &value para outRef.

Tambien se puede recuperar con las funciones de string o de bytes
*/
func (Space *Space) GetFieldJson(col int64, outRef interface{}) error {

	// Obtenemos el buffer de bytes del archivo
	buffer, err := Space.GetField(col)
	if err != nil {
		return err
	}

	return JsonBytesToTypesGolang(buffer, outRef)
}
