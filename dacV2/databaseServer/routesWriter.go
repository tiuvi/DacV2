package databaseServer

import (
	"dacV2"
	. "dacV2/httpReceiver"
	"net/http"
)

func init() {
 
	Routes["/CreateDirectory"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		dirPath := HR.ReadUrlMultiplesRaw("dirPath")

		err := dacv2.CreateDirectory(append([]string{globalPath}, dirPath...)...)
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

	Routes["/SetAt"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, err := handleRequestNoMaps(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		offSet, err := HR.ReadUrlInt64("offSet")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetAt(offSet, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/SetAtRange"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, err := handleRequestNoMaps(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		nRange, err := HR.ReadUrlInt64("nRange")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		bandwidth, err := HR.ReadUrlInt64("bandwidth")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetAtRange(body, nRange, bandwidth)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	//Añadir SetLineLimit y NewLineLimit si supera el tamaño maximo que de error
	//Esto es para las interfaces que ya no se validan en el cliente.
	Routes["/SetField"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, field, err := handleRequestField(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetField(field, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/SetFieldRaw"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, field, err := handleRequestField(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetFieldRaw(field, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/SetFieldRange"] = func(res http.ResponseWriter, req *http.Request) {

		//fieldsName linesName dirPath field
		HR, sf, field, err := handleRequestField(res, req)
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

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetFieldRange(field, body, rangue, bandwidth)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/SetLine"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, column, line, err := handleRequestLine(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetLine(column, line, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/SetLineRaw"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, column, line, err := handleRequestLine(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = sf.SetLineRaw(column, line, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteOk()
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/NewLine"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, column, err := handleRequestNewLine(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		line, err := sf.NewLine(column, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteInt64(line)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

	Routes["/NewLineRaw"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, column, err := handleRequestNewLine(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		line, err := sf.NewLineRaw(column, body)
		if err != nil {
			HR.ErrorStatusInternalServerError(err.Error())
			return
		}

		err = HR.WriteInt64(line)
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}
}
