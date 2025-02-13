package httpReceiver

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

func (SK *HttpReceiver) ExistUrl(key string) bool {

	return SK.url.Has(key)
}

func (SK *HttpReceiver) ReadUrlRaw(key string) string {

	return SK.url.Get(key)
}

func (SK *HttpReceiver) ReadUrlMultiplesRaw(key string) []string {

	return SK.request.URL.Query()[key]
}

func (SK *HttpReceiver) ReadUrlRawReq(key string) (data string, err error) {

	data = SK.url.Get(key)
	if data == "" {
		return "", errors.New(strings.Join([]string{"el campo ", key, " esta vacio"}, ""))
	}
	return
}

func (SK *HttpReceiver) ReadUrlRawLimit(key string, lowLimit int64, highLimit int64) (data string, err error) {

	data = SK.url.Get(key)

	dataLength := int64(len(data))

	if dataLength < lowLimit {
		err = errors.New(strings.Join([]string{"el campo ", key, " está vacío o tiene menos de ", strconv.FormatInt(lowLimit, 10), " caracteres"}, ""))
		return
	}

	if dataLength > highLimit {
		err = errors.New(strings.Join([]string{"el campo ", key, " excede el límite de ", strconv.FormatInt(highLimit, 10), " caracteres"}, ""))
		return
	}

	return
}

func (SK *HttpReceiver) ReadUrlInt64(key string) (data int64, err error) {

	number := SK.url.Get(key)

	data, err = strconv.ParseInt(number, 10, 64)
	if err != nil {
		return
	}

	return
}

func (SK *HttpReceiver) ReadUrlBool(key string) (data bool, err error) {

	// Obtén el valor de la URL correspondiente a la clave
	value := SK.url.Get(key)

	// Convierte el valor a booleano
	data, err = strconv.ParseBool(value)
	if err != nil {
		return
	}

	return
}

func (SK *HttpReceiver) ReadUrlBase64(key string) (string, error) {

	keyHeaderbytes, err := base64.StdEncoding.DecodeString(SK.url.Get(key))
	if err != nil {
		return "", err
	}

	keyHeader := string(keyHeaderbytes)

	return keyHeader, nil
}
