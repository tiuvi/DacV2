package databaseServer

import (
	"dacV2"
	. "dacV2/httpReceiver"
	"path/filepath"
	"net/http"
)
 

func init() {

	Routes["/DeleteFile"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		dirPath := HR.ReadUrlMultiplesRaw("dirPath")

		err := dacv2.DeleteFile(filepath.Join(append([]string{globalPath}, dirPath...)...))
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/DeleteDirectory"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		dirPath := HR.ReadUrlMultiplesRaw("dirPath")
	
		err := dacv2.DeleteDirectory(filepath.Join(append([]string{globalPath}, dirPath...)...))
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
 
	Routes["/TruncateFile"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		HR, sf, err := handleRequestCore(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		size, err := HR.ReadUrlInt64("size")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}
	
		err = sf.TruncateFile(size)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
	
	Routes["/TruncateFileLine"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		HR, sf, err := handleRequestCore(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		line, err := HR.ReadUrlInt64("line")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = sf.TruncateFileLine(line)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

}
