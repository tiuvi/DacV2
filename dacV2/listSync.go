package dacv2

import (
	jsonGo "encoding/json"
)

type ListSync struct {
	*Space
	sfList     *Space
	syncColumn int64
}

const (
	count int64 = 0
)

/*
Caso de uso de la lista sincronizada:

Imaginemos un sistema donde cada "like" se almacena en un archivo, y la posición en dicho
archivo corresponde al ID del usuario. Es decir, si un usuario con ID 100,000 da like, se
necesitaría un archivo de al menos 100,000 bytes para representarlo.

Si quisiéramos obtener todos los usuarios que han dado like, tendríamos que recorrer las
100,000 líneas del archivo, incluso si solo hay un like registrado.

Aquí es donde entra la lista sincronizada: en lugar de iterar sobre todas las líneas, esta
estructura optimiza el acceso guardando en la línea 0 el ID del usuario con el mayor valor
registrado. Así, en vez de recorrer 100,000 posiciones innecesarias, podemos acceder

	directamente a los usuarios que han interactuado.
*/
func InitListSync(syncColumn int64, mapFields map[int64][3]int64, sizeField int64, mapLines map[int64][3]int64, sizeLine int64, dirName ...string) (listSync *ListSync, err error) {

	sf, err := NewSpace(mapFields, sizeField, mapLines, sizeLine, dirName...)
	if err != nil {
		return
	}

	fieldList, fieldListSize,  err := CreateMap([]SpaceList{
		{Name: count, Len: 8},
	})
	if err != nil {
		return
	}

	columnsList, columnsListSize,  err := CreateMap([]SpaceList{
		{Name: syncColumn, Len: 8},
	})
	if err != nil {
		return
	}

	sfList, err := NewSpace(fieldList, fieldListSize, columnsList, columnsListSize, append(dirName, "syncList")...)
	if err != nil {
		return
	}

	listSync = &ListSync{
		Space:      sf,
		sfList:     sfList,
		syncColumn: syncColumn,
	}

	return
}

// Comprueba si un id existe en las dos listas
func (LS *ListSync) CheckIdSync(id int64) (found bool, err error) {

	//Buscamos el id en la lista unica
	IdInList, err := LS.GetLineInt64(LS.syncColumn, id)
	if err != nil {
		return
	}

	//Si el id es -1 significa que ya fue borrado
	if IdInList == -1 {
		return false, nil
	}

	//Buscamos el id en la lista
	idFound, err := LS.sfList.GetLineInt64(LS.syncColumn, IdInList)
	if err != nil {
		return
	}

	//Si el id encontrado no es igual al guardado significa que no existe, porque no estan sincronizados
	if idFound != id {
		return false, nil
	}

	//Si el id en la lista unica es igual al id en la lista, significa que existe
	return true, nil

}

// Crea un nuevo id sincronizado en ambas listas
func (LS *ListSync) NewIdSync(id int64) (line int64, err error) {

	//Guardamos el id en la lista
	idListPointer, err := LS.sfList.NewLineInt64(LS.syncColumn, id)
	if err != nil {
		return
	}

	//Guardamos el idList en la lista unica, guardandolo en el numero de linea del id.
	err = LS.SetLineInt64(LS.syncColumn, id, idListPointer)
	if err != nil {
		return
	}

	line, err = LS.sfList.AtomicFieldSumInt64(count, +1)
	if err != nil {
		return
	}

	return

}

func (LS *ListSync) DeleteIdSync(id int64) (err error) {

	//Primero buscamos el id para encontrarlo en la lista unica
	IdInList, err := LS.GetLineInt64(LS.syncColumn, id)
	if err != nil {
		return
	}

	//Borramos el id en la lista poniendolo en -1
	err = LS.sfList.SetLineInt64(LS.syncColumn, IdInList, -1)
	if err != nil {
		return
	}

	//Borramos el id en la lista unica poniendolo en -1
	err = LS.SetLineInt64(LS.syncColumn, id, -1)
	if err != nil {
		return
	}

	_, err = LS.sfList.AtomicFieldSumInt64(count, -1)
	if err != nil {
		return
	}

	return

}

func (LS *ListSync) CountIds() (count int64, err error) {

	return LS.sfList.GetFieldInt64(count)
}

type Int64Data struct {
	ID    int64 `json:"id"`
	Value int64 `json:"value"`
}

func (LS *ListSync) GetAllIdSyncRange(start int64, manyResults int64) (list []Int64Data, err error) {

	list = make([]Int64Data, 0, 10)

	var count int64
	var IdInList int64
	var id int64

	// Iterar sobre el rango especificado
	for index := start; index < LS.sfList.CountLines(); index++ {

		if count > manyResults {
			break
		}

		//Buscamos el id en la lista
		IdInList, err = LS.sfList.GetLineInt64(LS.syncColumn, index)
		if err != nil {
			return
		}

		//Si es -1 Significa que fue borrado
		if IdInList == -1 {
			continue
		}

		//Buscamos el id en la lista unica a ver si existe
		id, err = LS.GetLineInt64(LS.syncColumn, IdInList)
		if err != nil {
			return
		}

		//Si es -1 Significa que fue borrado
		if id == -1 {
			continue
		}

		if id != IdInList {
			continue
		}
		//Añadimos una cuenta
		count = count + 1

		list = append(list, Int64Data{ID: IdInList, Value: id})
	}

	return list, nil

}

func (LS *ListSync) GetAllIdSyncJson(start int64, manyResults int64) (json []byte, err error) {

	data, err := LS.GetAllIdSyncRange(start, manyResults)
	if err != nil {
		return nil, err
	}

	jsonData, err := jsonGo.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}
