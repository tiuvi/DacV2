package dacv2



func (Space *Space) GetLineUint64(col int64, line int64) (uint64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToUint64(RBuffer), nil
}

func (Space *Space) NewLineUint64(col int64, data uint64) (int64, error) {

	return Space.NewLineRaw(col, Uint64ToBytes(data))
}

func (Space *Space) SetLineUint64(col int64, line int64, data uint64) error {

	return Space.SetLineRaw(col, line, Uint64ToBytes(data))
}

func (Space *Space) GetLineUint32(col int64, line int64) (uint32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToUint32(RBuffer), nil
}

func (Space *Space) NewLineUint32(col int64, data uint32) (int64, error) {

	return Space.NewLineRaw(col, Uint32ToBytes(data))
}

func (Space *Space) SetLineUint32(col int64, line int64, data uint32) error {

	return Space.SetLineRaw(col, line, Uint32ToBytes(data))
}

func (Space *Space) GetLineUint(col int64, line int64) (uint, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return uint(BytesToUint32(RBuffer)), nil
}

func (Space *Space) NewLineUint(col int64, data uint) (int64, error) {

	return Space.NewLineRaw(col, Uint32ToBytes(uint32(data)))
}

func (Space *Space) SetLineUint(col int64, line int64, data uint) error {

	return Space.SetLineRaw(col, line, Uint32ToBytes(uint32(data)))
}

func (Space *Space) GetLineUint16(col int64, line int64) (uint16, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToUint16(RBuffer), nil
}

func (Space *Space) NewLineUint16(col int64, data uint16) (int64, error) {

	return Space.NewLineRaw(col, Uint16ToBytes(data))
}

func (Space *Space) SetLineUint16(col int64, line int64, data uint16) error {

	return Space.SetLineRaw(col, line, Uint16ToBytes(data))
}

func (Space *Space) GetLineUint8(col int64, line int64) (uint8, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToUint8(RBuffer), nil
}

func (Space *Space) NewLineUint8(col int64, data uint8) (int64, error) {

	return Space.NewLineRaw(col, Uint8ToBytes(data))
}

func (Space *Space) SetLineUint8(col int64, line int64, data uint8) error {

	return Space.SetLineRaw(col, line, Uint8ToBytes(data))
}

func (Space *Space) GetLineInt64(col int64, line int64) (int64, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToInt64(RBuffer), nil
}

func (Space *Space) NewLineInt64(col int64, data int64) (int64, error) {

	return Space.NewLineRaw(col, Int64ToBytes(data))
}

func (Space *Space) SetLineInt64(col int64, line int64, data int64) error {

	return Space.SetLineRaw(col, line, Int64ToBytes(data))
}


func (Space *Space) GetLineInt32(col int64, line int64) (int32, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToInt32(RBuffer), nil
}

func (Space *Space) NewLineInt32(col int64, data int32) (int64, error) {

	return Space.NewLineRaw(col, Int32ToBytes(data))
}

func (Space *Space) SetLineInt32(col int64, line int64, data int32) error {

	return Space.SetLineRaw(col, line, Int32ToBytes(data))
}

func (Space *Space) GetLineInt(col int64, line int64) (int, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return int(BytesToInt32(RBuffer)), nil
}

func (Space *Space) NewLineInt(col int64, data int) (int64, error) {

	return Space.NewLineRaw(col, Int32ToBytes(int32(data)))
}

func (Space *Space) SetLineInt(col int64, line int64, data int) error {

	return Space.SetLineRaw(col, line, Int32ToBytes(int32(data)))
}

func (Space *Space) GetLineInt16(col int64, line int64) (int16, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToInt16(RBuffer), nil
}

func (Space *Space) NewLineInt16(col int64, data int16) (int64, error) {

	return Space.NewLineRaw(col, Int16ToBytes(data))
}

func (Space *Space) SetLineInt16(col int64, line int64, data int16) error {

	return Space.SetLineRaw(col, line, Int16ToBytes(data))
}

func (Space *Space) GetLineInt8(col int64, line int64) (int8, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return 0, err
	}

	return BytesToInt8(RBuffer), nil
}

func (Space *Space) NewLineInt8(col int64, data int8) (int64, error) {

	return Space.NewLineRaw(col, Int8ToBytes(data))
}

func (Space *Space) SetLineInt8(col int64, line int64, data int8) error {

	return Space.SetLineRaw(col, line, Int8ToBytes(data))
}
