package dacv2

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type Space struct {
	Dir       string
	FilePath  string
	Extension string

	File *os.File
	/*
		defer no funciona con bucles por si se va usar para liberar el mutex
		Bloqueo de lectura
		Space.mu.RLock()
		Space.mu.RUnlock()
		Bloqueo de escritura
		Space.mu.Lock()
		Space.mu.Unlock()
	*/
	Mu sync.RWMutex

	Size            int64
	IndexSizeFields map[int64][3]int64
	SizeField       int64

	//Indice de columnas y tamaño de columna
	IndexSizeColumns map[int64][3]int64
	SizeLine         int64
	AtomicCountLines atomic.Int64
}

func CreateDirectory(dirPath ...string) (err error) {

	dir := filepath.Join(dirPath...)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	return
}

func OpenFileOrCreate(dirPath ...string) (file *os.File, filePath string, extension string, dir string, err error) {

	filePath = filepath.Join(dirPath...)

	dir = filepath.Dir(filePath)

	extension = filepath.Ext(filePath)

	// Intenta abrir el archivo
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil && !os.IsNotExist(err) {
		return
	}

	if err != nil && os.IsNotExist(err) {

		// Crea las carpetas necesarias
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return
		}

		// Reintenta abrir el archivo
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return
		}

	}

	return
}

func OpenFileIfExist(dirPath ...string) (file *os.File, filePath string, extension string, dir string, err error) {

	filePath = filepath.Join(dirPath...)
	dir = filepath.Dir(filePath)
	extension = filepath.Ext(filePath)

	// Intenta abrir el archivo sin crearlo
	file, err = os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	return
}

func CountLineFile(file *os.File, sizeField int64, sizeLine int64) (line int64, err error) {

	// Obtiene el tamaño total del archivo
	size, err := file.Seek(0, 2)
	if err != nil {
		return 0, err
	}

	// Verifica si el tamaño es menor o igual a sizeField
	if size <= sizeField {
		return 0, nil
	}

	// Calcula el tamaño restante después de sizeField
	size -= sizeField

	// Calcula el número de líneas
	lines := size / sizeLine

	if size%sizeLine != 0 {
		lines++ // Suma 1 si hay una línea parcial
	}

	return lines, nil
}

func NewSpace(mapFields map[int64][3]int64, sizeField int64, mapLines map[int64][3]int64, sizeLine int64, dirPath ...string) (newSpace *Space, err error) {

	file, filePath, extension, dir, err := OpenFileOrCreate(dirPath...)
	if err != nil {
		return
	}

	countLines, err := CountLineFile(file, sizeField, sizeLine)
	if err != nil {
		return
	}

	newSpace = &Space{

		Dir:       dir,
		FilePath:  filePath,
		Extension: extension,

		File: file,

		IndexSizeFields: mapFields,
		SizeField:       sizeField,

		IndexSizeColumns: mapLines,
		SizeLine:         sizeLine,
	}

	newSpace.AtomicCountLines.Store(countLines)

	return
}

func NewSpaceIfExist(mapFields map[int64][3]int64, sizeField int64, mapLines map[int64][3]int64, sizeLine int64, dirPath ...string) (newSpace *Space, err error) {

	file, filePath, extension, dir, err := OpenFileIfExist(dirPath...)
	if err != nil {
		return
	}

	countLines, err := CountLineFile(file, sizeField, sizeLine)
	if err != nil {
		return
	}

	newSpace = &Space{

		Dir:       dir,
		FilePath:  filePath,
		Extension: extension,

		File: file,

		IndexSizeFields: mapFields,
		SizeField:       sizeField,

		IndexSizeColumns: mapLines,
		SizeLine:         sizeLine,
	}

	newSpace.AtomicCountLines.Store(countLines)

	return
}
