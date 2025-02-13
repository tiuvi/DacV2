package databaseClient

import (
	dacv2 "dacV2"
)

func (spaceDB *SpaceDB) AtomicLineSumInt64(col int64, line int64, increment int64) (result int64, err error) {

	client, err := handleRequestLine(spaceDB, col, line, "AtomicLineSumInt64")
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
