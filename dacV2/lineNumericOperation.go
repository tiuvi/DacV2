package dacv2

/*
Incrementa un valor en la base de datos, segura para operacis concurrentes.
*/
func (Space *Space) AtomicLineSumInt64(col int64, line int64, increment int64) (result int64, err error) {

	Space.Mu.Lock()
	
	data, err := Space.GetLineInt64(col, line)
	if err != nil {
		Space.Mu.Unlock()
		return 0, err
	}

	result = data + increment

	err = Space.SetLineInt64(col, line, result)
	if err != nil {
		Space.Mu.Unlock()
		return 0, err
	}

	Space.Mu.Unlock()

	return
}
