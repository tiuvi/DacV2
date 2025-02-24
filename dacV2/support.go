package dacv2

import "os"



func FileExists(path string) (bool, error) {

    fileInfo, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }
        return false, err
    }
    
    return fileInfo.Mode().IsRegular(), nil
}


func (Space *Space) FileSize()(ret int64, err error){
   
    return Space.File.Seek(0, 2)
}

func (Space *Space) CalcRange(bandwidth int64) (nRange int64, err error) {

    fileSize , err := Space.File.Seek(0, 2)
    if err != nil {
    return
    }

    nRange = fileSize / bandwidth

    if fileSize % bandwidth != 0 {
        nRange = nRange + 1
    }
    
	return
}