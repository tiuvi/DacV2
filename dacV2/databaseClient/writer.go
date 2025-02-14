package databaseClient

import (
	"errors"
)

func (spaceDB *SpaceDB) SetAt(offSet int64, data []byte) (err error) {

	client, err := handleRequestCore(spaceDB, "SetAt")
	if err != nil {
		return
	}

	client.WriteUrlInt64("offSet" , offSet)

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) SetAtRange(data []byte, nRange int64, bandwidth int64) (err error) {

	client, err := handleRequestCore(spaceDB, "SetAtRange")
	if err != nil {
		return
	}

	client.WriteUrlInt64("nRange" , nRange)

	client.WriteUrlInt64("bandwidth" , bandwidth)

	requestBuild, err := client.CreatePostRequest(data)
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

// fieldsName linesName dirPath field
func (spaceDB *SpaceDB) SetField(field int64, data []byte) (err error) {

	client, err := handleRequestField(spaceDB, field, "SetField")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) SetFieldRaw(field int64, data []byte) (err error) {

	client, err := handleRequestField(spaceDB, field, "SetFieldRaw")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) SetFieldRange(field int64, data []byte, rangue int64, bandwidth int64) (err error) {

	client, err := handleRequestField(spaceDB, field, "SetFieldRange")
	if err != nil {
		return
	}

	client.WriteUrlInt64("rangue", rangue)

	client.WriteUrlInt64("bandwidth", bandwidth)

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) SetLine(column int64, line int64, data []byte) (err error) {

	client, err := handleRequestLine(spaceDB, column, line, "SetLine")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) SetLineRaw(column int64, line int64, data []byte) (err error) {

	client, err := handleRequestLine(spaceDB, column, line, "SetLineRaw")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
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

func (spaceDB *SpaceDB) NewLine(column int64, data []byte) (line int64, err error) {

	client, err := handleRequestNewLine(spaceDB, column, "NewLine")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyInt64()
}

func (spaceDB *SpaceDB) NewLineRaw(column int64, data []byte) (line int64, err error) {

	client, err := handleRequestNewLine(spaceDB, column, "NewLineRaw")
	if err != nil {
		return
	}

	requestBuild, err := client.CreatePostRequest(data)
	if err != nil {
		return
	}

	handlerResponse, err := requestBuild.StartSender()
	if err != nil {
		return
	}

	return handlerResponse.ReadBodyInt64()
}
