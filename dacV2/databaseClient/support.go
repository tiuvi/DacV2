package databaseClient

import (
	. "dacV2/httpSender"
)


func (spaceDB *SpaceDB) FileExists(path string) (exist bool, err error) {

	client, err := NewBuildURL(spaceDB.BaseURL, spaceDB.Port, "FileExists")
	if err != nil {
		return
	}

	client.WriteUrlRaw("path", path)

	requestBuild, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	body, err := handlerResponse.ReadBodyBytes()
	if err != nil {
		return
	}

	return body[0] == dacBoolTrue[0], nil
}
