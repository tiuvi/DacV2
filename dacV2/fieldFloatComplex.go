package dacv2

func (Space *Space) GetFieldFloat64(col int64) (float64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToFloat64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *Space) SetFieldFloat64(col int64, value float64) error {

	return Space.SetFieldRaw(col, Float64ToBytes(value))
}

func (Space *Space) GetFieldFloat32(col int64) (float32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToFloat32(buffer), nil
}

// Tamaño: 4 bytes
func (Space *Space) SetFieldFloat32(col int64, value float32) error {

	return Space.SetFieldRaw(col, Float32ToBytes(value))
}

func (Space *Space) GetFieldComplex128(col int64) (complex128, error) {
	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return BytesToComplex128(buffer), nil
}

// tamaño: 16 bytes
func (Space *Space) SetFieldComplex128(col int64, value complex128) error {

	return Space.SetFieldRaw(col, Complex128ToBytes(value))
}

func (Space *Space) GetFieldComplex64(col int64) (complex64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return BytesToComplex64(buffer), nil
}

func (Space *Space) SetFieldComplex64(col int64, value complex64) error {

	return Space.SetFieldRaw(col, Complex64ToBytes(value))
}

