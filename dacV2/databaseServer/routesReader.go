package databaseServer

import (
	"net/http"
)



func init() {

	Routes["/GetAt"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , err  := handleRequestCore(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		offSet , err := HR.ReadUrlInt64("offSet")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		sizeBuffer , err := HR.ReadUrlInt64("sizeBuffer")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetAt(offSet , sizeBuffer)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/GetAtRange"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , err  := handleRequestCore(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		nRange , err := HR.ReadUrlInt64("nRange")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		bandwidth , err := HR.ReadUrlInt64("bandwidth")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetAtRange(nRange , bandwidth)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/GetField"] = func(res http.ResponseWriter, req *http.Request) {


		HR , sf , field , err  := handleRequestField(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetField(field)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
	
	Routes["/GetFieldRaw"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , field , err  := handleRequestField(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetFieldRaw(field)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/GetFieldRange"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , field , err  := handleRequestField(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		rangue, err := HR.ReadUrlInt64("rangue")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		bandwidth, err := HR.ReadUrlInt64("bandwidth")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetFieldRange(field , rangue , bandwidth)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
	
	Routes["/GetLine"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , column , line , err  := handleRequestLine(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetLine(column , line)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
	
	Routes["/GetLineRaw"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , column , line , err  := handleRequestLine(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetLineRaw(column , line)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	
	Routes["/GetLinesRange"] = func(res http.ResponseWriter, req *http.Request) {

		HR , sf , err  := handleRequestCore(res , req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		startLine , err := HR.ReadUrlInt64("startLine")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		endLine , err := HR.ReadUrlInt64("endLine")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		data, err := sf.GetLinesRange(startLine , endLine )
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteBytes(data)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
}