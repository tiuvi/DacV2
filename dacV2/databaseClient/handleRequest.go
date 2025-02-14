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
