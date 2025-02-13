package httpSender

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type HttpResponseHandler struct {
	*HttpRequestBuilder
	Response      *http.Response
	Status        string
	StatusCode    int
	ContentLength int64
}

func (sk *HttpResponseHandler) IsError(body []byte) error {

	switch sk.StatusCode {
	case http.StatusBadRequest:
		return errors.New(string(body))
	case http.StatusInternalServerError:
		return errors.New(string(body))
	case http.StatusNotFound:
		return errors.New(strings.TrimSpace(string(body)))
	default:
		return nil // No es un error
	}
}

func (MG *HttpRequestBuilder) StartSender() (handler *HttpResponseHandler, err error) {

	resp, err := MG.Client.Do(MG.Request)
	if err != nil {
		return 
	}

	return &HttpResponseHandler{
		HttpRequestBuilder: MG,         // Asignamos HttpRequestBuilder de MG
		Response:    resp,       // Asignamos la respuesta HTTP
		Status:      resp.Status,
		StatusCode:  resp.StatusCode,
		ContentLength: resp.ContentLength,
	}, nil
}

func (MG *HttpResponseHandler) ReadHeaderRaw(key string) string {

	return MG.Response.Header.Get(key)

}

func (MG *HttpResponseHandler) ReadHeaderBase64(key string) (keyHeader string, err error) {

	keyHeaderbytes, err := base64.StdEncoding.DecodeString(MG.ReadHeaderRaw(key))
	if err != nil {
		return
	}

	keyHeader = string(keyHeaderbytes)
	return
}

func (SK *HttpResponseHandler) ReadBodyBytes() (body []byte, err error) {

	body, errRead := io.ReadAll(SK.Response.Body)

	errClose := SK.Response.Body.Close()

	if errRead != nil {
		return
	}

	if errClose != nil {
		return
	}

	err = SK.IsError(body)
	if err != nil {
		return
	}

	return
}

func (SK *HttpResponseHandler) ReadBodyString() (string, error) {

	body , err := SK.ReadBodyBytes()
	if err != nil {
		return "" , err
	}

	return string(body) , nil
}

func (SK *HttpResponseHandler) ReadBodyInt64() (int64, error) {

	body , err := SK.ReadBodyBytes()
	if err != nil {
		return 0 , err
	}

	return strconv.ParseInt(string(body) , 10 , 64)
}

func (SK *HttpResponseHandler) ReadBodyMaxBytes(maxSize int64) (body []byte, err error) {

	limitBody := &io.LimitedReader{R: SK.Response.Body, N: maxSize}

	body, errRead := io.ReadAll(limitBody)

	errBody := SK.Response.Body.Close()

	if errRead != nil {
		return
	}

	if errBody != nil {
		return
	}
	
	err = SK.IsError(body)
	if err != nil {
		return
	}

	return
}
