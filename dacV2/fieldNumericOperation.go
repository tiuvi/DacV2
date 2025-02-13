package dacv2

// Esta funcion incrementa y ademas devuelve el resultado incrementado
func (Space *Space) AtomicFieldSumInt64(col int64, increment int64) (result int64, err error) {

	Space.Mu.Lock()

	data, err := Space.GetFieldInt64(col)
	if err != nil {
		Space.Mu.Unlock()
		return 0, err
	}

	result = data+increment

	err = Space.SetFieldInt64(col, result)
	if err != nil {
		Space.Mu.Unlock()
		return 0, err
	}

	Space.Mu.Unlock()

	return
}
