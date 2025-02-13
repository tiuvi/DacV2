package dacv2



var dacBoolTrue []byte = []byte("t")
var dacBoolFalse []byte = []byte("f")

func (Space *Space) NewLineBool(col int64, active bool) (int64, error) {

	if active {
		return Space.NewLine(col, dacBoolTrue)
	}

	return Space.NewLine(col, dacBoolFalse)
}


//Necesita un solo byte para guardar el valor
func (Space *Space) SetLineBool(col int64, line int64, active bool) error {
	if active {
		return Space.SetLine(col, line, dacBoolTrue)
	}

	return Space.SetLine(col, line, dacBoolFalse)
}


func (Space *Space) GetLineBool(col int64, line int64) (bool, error) {

	buffer, err := Space.GetLine(col, line)
	if err != nil {
		return false, err
	}


	if len(buffer) == 0 {
		return false, nil
	}

	if buffer[0] == dacBoolFalse[0] {
		return false, nil
	}
	
	if  buffer[0] == dacBoolTrue[0] {
		return true, nil
	}

	return true, nil
}
