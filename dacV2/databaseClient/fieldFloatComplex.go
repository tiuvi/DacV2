package databaseClient

import (
	dacv2 "dacV2"
)

func (Space *SpaceDB) GetFieldFloat64(col int64) (float64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToFloat64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *SpaceDB) SetFieldFloat64(col int64, value float64) error {

	return Space.SetFieldRaw(col, dacv2.Float64ToBytes(value))
}

func (Space *SpaceDB) GetFieldFloat32(col int64) (float32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToFloat32(buffer), nil
}

// Tamaño: 4 bytes
func (Space *SpaceDB) SetFieldFloat32(col int64, value float32) error {

	return Space.SetFieldRaw(col, dacv2.Float32ToBytes(value))
}

func (Space *SpaceDB) GetFieldComplex128(col int64) (complex128, error) {
	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToComplex128(buffer), nil
}

// tamaño: 16 bytes
func (Space *SpaceDB) SetFieldComplex128(col int64, value complex128) error {

	return Space.SetFieldRaw(col, dacv2.Complex128ToBytes(value))
}

func (Space *SpaceDB) GetFieldComplex64(col int64) (complex64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToComplex64(buffer), nil
}

func (Space *SpaceDB) SetFieldComplex64(col int64, value complex64) error {

	return Space.SetFieldRaw(col, dacv2.Complex64ToBytes(value))
}
