package databaseServer

import (
	"dacV2"
	. "dacV2/httpReceiver"
	"errors"
	"net/http"
)

func handleRequestNoMaps(res http.ResponseWriter, req *http.Request, )(HR *HttpReceiver, sf *dacv2.Space , err error){

	HR = InitHttpReceiver(res, req)

	dirPath := HR.ReadUrlMultiplesRaw("dirPath")

	sf, err = globalCache.OpenSpaceRange(append([]string{globalPath}, dirPath...)... )
	if err != nil {
		HR.ErrorStatusInternalServerError(err.Error())
		return
	}
	
	return
}

func handleRequestCore(res http.ResponseWriter, req *http.Request, )(HR *HttpReceiver, sf *dacv2.Space , err error){

	HR = InitHttpReceiver(res, req)

	fieldsName , _ := HR.ReadUrlInt64("fieldsMapName")

	linesName , _ := HR.ReadUrlInt64("linesMapName")

	dirPath := HR.ReadUrlMultiplesRaw("dirPath")

	fileMaps.Mu.RLock()

	mapFields := mapsDac[fieldsName]

	mapLines := mapsDac[linesName]

	fileMaps.Mu.RUnlock()

	if mapFields.Map == nil && mapLines.Map == nil{
		err = errors.New("no se seleccionaron mapas de campos o lineas")
		return
	}

	sf, err = globalCache.Open(mapFields.Map, mapFields.sizeMap, mapLines.Map, mapLines.sizeMap, append([]string{globalPath}, dirPath...)... )
	if err != nil {
		HR.ErrorStatusInternalServerError(err.Error())
		return
	}
	
	return
}

func handleRequestField(res http.ResponseWriter, req *http.Request, )(HR *HttpReceiver, sf *dacv2.Space , field int64 ,err error){

	HR = InitHttpReceiver(res, req)

	field, err = HR.ReadUrlInt64("field")
	if err != nil {
		return 
	}

	fieldsName , _ := HR.ReadUrlInt64("fieldsMapName")

	linesName , _ := HR.ReadUrlInt64("linesMapName")

	dirPath := HR.ReadUrlMultiplesRaw("dirPath")

	fileMaps.Mu.RLock()

	mapFields := mapsDac[fieldsName]

	mapLines := mapsDac[linesName]

	fileMaps.Mu.RUnlock()

	if mapFields.Map == nil && mapLines.Map == nil{
		err = errors.New("no se seleccionaron mapas de campos o lineas")
		return
	}

	sf, err = globalCache.Open(mapFields.Map, mapFields.sizeMap, mapLines.Map, mapLines.sizeMap, append([]string{globalPath}, dirPath...)... )
	if err != nil {
		HR.ErrorStatusInternalServerError(err.Error())
		return
	}
	
	return
}

func handleRequestLine(res http.ResponseWriter, req *http.Request, )(HR *HttpReceiver, sf *dacv2.Space , column int64 ,line int64,err error){

	HR = InitHttpReceiver(res, req)


	column, err = HR.ReadUrlInt64("column")
	if err != nil {
		HR.ErrorStatusBadRequest(err.Error())
		return
	}

	line , err = HR.ReadUrlInt64("line")
	if err != nil {
		HR.ErrorStatusBadRequest(err.Error())
		return
	}

	fieldsName , _ := HR.ReadUrlInt64("fieldsMapName")

	linesName , _ := HR.ReadUrlInt64("linesMapName")

	dirPath := HR.ReadUrlMultiplesRaw("dirPath")

	fileMaps.Mu.RLock()

	mapFields := mapsDac[fieldsName]

	mapLines := mapsDac[linesName]

	fileMaps.Mu.RUnlock()

	if mapFields.Map == nil && mapLines.Map == nil{
		err = errors.New("no se seleccionaron mapas de campos o lineas")
		return
	}

	sf, err = globalCache.Open(mapFields.Map, mapFields.sizeMap, mapLines.Map, mapLines.sizeMap, append([]string{globalPath}, dirPath...)... )
	if err != nil {
		HR.ErrorStatusInternalServerError(err.Error())
		return
	}
	
	return
}

func handleRequestNewLine(res http.ResponseWriter, req *http.Request, )(HR *HttpReceiver, sf *dacv2.Space , column int64 ,err error){

	HR = InitHttpReceiver(res, req)

	column, err = HR.ReadUrlInt64("column")
	if err != nil {
		HR.ErrorStatusBadRequest(err.Error())
		return
	}

	fieldsName , _ := HR.ReadUrlInt64("fieldsMapName")

	linesName , _ := HR.ReadUrlInt64("linesMapName")

	dirPath := HR.ReadUrlMultiplesRaw("dirPath")

	fileMaps.Mu.RLock()

	mapFields := mapsDac[fieldsName]

	mapLines := mapsDac[linesName]

	fileMaps.Mu.RUnlock()

	if mapFields.Map == nil && mapLines.Map == nil{
		err = errors.New("no se seleccionaron mapas de campos o lineas")
		return
	}

	sf, err = globalCache.Open(mapFields.Map, mapFields.sizeMap, mapLines.Map, mapLines.sizeMap, append([]string{globalPath}, dirPath...)... )
	if err != nil {
		HR.ErrorStatusInternalServerError(err.Error())
		return
	}
	
	return
}