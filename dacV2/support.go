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


