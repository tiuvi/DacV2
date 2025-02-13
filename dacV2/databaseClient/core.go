package databaseClient

import (
	dacv2 "dacV2"
)

func (spaceDB *SpaceDB) CountLines() (result int64, err error) {

	client, err := handleRequestCore(spaceDB, "CountLines")
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

	body, err := handlerResponse.ReadBodyBytes()
	if err != nil {
		return
	}

	return dacv2.BytesToInt64(body), nil
}
