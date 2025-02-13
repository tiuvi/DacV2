package databaseClient

import (
	dacv2 "dacV2"
	"time"
)

func (Space *SpaceDB) GetFieldUnixMilliInt64(col int64) (int64, error) {

	return Space.GetFieldInt64(col)
}

// Tamaño 8 bytes
func (Space *SpaceDB) GetFieldUnixMilli(col int64) (time.Time, error) {

	buffer, err := Space.GetFieldRaw(col)
	if err != nil {
		return time.Time{}, err
	}

	return time.UnixMilli(dacv2.BytesToInt64(buffer)), nil
}

// Tamaño 8 bytes
func (Space *SpaceDB) SetFieldUnixMilli(col int64, data time.Time) error {

	return Space.SetFieldRaw(col, dacv2.Int64ToBytes(data.UnixMilli()))
}

func (Space *SpaceDB) SetFieldUnixMilliNow(col int64, line int64) (time.Time, error) {

	timenow := time.Now()

	err := Space.SetFieldRaw(col, dacv2.Int64ToBytes(timenow.UnixMilli()))

	return timenow, err
}
