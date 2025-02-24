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

    Space.AtomicCountLines.Store(0)

    return Space.File.Truncate(size)
}

func (Space *Space) TruncateZeroFile()error {

    return Space.TruncateFile(0)
}

func (Space *Space) TruncateFileLine(line int64)error {

    return Space.TruncateFile(Space.SizeField+(line*Space.SizeLine))
}