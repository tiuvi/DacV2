package dacv2

import (
	"errors"
	"io"
)

func (Space *Space) GetAt(offSet int64 , sizeBuffer int64) (buf []byte, err error) {

	buf = make([]byte, sizeBuffer)

	_, err = Space.File.ReadAt(buf, offSet)
	if err != nil && err != io.EOF {
		return
	}

	if err != nil && err == io.EOF {
		return buf , nil
	}

	return
}

func (s *Space) GetAtRange(nRange int64, bandwidth int64) ([]byte, error) {
	position := (nRange * bandwidth)

	buf := make([]byte, bandwidth)

	n, err := s.File.ReadAt(buf, position)
	if  err != nil && err != io.EOF {
		return nil, err
	}

	if err != nil && err == io.EOF {
		return buf[:n], nil
	}

	// Si se leyó menos de `bandwidth`, recortar el buffer
	return buf , nil
}


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
		return buf , nil
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
		return buf , nil
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
		return buf , nil
	}

	return
}

/*
Esta funcion unicamente devuelve un rango de lineas
Imagina que tenemos 100 000 lineas bool que serian 100 000 bytes hacer
100 000 peticiones usando getLineBool no es eficiente estarimos mandando 100 000 operaciones al
ssd en ese caso lo mas eficiente es pedir todos los bytes y formatearlos en el cliente.
*/
func (Space *Space) GetLinesRange(startLine int64, endLine int64) (buf []byte, err error) {
	
	// Verificamos que el rango sea válido
	if startLine > endLine || startLine < 0 {
		return nil , errors.New("rango de lineas invalido")
	}

	buf = make([]byte, (endLine-startLine+1)*Space.SizeLine)

	_, err = Space.File.ReadAt(buf, Space.SizeField+(startLine*Space.SizeLine))
	if err != nil && err != io.EOF {
		return
	}
	
	if err != nil && err == io.EOF {
		return buf , nil
	}

	return
}
