package dacv2

func SpacePadding(buf []byte, size [3]int64) []byte {

	bytesCount := int64(len(buf))
	sizeTotal := size[2]

	if bytesCount < sizeTotal {
		buf = append(buf, make([]byte, sizeTotal-bytesCount)...)
	}
 
	if bytesCount > sizeTotal {

		buf = (buf)[:sizeTotal]
	}

	return buf
}

func SpaceTrimNull(buf []byte)([]byte) {

	//Limpiamos nulos
	for len(buf) > 0 && buf[len(buf)-1] == 0 {

		buf = (buf)[:len(buf)-1]
	}

	return buf
}

func (Space *Space)checkField(columnName int64)(size [3]int64 , err error){

	if Space.IndexSizeFields == nil {
		return [3]int64{} , ErrNoFieldsOrColumnMap
	}

	size, found := Space.IndexSizeFields[columnName]
	if !found {
		return  [3]int64{} , ErrNoFieldsOrColumn
	}

	return
}

func (Space *Space)checkColumn(columnName int64)(size [3]int64 , err error){

	
	if Space.IndexSizeColumns == nil {
		return [3]int64{} , ErrNoFieldsOrColumnMap
	}
	
	size, found := Space.IndexSizeColumns[columnName]
	if !found {
		return  [3]int64{} , ErrNoFieldsOrColumn
	}
	
	return
}

//El numero de lineas siempre va ser uno mas , ya que las lineas empiezan en 0.
//Una linea significa que la linea 0 esta escrita 2 lineas que la linea 0 y 1 estan escritas.
func (Space *Space)CountLines()(int64){

	return Space.AtomicCountLines.Load()
}