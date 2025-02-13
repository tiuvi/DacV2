package databaseClient

import (
	dacv2 "dacV2"
)

func (Space *SpaceDB) GetLineFloat64(col int64, line int64) (float64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToFloat64(RBuffer), nil
}

func (Space *SpaceDB) NewLineFloat64(col int64, data float64) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Float64ToBytes(data))
}

func (Space *SpaceDB) SetLineFloat64(col int64, line int64, data float64) error {

	return Space.SetLineRaw(col, line, dacv2.Float64ToBytes(data))
}

func (Space *SpaceDB) GetLineFloat32(col int64, line int64) (float32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToFloat32(RBuffer), nil
}

func (Space *SpaceDB) NewLineFloat32(col int64, data float32) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Float32ToBytes(data))
}

func (Space *SpaceDB) SetLineFloat32(col int64, line int64, data float32) error {

	return Space.SetLineRaw(col, line, dacv2.Float32ToBytes(data))
}

func (Space *SpaceDB) GetLineComplex128(col int64, line int64) (complex128, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToComplex128(RBuffer), nil
}

func (Space *SpaceDB) NewLineComplex128(col int64, data complex128) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Complex128ToBytes(data))
}

func (Space *SpaceDB) SetLineComplex128(col int64, line int64, data complex128) error {

	return Space.SetLineRaw(col, line, dacv2.Complex128ToBytes(data))
}

func (Space *SpaceDB) GetLineComplex64(col int64, line int64) (complex64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToComplex64(RBuffer), nil
}

func (Space *SpaceDB) NewLineComplex64(col int64, data complex64) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Complex64ToBytes(data))
}

func (Space *SpaceDB) SetLineComplex64(col int64, line int64, data complex64) error {

	return Space.SetLineRaw(col, line, dacv2.Complex64ToBytes(data))
}
