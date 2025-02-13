package dacv2



func NewSpaceContent(sizeContent int64 , dirPath ...string) (newSpace *Space, err error) {

	file, filePath, extension, dir, err := OpenFileOrCreate(dirPath...)
	if err != nil {
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return nil, err
	}

	buf := Int64ToBytes(sizeContent)
	
	_, err = file.WriteAt(buf, 0)
	if err != nil {
		return
	}

	mapDac , sizeMap , err := CreateMap([]SpaceList{
		{FieldSizeFile, 8},
		{FieldContentFile, sizeContent},
	})
	if err != nil {
		return
	}

	newSpace = &Space{

		Dir:       dir,
		FilePath:  filePath,
		Extension: extension,

		File: file,

		IndexSizeFields: mapDac,
		SizeField:       sizeMap,
	}

	return
}

func OpenSpaceContent(dirPath ...string) (newSpace *Space, err error) {

	file, filePath, extension, dir, err := OpenFileIfExist(dirPath...)
	if err != nil {
		return
	}

	buf := make([]byte , 8 )
	_, err = file.ReadAt(buf, 0)
	if err != nil {
		return
	}

	sizeContent := BytesToInt64(buf)

	mapDac , sizeMap , err := CreateMap([]SpaceList{
		{FieldSizeFile, 8},
		{FieldContentFile, sizeContent},
	})
	if err != nil {
		return
	}

	newSpace = &Space{

		Dir:       dir,
		FilePath:  filePath,
		Extension: extension,

		File: file,

		IndexSizeFields: mapDac,
		SizeField:       sizeMap,
	}

	return
}