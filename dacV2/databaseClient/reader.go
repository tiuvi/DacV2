package databaseClient


func (spaceDB *SpaceDB) GetAt(offSet int64 ,sizeBuffer int64) (body []byte, err error) {

	client, err := handleRequestCore(spaceDB, "GetAt")
	if err != nil {
		return
	}

	client.WriteUrlInt64("offSet" , offSet)

	client.WriteUrlInt64("sizeBuffer" , sizeBuffer)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetAtRange(nRange int64 ,bandwidth int64) (body []byte, err error) {

	client, err := handleRequestCore(spaceDB, "GetAtRange")
	if err != nil {
		return
	}

	client.WriteUrlInt64("nRange" , nRange)

	client.WriteUrlInt64("bandwidth" , bandwidth)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetField(field int64) (body []byte, err error) {

	client, err := handleRequestField(spaceDB, field, "GetField")
	if err != nil {
		return
	}

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetFieldRaw(field int64) (body []byte, err error) {

	client, err := handleRequestField(spaceDB, field, "GetFieldRaw")
	if err != nil {
		return
	}

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetFieldRange(field int64, rangue int64, bandwidth int64) (body []byte, err error) {

	client, err := handleRequestField(spaceDB, field, "GetFieldRange")
	if err != nil {
		return
	}

	client.WriteUrlInt64("rangue", rangue)

	client.WriteUrlInt64("bandwidth", bandwidth)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetLine(column int64, line int64) (body []byte, err error) {

	client, err := handleRequestLine(spaceDB, column, line, "GetLine")
	if err != nil {
		return
	}

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetLineRaw(column int64, line int64) (body []byte, err error) {

	client, err := handleRequestLine(spaceDB, column, line, "GetLineRaw")
	if err != nil {
		return
	}

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}

func (spaceDB *SpaceDB) GetLinesRange(startLine int64, endLine int64) (body []byte, err error) {

	client, err := handleRequestCore(spaceDB , "GetLinesRange")
	if err != nil {
		return
	}

	client.WriteUrlInt64("startLine", startLine)

	client.WriteUrlInt64("endLine", endLine)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyBytes()
}
