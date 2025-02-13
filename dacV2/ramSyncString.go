package dacv2

import (
	"strings"
)


type SpaceRamSyncString struct {
	*Space
	colName int64
	size    [3]int64
	Map     map[string]int64
}


 
// Ram Sync String: Sincroniza un mapa con las lineas y los valores de esas lineas que son string
func (sF *Space) InitRamSyncString(colName int64) (*SpaceRamSyncString, error) {

	size, found := sF.IndexSizeColumns[colName]
	if found {

		//Creamos un puntero a la estructura.
		SGMS := &SpaceRamSyncString{
			Space: sF,
			colName:   colName,
			size:      size,
			Map:       make(map[string]int64),
		}

		//Activamos candados de lectura y escritura.
		sF.Mu.Lock()
		defer sF.Mu.Unlock()

		for idLine := int64(0); idLine < sF.CountLines(); idLine++ {

			idString, err := sF.GetLineString(colName, idLine)
			if err != nil {
				return nil, err
			}

			if idString == "" {
				continue
			}

			SGMS.Map[idString] = idLine
		}

		return SGMS, nil
	}

	return nil, ErrNoFieldsOrColumn

}

func (SGMS *SpaceRamSyncString) NewRamLineString(data string) (int64, error) {

	if data == "" {
		return -1, ErrNoEmptyString
	}

	//Eliminamos nulos porque dac siempre borra todos los nulos a la derecha.
	data = strings.TrimRight(data, "\x00")

	SGMS.Mu.Lock()

	//Si la linea es -1  buscamos en el mapa, creamos una nueva entrada y una nueva linea en el archivo.
	_, found := SGMS.Map[data]
	if !found {

		linePointer, err := SGMS.NewLineString(SGMS.colName, data)
		if err != nil {
			SGMS.Mu.Unlock()
			return -1, err
		}

		SGMS.Map[data] = linePointer

		SGMS.Mu.Unlock()

		return linePointer, nil

	}

	SGMS.Mu.Unlock()

	return -1, ErrNonUniqueValues
}

func (SGMS *SpaceRamSyncString) SetRamLineString(line int64, idStringSet string) (err error) {

	if line < 0 {
		return ErrInvalidLine
	}

	if idStringSet == "" {
		return ErrNoEmptyString
	}

	//Eliminamos nulos porque dac siempre borra todos los nulos a la derecha.
	idStringSet = strings.TrimRight(idStringSet, "\x00")

	//creamos los bloqueos
	SGMS.Mu.Lock()

	//Primero leemos el fichero
	idStringFoundCurrent, err := SGMS.GetLineString(SGMS.colName, line)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	_, found := SGMS.Map[idStringSet]
	if found {
		SGMS.Mu.Unlock()
		return ErrNonUniqueValues
	}

	delete(SGMS.Map, idStringFoundCurrent)

	//Despues escribimos la nueva linea en el archivo
	err = SGMS.SetLineString(SGMS.colName, line, idStringSet)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	//Añadimos esa linea al mapa tambien
	SGMS.Map[idStringSet] = line

	SGMS.Mu.Unlock()
	return
}

func (SGMS *SpaceRamSyncString) GetRamIdString(line int64) (idString string, err error) {

	idString , err = SGMS.GetLineString(SGMS.colName, line)
	if err != nil {
		return
	}

	SGMS.Mu.Lock()

	_, found := SGMS.Map[idString]
	if !found {
		SGMS.Mu.Unlock()
		return "" , ErrNoExistValue
	}

	SGMS.Mu.Unlock()

	return
}

func (SGMS *SpaceRamSyncString) ExistRamLineString(line int64) (found bool, err error) {

	idString , err := SGMS.GetLineString(SGMS.colName, line)
	if err != nil {
		return
	}

	SGMS.Mu.Lock()

	_, found = SGMS.Map[idString]

	SGMS.Mu.Unlock()
	
	return
}


func (SGMS *SpaceRamSyncString) GetRamLineString(idString string)(line int64, err error){

	idString = strings.TrimRight(idString, "\x00")

	//Añadimos un bloqueo de lectura
	SGMS.Mu.RLock()

	//Si existe devolvemos true si no false.
	line, found := SGMS.Map[idString]
	if !found {
		SGMS.Mu.RUnlock()
		return -1 , ErrNoExistValue 

	}

	SGMS.Mu.RUnlock()

	return
}

func (SGMS *SpaceRamSyncString) ExistRamIdString(idString string)(found bool){

	idString = strings.TrimRight(idString, "\x00")

	SGMS.Mu.RLock()

	_, found = SGMS.Map[idString]

	SGMS.Mu.RUnlock()

	return
}

func (SGMS *SpaceRamSyncString) DeleteRamLineString(line int64) (err error) {

	SGMS.Mu.Lock()

	idString, err := SGMS.GetLineString(SGMS.colName, line)
	if err != nil  {
		SGMS.Mu.Unlock()
		return 
	}
	 
	err = SGMS.SetLineString(SGMS.colName, line, "")
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	delete(SGMS.Map, idString)

	SGMS.Mu.Unlock()

	return
}

func (SGMS *SpaceRamSyncString) DeleteRamIdString(idString string) (err error) {

	SGMS.Mu.Lock()

	line, found := SGMS.Map[idString]
	if !found {
		SGMS.Mu.Unlock()
		return ErrNoExistValue
	}

	err = SGMS.SetLineString(SGMS.colName, line, "")
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	delete(SGMS.Map, idString)

	SGMS.Mu.Unlock()

	return
}