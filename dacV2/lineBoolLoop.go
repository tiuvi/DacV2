package dacv2

import (
	"encoding/json"
)

func (Space *Space) GetAllLinesBool(col int64) ([]bool, error) {

	var x int64
	var data = make([]bool, Space.CountLines())

	for x = 0; x < Space.CountLines(); x++ {

		RBuffer, err := Space.GetLine(col, x)
		if err != nil {
			return nil, err
		}

		value := string(RBuffer)
		if len(value) == 0 {
			data[x] = false
		}

		data[x] = (string(RBuffer) == "t")

	}

	return data, nil
}

type BoolData struct {
	ID    int64 `json:"id"`
	Value bool  `json:"value"`
}

func (Space *Space) GetAllLinesRangesBool(col int64, start int64, end int64) ([]BoolData, error) {

	// Crear el slice para almacenar los resultados
	var data = make([]BoolData, 0, end-start+1)

	// Iterar sobre el rango especificado
	for index := start; index <= end; index++ {

		if index > Space.CountLines() {
			break
		}

		buffer, err := Space.GetLine(col, index)
		if err != nil {
			return nil, err
		}

		if len(buffer) == 0 {
			data = append(data, BoolData{
				ID:    index,
				Value: false,
			})
		}

		if buffer[0] == dacBoolFalse[0] {
			data = append(data, BoolData{
				ID:    index,
				Value: false,
			})
		}

		if buffer[0] == dacBoolTrue[0] {
			data = append(data, BoolData{
				ID:    index,
				Value: true,
			})
		}

	}

	return data, nil
}

func (Space *Space) GetAllLinesRangesBoolTrue(col int64, start int64, manyResults int64) ([]BoolData, error) {

	// Crear un slice para almacenar los resultados positivos encontrados
	data := make([]BoolData, 0, manyResults)

	var count int64
	// Iterar sobre el rango especificado
	for index := start; index < Space.CountLines(); index++ {

		// Si ya hemos encontrado suficientes resultados, salir del bucle
		if count > manyResults {
			break
		}

		buffer, err := Space.GetLine(col, index)
		if err != nil {
			return nil, err
		}

		if buffer[0] == dacBoolTrue[0] {
			data = append(data, BoolData{
				ID:    index,
				Value: true,
			})
			count++
		}
	}

	return data, nil
}

func (Space *Space) GetAllLinesRangesBoolJson(col int64, start int64, end int64) ([]byte, error) {

	// Obtener el array de BoolData
	data, err := Space.GetAllLinesRangesBool(col, start, end)
	if err != nil {
		return nil, err
	}

	// Convertir el slice a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (Space *Space) GetAllLinesRangesBoolTrueJson(col int64, start int64, manyResults int64) ([]byte, error) {

	// Obtener el array de BoolData
	data, err := Space.GetAllLinesRangesBoolTrue(col, start, manyResults)
	if err != nil {
		return nil, err
	}

	// Convertir el slice a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
