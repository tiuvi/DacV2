package dacv2



// Tamaño: 8 bytes
func (Space *Space) GetFieldUint64(col int64) (uint64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToUint64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *Space) SetFieldUint64(col int64, value uint64) error {

	return Space.SetFieldRaw(col, Uint64ToBytes(value))

}

func (Space *Space) GetFieldUint32(col int64) (uint32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToUint32(buffer), nil

}

// Tamaño: 4 bytes
func (Space *Space) SetFieldUint32(col int64, value uint32) {

	Space.SetFieldRaw(col, Uint32ToBytes(value))
}

func (Space *Space) GetFieldUint(col int64) (uint, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return uint(BytesToUint32(buffer)), nil
}

// Tamaño: 4 bytes
func (Space *Space) SetFieldUint(col int64, value uint) error {

	return Space.SetFieldRaw(col, Uint32ToBytes(uint32(value)))

}

func (Space *Space) GetFieldUint16(col int64) (uint16, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToUint16(buffer), nil
}

// Tamaño: 2 bytes
func (Space *Space) SetFieldUint16(col int64, value uint16) error {

	return Space.SetFieldRaw(col, Uint16ToBytes(value))
}

func (Space *Space) GetFieldUint8(col int64) (uint8, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToUint8(buffer), nil
}

// Tamaño: 1 bytes
func (Space *Space) SetFieldUint8(col int64, value uint8) error {

	return Space.SetFieldRaw(col, Uint8ToBytes(value))
}

// Tamaño: 8 bytes 
func (Space *Space) GetFieldInt64(col int64) (int64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToInt64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *Space) SetFieldInt64(col int64, value int64) error {

	return Space.SetFieldRaw(col, Int64ToBytes(value))
}


func (Space *Space) GetFieldInt(col int64) (int, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToInt(buffer), nil
}

// Tamaño: 4 bytes
func (Space *Space) SetFieldInt(col int64, value int) error {

	return Space.SetFieldRaw(col, IntToBytes(value))
}

func (Space *Space) GetFieldInt32(col int64) (int32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToInt32(buffer), nil
}

// Tamaño: 4 bytes
func (Space *Space) SetFieldInt32(col int64, value int32) error {

	return Space.SetFieldRaw(col, Int32ToBytes(value))
}

func (Space *Space) GetFieldInt16(col int64) (int16, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToInt16(buffer), nil

}

// Tamaño: 2 bytes
func (Space *Space) SetFieldInt16(col int64, value int16) error {

	return Space.SetFieldRaw(col, Int16ToBytes(value))
}

func (Space *Space) GetFieldInt8(col int64) (int8, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToInt8(buffer), nil

}

// Tamaño: 1 bytes
func (Space *Space) SetFieldInt8(col int64, value int8) error {

	return Space.SetFieldRaw(col, Int8ToBytes(value))
}

