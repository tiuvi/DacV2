package databaseClient

import (
	dacv2 "dacV2"
)

func (spaceDB *SpaceDB) AtomicFieldSumInt64(field int64, increment int64) (result int64, err error) {

	client, err := handleRequestField(spaceDB, field, "AtomicFieldSumInt64")
	if err != nil {
		return
	}

	client.WriteUrlInt64("increment", increment)

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
