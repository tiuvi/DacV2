package databaseClient

import (
	. "dacV2/httpSender"
	"strconv"
)

func handleRequestCore(spaceDB *SpaceDB, endpoint string) (client *HttpClientBuilder, err error) {

	client, err = NewBuildURL(spaceDB.BaseURL, spaceDB.Port, endpoint)
	if err != nil {
		return
	}

	client.WriteUrlInt64("fieldsMapName", spaceDB.FieldsMapName)

	client.WriteUrlInt64("linesMapName", spaceDB.LinesMapName)

	client.WriteUrlMultipleRaw("dirPath", spaceDB.DirPath)

	return
}

func handleRequestField(spaceDB *SpaceDB, field int64, endpoint string) (client *HttpClientBuilder, err error) {

	client, err = NewBuildURL(spaceDB.BaseURL, spaceDB.Port, endpoint)
	if err != nil {
		return
	}

	client.WriteUrlRaw("field", strconv.FormatInt(field, 10))

	client.WriteUrlInt64("fieldsMapName", spaceDB.FieldsMapName)

	client.WriteUrlInt64("linesMapName", spaceDB.LinesMapName)

	client.WriteUrlMultipleRaw("dirPath", spaceDB.DirPath)

	return
}

func handleRequestLine(spaceDB *SpaceDB, column int64, line int64, endpoint string) (client *HttpClientBuilder, err error) {

	client, err = NewBuildURL(spaceDB.BaseURL, spaceDB.Port, endpoint)
	if err != nil {
		return
	}

	client.WriteUrlRaw("column", strconv.FormatInt(column, 10))

	client.WriteUrlInt64("line", line)

	client.WriteUrlInt64("fieldsMapName", spaceDB.FieldsMapName)

	client.WriteUrlInt64("linesMapName", spaceDB.LinesMapName)

	client.WriteUrlMultipleRaw("dirPath", spaceDB.DirPath)

	return client, nil
}

func handleRequestNewLine(spaceDB *SpaceDB, column int64, endpoint string) (client *HttpClientBuilder, err error) {

	client, err = NewBuildURL(spaceDB.BaseURL, spaceDB.Port, endpoint)
	if err != nil {
		return
	}

	client.WriteUrlRaw("column", strconv.FormatInt(column, 10))

	client.WriteUrlInt64("fieldsMapName", spaceDB.FieldsMapName)

	client.WriteUrlInt64("linesMapName", spaceDB.LinesMapName)

	client.WriteUrlMultipleRaw("dirPath", spaceDB.DirPath)

	return
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
