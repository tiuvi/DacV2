package dacv2

import "io"


// Funcion de bytes con limpieza de bytes nulos
func (Space *Space) GetField(columnName int64) (buf []byte, err error) {

	buf, err = Space.GetFieldRaw(columnName)

	buf = SpaceTrimNull(buf)

	return
}

// Lectura de bytes sin limpieza de bytes nulos
func (Space *Space) GetFieldRaw(columnName int64) (buf []byte, err error) {

	size, err := Space.checkField(columnName)
	if err != nil {
		return
	}

	buf = make([]byte, size[2])

	_, err = Space.File.ReadAt(buf, size[0])
	if err != nil && err != io.EOF {
		return
	}

	if err != nil && err == io.EOF {
		return []byte{} , nil
	}

	return
}

// Lectura de rangos sin limpieza de bytes nulos
func (Space *Space) GetFieldRange(columnName int64, rangue int64, bandwidth int64) (buf []byte, err error) {

	if bandwidth == 0 {
		return nil, ErrRangeFieldNoZero
	}

	size, err := Space.checkField(columnName)
	if err != nil {
		return
	}

	var initRangue int64 = rangue * bandwidth
	var endRangue int64 = (rangue + 1) * bandwidth
	if endRangue > size[1] {
		endRangue = size[1]
	}

	buf = make([]byte, endRangue-initRangue)

	_, err = Space.File.ReadAt(buf, size[0]+initRangue)
	if err != nil && err != io.EOF {
		return
	}

	if err != nil && err == io.EOF {
		return []byte{} , nil
	}

	return
}

func (Space *Space) GetLine(columnName int64, line int64) (buf []byte, err error) {

	buf, err = Space.GetLineRaw(columnName, line)

	buf = SpaceTrimNull(buf)

	return
}

func (Space *Space) GetLineRaw(columnName int64, line int64) (buf []byte, err error) {

	size, err := Space.checkColumn(columnName)
	if err != nil {
		return
	}

	buf = make([]byte, size[2])

	_, err = Space.File.ReadAt(buf, Space.SizeField+(line*Space.SizeLine)+size[0])
	if err != nil && err != io.EOF {
		return
	}
	
	if err != nil && err == io.EOF {
		return []byte{} , nil
	}

	return
}
