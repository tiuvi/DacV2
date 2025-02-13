package dacv2


func (Space *Space) GetAllLinesInt64(col int64) ([]int64, error) {

	var x int64
	var data = make([]int64, Space.CountLines())

	for x = 0; x < Space.CountLines(); x++ {

		RBuffer, err := Space.GetLineRaw(col, x)
		if err != nil {
			return nil, err
		}

		data[x] = BytesToInt64(RBuffer)

	}

	return data, nil
}

func (Space *Space) SearchLineInt64(col int64, colunique int64, idFind int64) ([]string, error) {

	var x int64
	var data = make([]string, 0)

	for x = 0; x < Space.CountLines(); x++ {

		id, err := Space.GetLineInt64(col, x)
		if err != nil {
			return nil, err
		}

		if id == idFind {

			value, err := Space.GetLineString(colunique, x)
			if err != nil {
				return nil, err
			}
			data = append(data, value)
		}

	}

	return data, nil
}