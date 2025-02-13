package dacv2

import (
	"time"
)

func (Space *Space) GetFieldUnixMilliInt64(col int64) (int64, error) {

	return Space.GetFieldInt64(col)
}

// Tamaño 8 bytes
func (Space *Space) GetFieldUnixMilli(col int64) (time.Time, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return time.Time{}, err
	}

	return time.UnixMilli(BytesToInt64(buffer)), nil
}

// Tamaño 8 bytes
func (Space *Space) SetFieldUnixMilli(col int64, data time.Time) error {

	return Space.SetFieldRaw(col, Int64ToBytes(data.UnixMilli()))
}


func (Space *Space) SetFieldUnixMilliNow(col int64, line int64) (time.Time, error) {

	timenow := time.Now()

	err := Space.SetFieldRaw(col, Int64ToBytes(timenow.UnixMilli()))

	return timenow, err
}
