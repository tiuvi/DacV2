package databaseClient

import (
	. "dacV2/httpSender"
	"errors"
)

func (spaceDB *SpaceDB) DeleteFile() (err error) {

	client, err := NewBuildURL(spaceDB.BaseURL, spaceDB.Port, "DeleteFile")
	if err != nil {
		return
	}

	client.WriteUrlMultipleRaw("dirPath", spaceDB.DirPath)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyString()
	if err != nil {
		return
	}

	if body != "ok" {
		return errors.New(body)
	}

	return
}

func (spaceDB *SpaceDB) DeleteDirectory() (err error) {

	client, err := NewBuildURL(spaceDB.BaseURL, spaceDB.Port, "DeleteDirectory")
	if err != nil {
		return
	}

	var dir []string
	if len(spaceDB.DirPath) > 0 {
		dir = spaceDB.DirPath[:len(spaceDB.DirPath)-1]
	}

	client.WriteUrlMultipleRaw("dirPath", dir)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyString()
	if err != nil {
		return
	}

	if body != "ok" {
		return errors.New(body)
	}

	return
}

//Falta truncate y truncateZero
func (spaceDB *SpaceDB) TruncateZeroFile() (err error) {
	
	client, err := handleRequestCore(spaceDB, "TruncateFile")
	if err != nil {
		return
	}

	client.WriteUrlInt64("size" , 0)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyString()
	if err != nil {
		return
	}

	if body != "ok" {
		return errors.New(body)
	}

	return
}

func (spaceDB *SpaceDB) TruncateFile(size int64) (err error) {
	
	client, err := handleRequestCore(spaceDB, "TruncateFile")
	if err != nil {
		return
	}

	client.WriteUrlInt64("size" , size)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyString()
	if err != nil {
		return
	}

	if body != "ok" {
		return errors.New(body)
	}

	return
}


func (spaceDB *SpaceDB) TruncateFileLine(line int64) (err error) {
	
	client, err := handleRequestCore(spaceDB, "TruncateFileLine")
	if err != nil {
		return
	}

	client.WriteUrlInt64("line" , line)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyString()
	if err != nil {
		return
	}

	if body != "ok" {
		return errors.New(body)
	}

	return
}