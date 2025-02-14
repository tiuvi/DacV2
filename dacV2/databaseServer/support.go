package databaseServer

import (
	"dacV2"
	. "dacV2/httpReceiver"
	"path/filepath"

	"net/http"
)
 
var dacBoolTrue []byte = []byte("t")
var dacBoolFalse []byte = []byte("f")


func init() {

	Routes["/FileExists"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		path, err := HR.ReadUrlRawReq("path")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		exist , err := dacv2.FileExists(filepath.Join(globalPath , path))
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		var data []byte
		if exist {
			data = dacBoolTrue
		}else{
			data = dacBoolFalse
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}

	}

}
