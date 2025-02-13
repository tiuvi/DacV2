package databaseClient

var dacBoolTrue []byte = []byte("t")
var dacBoolFalse []byte = []byte("f")

// tamaño: 1byte
// Por defecto false
func (Space *SpaceDB) GetFieldBool(col int64) (bool, error) {

	buffer, err := Space.GetField(col)
	if err != nil {
		return false, err
	}

	value := string(buffer)
	if len(value) == 0 {
		return false, nil
	}

	if value == "f" {
		return false, nil
	}
	if value == "t" {
		return true, nil
	}

	return true, nil
}

// tamaño: 1byte
// Por defecto false
func (Space *SpaceDB) SetFieldBool(col int64, active bool) error {

	if active {
		return Space.SetField(col, dacBoolTrue)
	}

	return Space.SetField(col, []byte("f"))
}
