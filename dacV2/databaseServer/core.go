package databaseServer

import (
	dacv2 "dacV2"
	"net/http"
)



func init(){

	Routes["/CountLines"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, err := handleRequestCore(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data := sf.CountLines()
	
		err = HR.WriteBytes(dacv2.Int64ToBytes(data))
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

}