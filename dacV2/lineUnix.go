package dacv2

import (
	"time"
)


func (Space *Space) GetLineUnixMilliInt64(col int64 , line int64) (int64, error) {

	return Space.GetLineInt64(col , line)
}

func (Space *Space) GetLineUnixMilli(col int64, line int64) (time.Time, error) {

	RBuffer, err := Space.GetLineRaw(col, line)
	if err != nil {
		return time.Time{}, err
	}

	return time.UnixMilli(BytesToInt64(RBuffer)), nil
}


// Esta funcion actualiza con el tiempo actual
func (Space *Space) NewLineUnixMilli(col int64, data time.Time) (int64, error) {

	return Space.NewLineRaw(col, Int64ToBytes(data.UnixMilli()))

}

func (Space *Space) SetLineUnixMilli(col int64, line int64, data time.Time) error {

	return Space.SetLineRaw(col, line, Int64ToBytes(data.UnixMilli()))
}

// Esta funcion actualiza con el tiempo actual y devuelve el tiempo
func (Space *Space) NewLineUnixMilliNow(col int64) (time.Time, int64, error) {

	timenow := time.Now()

	id, err := Space.NewLineRaw(col, Int64ToBytes(timenow.UnixMilli()))

	return timenow, id, err
}

func (Space *Space) SetLineUnixMilliNow(col int64, line int64) (time.Time, error) {

	timenow := time.Now()

	err := Space.SetLineRaw(col, line, Int64ToBytes(timenow.UnixMilli()))

	return timenow, err
}

