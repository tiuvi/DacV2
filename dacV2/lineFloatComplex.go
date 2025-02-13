package dacv2

func (Space *Space) GetLineFloat64(col int64, line int64) (float64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToFloat64(RBuffer), nil
}

func (Space *Space) NewLineFloat64(col int64, data float64) (int64, error) {

	return Space.NewLineRaw(col, Float64ToBytes(data))
}

func (Space *Space) SetLineFloat64(col int64, line int64, data float64) error {

	return Space.SetLineRaw(col, line, Float64ToBytes(data))
}

func (Space *Space) GetLineFloat32(col int64, line int64) (float32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToFloat32(RBuffer), nil
}

func (Space *Space) NewLineFloat32(col int64, data float32) (int64, error) {

	return Space.NewLineRaw(col, Float32ToBytes(data))
}

func (Space *Space) SetLineFloat32(col int64, line int64, data float32) error {

	return Space.SetLineRaw(col, line, Float32ToBytes(data))
}

func (Space *Space) GetLineComplex128(col int64, line int64) (complex128, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToComplex128(RBuffer), nil
}

func (Space *Space) NewLineComplex128(col int64, data complex128) (int64, error) {

	return Space.NewLineRaw(col, Complex128ToBytes(data))
}

func (Space *Space) SetLineComplex128(col int64, line int64, data complex128) error {

	return Space.SetLineRaw(col, line, Complex128ToBytes(data))
}

func (Space *Space) GetLineComplex64(col int64, line int64) (complex64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToComplex64(RBuffer), nil
}

func (Space *Space) NewLineComplex64(col int64, data complex64) (int64, error) {

	return Space.NewLineRaw(col, Complex64ToBytes(data))
}

func (Space *Space) SetLineComplex64(col int64, line int64, data complex64) error {

	return Space.SetLineRaw(col, line, Complex64ToBytes(data))
}
