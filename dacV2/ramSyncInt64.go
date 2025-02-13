package dacv2


type SpaceRamSyncInt64 struct {
	*Space
	colName int64
	size    [3]int64
	Map     map[int64]int64
}

// Ram Sync Int64: Sincroniza los numeros de linea con los valores de esas lineas que son int64
func (sF *Space) InitRamSyncInt64(colName int64) (*SpaceRamSyncInt64, error) {

	size, found := sF.IndexSizeColumns[colName]
	if found {

		//Creamos un puntero a la estructura.
		SGMS := &SpaceRamSyncInt64{
			Space:   sF,
			colName: colName,
			size:    size,
			Map:     make(map[int64]int64),
		}

		//Activamos candados de lectura y escritura.
		sF.Mu.Lock()
		defer sF.Mu.Unlock()

		//Las lineas empiezan en -1, si sizeFileLine tiene una linea tiene 0 lineas
		for idLine := int64(0); idLine < sF.CountLines(); idLine++ {

			idFound, err := sF.GetLineInt64(colName, idLine)
			if err != nil {
				return nil, err
			}

			if idFound == -1 {
				continue
			}

			SGMS.Map[idFound] = idLine
		}

		return SGMS, nil
	}

	return nil, ErrNoFieldsOrColumn

}

// Ram Sync Int64: Crea una nueva linea con un id que no exista.
func (SGMS *SpaceRamSyncInt64) NewRamLineInt64(idFound int64) (int64, error) {

	if idFound < 0 {
		return -1, ErrNonNegativeValues
	}

	SGMS.Mu.Lock()

	_, found := SGMS.Map[idFound]
	if !found {

		linePointer, err := SGMS.NewLineInt64(SGMS.colName, idFound)
		if err != nil {
			SGMS.Mu.Unlock()
			return -1, err
		}

		SGMS.Map[idFound] = linePointer

		SGMS.Mu.Unlock()

		return linePointer, nil

	}

	SGMS.Mu.Unlock()

	return -1, ErrNonUniqueValues
}

// Ram Sync Int64: Actualiza una nueva linea y verifica que no existe el id
func (SGMS *SpaceRamSyncInt64) SetRamLineInt64(line int64, idFoundSet int64) (err error) {

	//Verificamos que la linea no sea menor que 0
	if line < 0 {
		return ErrInvalidLine
	}

	if idFoundSet < 0 {
		return ErrNonNegativeValues
	}

	//creamos los bloqueos
	SGMS.Mu.Lock()

	//Buscamos la linea con el id actual
	idFoundCurrent, err := SGMS.GetLineInt64(SGMS.colName, line)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	//Verificamos que el nuevo id para actualizar no exista ya en el mapa
	_, found := SGMS.Map[idFoundSet]
	if found {
		SGMS.Mu.Unlock()
		return ErrNonUniqueValues
	}

	//Borramos el id actual
	delete(SGMS.Map, idFoundCurrent)

	//Despues escribimos la nueva linea en el archivo
	err = SGMS.SetLineInt64(SGMS.colName, line, idFoundSet)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	//Añadimos esa linea al mapa tambien
	SGMS.Map[idFoundSet] = line

	SGMS.Mu.Unlock()

	return
}

// Ram Sync Int64: Obtiene un id sabiendo el numero de linea
func (SGMS *SpaceRamSyncInt64) GetRamIdInt64(line int64) (idFound int64, err error) {

	idFound, err = SGMS.GetLineInt64(SGMS.colName, line)
	if err != nil {
		return
	}
	SGMS.Mu.Lock()

	_, found := SGMS.Map[idFound]
	if !found {
		SGMS.Mu.Unlock()
		return -1, ErrNoExistValue
	}

	SGMS.Mu.Unlock()
	return
}

func (SGMS *SpaceRamSyncInt64) ExistRamLineInt64(line int64) (found bool, err error) {

	idFound, err := SGMS.GetLineInt64(SGMS.colName, line)
	if err != nil {
		return
	}

	SGMS.Mu.Lock()

	_, found = SGMS.Map[idFound]

	SGMS.Mu.Unlock()
	return
}

// Ram Sync Int64: Obtiene el numero de linea a traves de un id, buscando en un mapa.
func (SGMS *SpaceRamSyncInt64) GetRamLineInt64(idFound int64) (line int64, err error) {

	SGMS.Mu.RLock()

	line, found := SGMS.Map[idFound]
	if !found {
		SGMS.Mu.RUnlock()
		return -1, ErrNoExistValue
	}

	SGMS.Mu.RUnlock()

	return
}

func (SGMS *SpaceRamSyncInt64) ExistRamIdInt64(idFound int64) (found bool) {

	SGMS.Mu.RLock()

	_, found = SGMS.Map[idFound]

	SGMS.Mu.RUnlock()

	return
}

// Ram Sync Int64: Borra un numero de linea
func (SGMS *SpaceRamSyncInt64) DeleteRamLineInt64(line int64) (err error) {

	SGMS.Mu.Lock()

	idFound, err := SGMS.GetLineInt64(SGMS.colName, line)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	err = SGMS.SetLineInt64(SGMS.colName, line, -1)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	delete(SGMS.Map, idFound)

	SGMS.Mu.Unlock()

	return nil
}

// Ram Sync Int64: Borra la linea donde se encuentre el id
func (SGMS *SpaceRamSyncInt64) DeleteRamIdInt64(idFound int64) (err error) {

	SGMS.Mu.Lock()

	line, found := SGMS.Map[idFound]
	if !found {
		SGMS.Mu.Unlock()
		return ErrNoExistValue
	}

	err = SGMS.SetLineInt64(SGMS.colName, line, -1)
	if err != nil {
		SGMS.Mu.Unlock()
		return
	}

	delete(SGMS.Map, idFound)

	SGMS.Mu.Unlock()

	return
}
