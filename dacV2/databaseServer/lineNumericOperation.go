package databaseServer

import (
	"dacV2"
	"net/http"
)


	
func init() {
 
	Routes["/AtomicLineSumInt64"] = func(res http.ResponseWriter, req *http.Request) {

		HR, sf, column, line , err := handleRequestLine(res, req)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		increment , err := HR.ReadUrlInt64("increment")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		sf.Mu.Lock()
		defer sf.Mu.Unlock()

		data, err := sf.GetLineInt64(column , line)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		result := data + increment

		err = sf.SetLineInt64(column,line , result)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = HR.WriteBytes(dacv2.Int64ToBytes(result))
		if err != nil {
			println("databaseServer - WriteOk: ", err.Error())
			return
		}
	}

}
