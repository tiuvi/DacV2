package dacv2

/*
Obtiene todas las lineas en un array
*/
func (Space *Space) GetAllLinesString(col int64) ([]string, error) {

	var x int64
	var data = make([]string, Space.CountLines())

	for x = 0; x < Space.CountLines(); x++ {

		RBuffer, err := Space.GetLine(col, x)
		if err != nil {
			return nil, err
		}

		data[x] = string(RBuffer)
	}

	return data, nil
}
