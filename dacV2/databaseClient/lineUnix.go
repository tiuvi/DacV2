package databaseClient

import (
	dacv2 "dacV2"
	"time"
)

func (Space *SpaceDB) GetLineUnixMilliInt64(col int64, line int64) (int64, error) {

	return Space.GetLineInt64(col, line)
}

func (Space *SpaceDB) GetLineUnixMilli(col int64, line int64) (time.Time, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return time.Time{}, err
	}

	return time.UnixMilli(dacv2.BytesToInt64(RBuffer)), nil
}

// Esta funcion actualiza con el tiempo actual
func (Space *SpaceDB) NewLineUnixMilli(col int64, data time.Time) (int64, error) {

	return Space.NewLineRaw(col, dacv2.Int64ToBytes(data.UnixMilli()))

}

func (Space *SpaceDB) SetLineUnixMilli(col int64, line int64, data time.Time) error {

	return Space.SetLineRaw(col, line, dacv2.Int64ToBytes(data.UnixMilli()))
}

// Esta funcion actualiza con el tiempo actual y devuelve el tiempo
func (Space *SpaceDB) NewLineUnixMilliNow(col int64) (time.Time, int64, error) {

	timenow := time.Now()

	id, err := Space.NewLineRaw(col, dacv2.Int64ToBytes(timenow.UnixMilli()))

	return timenow, id, err
}

func (Space *SpaceDB) SetLineUnixMilliNow(col int64, line int64) (time.Time, error) {

	timenow := time.Now()

	err := Space.SetLineRaw(col, line, dacv2.Int64ToBytes(timenow.UnixMilli()))

	return timenow, err
}
