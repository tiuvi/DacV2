package databaseServer

import (
	. "dacV2"
	. "dacV2/httpReceiver"
	"net/http"
)

func init() {

	Routes["/createMap"] = func(res http.ResponseWriter, req *http.Request) {

		HR := InitHttpReceiver(res, req)

		body, err := HR.ReadBodyBytes()
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		outRef := make(map[int64][3]int64)

		err = BytesToTypesGolang(body, &outRef)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		//nameMap sizeMap
		nameMap, err := HR.ReadUrlInt64("nameMap")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		if nameMap == 0 {
			HR.ErrorStatusBadRequest("El valor 0 no esta disponible")
			return
		}

		sizeMap, err := HR.ReadUrlInt64("sizeMap")
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		fileMaps.Mu.Lock()

		_, found := mapsDac[nameMap]
		if found {
			fileMaps.Mu.Unlock()
			HR.WriteBytes([]byte("map exist"))
			return
		}

		mapsDac[nameMap] = mapDacItem{
			Map:     outRef,
			sizeMap: sizeMap,
		}

		fileMaps.Mu.Unlock()

		line, err := fileMaps.NewLineInt64(colNameMap, nameMap)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = fileMaps.SetLineInt64(colSizeMap, line, sizeMap)
		if err != nil {
			HR.ErrorStatusBadRequest(err.Error())
			return
		}

		err = fileMaps.SetLineEncoder(colMap, line, outRef)
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