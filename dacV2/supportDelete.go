package dacv2

import "os"

  
func DeleteFile(path string)error {

    return os.Remove(path)
}

func DeleteDirectory(path string)error {

    return os.RemoveAll(path)
}

func (Space *Space) DeleteDirectory()error {

    return os.RemoveAll(Space.Dir)
}

func (Space *Space) DeleteFile()error {

    return os.Remove(Space.FilePath)
}

func (Space *Space) TruncateFile(size int64)error {

    return Space.File.Truncate(size)
}

func (Space *Space) TruncateZeroFile()error {

    return Space.File.Truncate(0)
}

func (Space *Space) TruncateFileLine(line int64)error {

    return Space.File.Truncate(Space.SizeField+(line*Space.SizeLine))
}