package dacv2



func NewSpaceContent(dirPath ...string) (newSpace *Space, err error) {

	file, filePath, extension, dir, err := OpenFileOrCreate(dirPath...)
	if err != nil {
		return
	}

	newSpace = &Space{

		Dir:       dir,
		FilePath:  filePath,
		Extension: extension,

		File: file,
	}

	return
}
