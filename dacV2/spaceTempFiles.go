package dacv2

import (
	"io"
	"os"
)

type  SpaceTemp struct{
	File *os.File
}

func NewTempSpace()(spTemp *SpaceTemp ,err error){

	file, err := os.CreateTemp("", "dac")
	if err != nil {

		return
	}

	return &SpaceTemp{
		File: file,
	}, nil
}

func (ST * SpaceTemp)Write(data []byte)(err error){
	
	_ , err = ST.File.Write(data)
	return 
}

func (ST * SpaceTemp)Update(data []byte)(err error){
	
	ST.File.Truncate(0)
	_ , err = ST.File.Write(data)
	return 
}


func (ST * SpaceTemp)Name()string{
	
	return ST.File.Name()
}

func (ST * SpaceTemp)ReadAll()(data []byte, err error){

	fileInfo, err := ST.File.Stat()
	if err != nil {
		return
	}

	// Crea un slice de bytes con el tamaño del archivo
	data = make([]byte, fileInfo.Size())

	_, err = ST.File.Read(data)
	if err != nil && err != io.EOF{
		return
	}

	if err != nil && err == io.EOF{
		return []byte{} , nil
	}

	return data , nil
}


func CreateTempBytes(data []byte)(url string ,err error){

	// Crear un archivo temporal en el directorio de archivos temporales predeterminado
	File, err := os.CreateTemp("", "dac")
	if err != nil {

		return
	}

	// Escribir en el archivo temporal
	_, err = File.Write(data)
	if err != nil {
		return
	}

	return File.Name(), nil
}
