package databaseClient

import (
	dacv2 "dacV2"
)

// Tamaño: 8 bytes
func (Space *SpaceDB) GetFieldUint64(col int64) (uint64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *SpaceDB) SetFieldUint64(col int64, value uint64) error {

	return Space.SetFieldRaw(col, dacv2.Uint64ToBytes(value))

}

func (Space *SpaceDB) GetFieldUint32(col int64) (uint32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint32(buffer), nil

}

// Tamaño: 4 bytes
func (Space *SpaceDB) SetFieldUint32(col int64, value uint32) {

	Space.SetFieldRaw(col, dacv2.Uint32ToBytes(value))
}

func (Space *SpaceDB) GetFieldUint(col int64) (uint, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return uint(dacv2.BytesToUint32(buffer)), nil
}

// Tamaño: 4 bytes
func (Space *SpaceDB) SetFieldUint(col int64, value uint) error {

	return Space.SetFieldRaw(col, dacv2.Uint32ToBytes(uint32(value)))

}

func (Space *SpaceDB) GetFieldUint16(col int64) (uint16, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint16(buffer), nil
}

// Tamaño: 2 bytes
func (Space *SpaceDB) SetFieldUint16(col int64, value uint16) error {

	return Space.SetFieldRaw(col, dacv2.Uint16ToBytes(value))
}

func (Space *SpaceDB) GetFieldUint8(col int64) (uint8, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint8(buffer), nil
}

// Tamaño: 1 bytes
func (Space *SpaceDB) SetFieldUint8(col int64, value uint8) error {

	return Space.SetFieldRaw(col, dacv2.Uint8ToBytes(value))
}

// Tamaño: 8 bytes
func (Space *SpaceDB) GetFieldInt64(col int64) (int64, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToInt64(buffer), nil
}

// Tamaño: 8 bytes
func (Space *SpaceDB) SetFieldInt64(col int64, value int64) error {

	return Space.SetFieldRaw(col, dacv2.Int64ToBytes(value))
}

func (Space *SpaceDB) GetFieldInt(col int64) (int, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToInt(buffer), nil
}

// Tamaño: 4 bytes
func (Space *SpaceDB) SetFieldInt(col int64, value int) error {

	return Space.SetFieldRaw(col, dacv2.IntToBytes(value))
}

func (Space *SpaceDB) GetFieldInt32(col int64) (int32, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToInt32(buffer), nil
}

// Tamaño: 4 bytes
func (Space *SpaceDB) SetFieldInt32(col int64, value int32) error {

	return Space.SetFieldRaw(col, dacv2.Int32ToBytes(value))
}

func (Space *SpaceDB) GetFieldInt16(col int64) (int16, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToInt16(buffer), nil

}

// Tamaño: 2 bytes
func (Space *SpaceDB) SetFieldInt16(col int64, value int16) error {

	return Space.SetFieldRaw(col, dacv2.Int16ToBytes(value))
}

func (Space *SpaceDB) GetFieldInt8(col int64) (int8, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return 0, err
	}
	return dacv2.BytesToInt8(buffer), nil

}

// Tamaño: 1 bytes
func (Space *SpaceDB) SetFieldInt8(col int64, value int8) error {

	return Space.SetFieldRaw(col, dacv2.Int8ToBytes(value))
}
