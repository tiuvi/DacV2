package databaseClient

import (
	dacv2 "dacV2"
)

func (Space *SpaceDB) GetLineUint64(col int64, line int64) (uint64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint64(RBuffer), nil
}

func (Space *SpaceDB) NewLineUint64(col int64, data uint64) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Uint64ToBytes(data))
}

func (Space *SpaceDB) SetLineUint64(col int64, line int64, data uint64) error {

	return Space.SetLineRaw(col, line, dacv2.Uint64ToBytes(data))
}

func (Space *SpaceDB) GetLineUint32(col int64, line int64) (uint32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint32(RBuffer), nil
}

func (Space *SpaceDB) NewLineUint32(col int64, data uint32) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Uint32ToBytes(data))
}

func (Space *SpaceDB) SetLineUint32(col int64, line int64, data uint32) error {

	return Space.SetLineRaw(col, line, dacv2.Uint32ToBytes(data))
}

func (Space *SpaceDB) GetLineUint(col int64, line int64) (uint, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return uint(dacv2.BytesToUint32(RBuffer)), nil
}

func (Space *SpaceDB) NewLineUint(col int64, data uint) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Uint32ToBytes(uint32(data)))
}

func (Space *SpaceDB) SetLineUint(col int64, line int64, data uint) error {

	return Space.SetLineRaw(col, line, dacv2.Uint32ToBytes(uint32(data)))
}

func (Space *SpaceDB) GetLineUint16(col int64, line int64) (uint16, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint16(RBuffer), nil
}

func (Space *SpaceDB) NewLineUint16(col int64, data uint16) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Uint16ToBytes(data))
}

func (Space *SpaceDB) SetLineUint16(col int64, line int64, data uint16) error {

	return Space.SetLineRaw(col, line, dacv2.Uint16ToBytes(data))
}

func (Space *SpaceDB) GetLineUint8(col int64, line int64) (uint8, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToUint8(RBuffer), nil
}

func (Space *SpaceDB) NewLineUint8(col int64, data uint8) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Uint8ToBytes(data))
}

func (Space *SpaceDB) SetLineUint8(col int64, line int64, data uint8) error {

	return Space.SetLineRaw(col, line, dacv2.Uint8ToBytes(data))
}

func (Space *SpaceDB) GetLineInt64(col int64, line int64) (int64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToInt64(RBuffer), nil
}

func (Space *SpaceDB) NewLineInt64(col int64, data int64) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int64ToBytes(data))
}

func (Space *SpaceDB) SetLineInt64(col int64, line int64, data int64) error {

	return Space.SetLineRaw(col, line, dacv2.Int64ToBytes(data))
}

func (Space *SpaceDB) GetLineInt32(col int64, line int64) (int32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToInt32(RBuffer), nil
}

func (Space *SpaceDB) NewLineInt32(col int64, data int32) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int32ToBytes(data))
}

func (Space *SpaceDB) SetLineInt32(col int64, line int64, data int32) error {

	return Space.SetLineRaw(col, line, dacv2.Int32ToBytes(data))
}

func (Space *SpaceDB) GetLineInt(col int64, line int64) (int, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return int(dacv2.BytesToInt32(RBuffer)), nil
}

func (Space *SpaceDB) NewLineInt(col int64, data int) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int32ToBytes(int32(data)))
}

func (Space *SpaceDB) SetLineInt(col int64, line int64, data int) error {

	return Space.SetLineRaw(col, line, dacv2.Int32ToBytes(int32(data)))
}

func (Space *SpaceDB) GetLineInt16(col int64, line int64) (int16, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToInt16(RBuffer), nil
}

func (Space *SpaceDB) NewLineInt16(col int64, data int16) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int16ToBytes(data))
}

func (Space *SpaceDB) SetLineInt16(col int64, line int64, data int16) error {

	return Space.SetLineRaw(col, line, dacv2.Int16ToBytes(data))
}

func (Space *SpaceDB) GetLineInt8(col int64, line int64) (int8, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return dacv2.BytesToInt8(RBuffer), nil
}

func (Space *SpaceDB) NewLineInt8(col int64, data int8) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int8ToBytes(data))
}

func (Space *SpaceDB) SetLineInt8(col int64, line int64, data int8) error {

	return Space.SetLineRaw(col, line, dacv2.Int8ToBytes(data))
}
