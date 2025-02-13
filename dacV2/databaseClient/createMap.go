package databaseClient

import (
	dacv2 "dacV2"
	. "dacV2/httpSender"
	"errors"
	"strconv"
)

func CreateMap(baseURL string, port uint16, nameMap int64, spaceList []dacv2.SpaceList) (mapDac map[int64][3]int64, err error) {

	mapDac, sizeMap, err := dacv2.CreateMap(spaceList)
	if err != nil {
		return
	}

	client, err := NewBuildURL(baseURL, port, "createMap")
	if err != nil {
		return
	}

	client.WriteUrlRaw("nameMap", strconv.FormatInt(nameMap, 10))

	client.WriteUrlRaw("sizeMap", strconv.Itoa(int(sizeMap)))

	data, err := dacv2.TypesGolangToBytes(&mapDac)
	if err != nil {
		return
	}

	request, err := client.CreateGetRequest()
	if err != nil {
		return
	}

	requestBuild, err := request.CreatePostRequest(data)
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

	if body == "map exist" {
		return
	}

	if body == "ok" {
		return
	}

	return mapDac, errors.New(body)
}
