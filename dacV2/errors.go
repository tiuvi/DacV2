package dacv2

import (
	"errors"
)

// Definición de errores
var (
	ErrNegativeLength = errors.New("unidad negativa no permitida")
	ErrDuplicateName  = errors.New("nombre duplicado de columna o campo")
	ErrNoFieldsOrColumn  = errors.New("no existen columnas o campos con ese nombre")
	ErrNoFieldsOrColumnMap  = errors.New("no existen mapas de columnas o campos")
	ErrRangeFieldNoZero   = errors.New("los rangos de los campos no pueden ser 0 de ancho de banda")
	ErrLenBufferNoEqualSize   = errors.New("el tamaño del buffer no coincide con el tamaño del campo o columna")
	ErrInterfaceTooLarge = 	errors.New("el tamaño de la interfaz excede el tamaño permitido")

	ErrNonUniqueValues = errors.New("no es un valor unico, ese valor ya existe")
	ErrNoExistValue = errors.New("ese valor no existe")
	ErrNoEmptyString = errors.New("no se permite el uso de cadenas vacias")
	
	ErrInvalidLine = errors.New("no se permite numeros de linea negativos")
	ErrNonNegativeValues = errors.New("no se permite valores negativos")
)

