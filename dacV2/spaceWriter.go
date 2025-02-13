package dacv2


func (Space *Space) SetField(columnName int64, buf []byte) (err error) {

	size , err := Space.checkField(columnName)
	if err != nil{
		return
	}

	buf = SpacePadding(buf, size)

	_, err = Space.File.WriteAt(buf, size[0])
	if err != nil {
		return err
	}

	return
}

func (Space *Space) SetFieldRaw(columnName int64, buf []byte) (err error) {

	size , err := Space.checkField(columnName)
	if err != nil{
		return
	}

	if int64(len(buf)) != size[2]{
		return ErrLenBufferNoEqualSize
	}

	_, err = Space.File.WriteAt(buf, size[0])
	if err != nil {
		return err
	}

	return
}

func (Space *Space) SetFieldRange(columnName int64, buf []byte, rangue int64, bandwidth int64) (err error) {

	if(bandwidth == 0){
		return ErrRangeFieldNoZero
	}

	size , err := Space.checkField(columnName)
	if err != nil{
		return
	}

	var initRangue int64 = rangue * bandwidth
	var endRangue int64 = (rangue + 1) * bandwidth
	if endRangue > size[1] {
		endRangue = size[1]
	}

	if int64(len(buf)) != endRangue{
		return ErrLenBufferNoEqualSize
	}

	_, err = Space.File.WriteAt(buf, size[0] + initRangue)
	if err != nil {
		return err
	}

	return
}
 
func (Space *Space) SetLine(columnName int64, line int64, buf []byte) (err error) {

	size , err := Space.checkColumn(columnName)
	if err != nil{
		return
	}

	atomicCountLines := Space.AtomicCountLines.Load()
	if line >= atomicCountLines {
		Space.AtomicCountLines.Add( (line + 1) - atomicCountLines )
	}

	buf = SpacePadding(buf, size)

	_, err = Space.File.WriteAt(buf, Space.SizeField+(Space.SizeLine*line)+size[0])
	if err != nil {
		return err
	}

	return
}

func (Space *Space) SetLineRaw(columnName int64, line int64, buf []byte) (err error) {

	size , err := Space.checkColumn(columnName)
	if err != nil{
		return
	}

	atomicCountLines := Space.AtomicCountLines.Load()
	if line >= atomicCountLines {
		Space.AtomicCountLines.Add( (line + 1) - atomicCountLines )
	}

	if int64(len(buf)) != size[2]{
		return ErrLenBufferNoEqualSize
	}

	_, err = Space.File.WriteAt(buf, Space.SizeField+(Space.SizeLine*line)+size[0])
	if err != nil {
		return err
	}

	return
}

func (Space *Space) NewLine(columnName int64, buf []byte) (line int64 , err error) {

	size , err := Space.checkColumn(columnName)
	if err != nil{
		return
	}

	line = Space.AtomicCountLines.Add(1)
	currentLine := (line - 1)
	buf = SpacePadding(buf, size)

	_, err = Space.File.WriteAt(buf, Space.SizeField+(Space.SizeLine*currentLine) +size[0])
	if err != nil {
		return  0  , err
	}

	return currentLine ,nil
}

func (Space *Space) NewLineRaw(columnName int64, buf []byte) (line int64 , err error) {

	size , err := Space.checkColumn(columnName)
	if err != nil{
		return
	}

	line = Space.AtomicCountLines.Add(1)
	currentLine := (line - 1)

	if int64(len(buf)) != size[2]{
		return 0 , ErrLenBufferNoEqualSize
	}
	
	_, err = Space.File.WriteAt(buf, Space.SizeField+(Space.SizeLine*currentLine)+size[0])
	if err != nil {
		return 0 , err
	}

	return currentLine ,nil
}
