package httpReceiver

import (
	"encoding/base64"
	"errors"
	"strings"
)

func (SK *HttpReceiver) WriteHeaderCode(code int) {

	SK.response.WriteHeader(code)
}

func (SK *HttpReceiver) ReadHeaderRaw(key string) string {

	return SK.request.Header.Get(key)
}

func (SK *HttpReceiver) ReadHeaderMultiplesRaw(key string) []string {


	return SK.request.Header.Values(key)
}


func (SK *HttpReceiver) ReadHeaderRawRequired(key string) (string, error) {

	value := SK.ReadHeaderRaw(key)
	if value == "" {

		return "", errors.New(strings.Join([]string{"el campo ", key, " esta vacio"}, ""))
	}
	return value, nil
}

func (SK *HttpReceiver) WriteHeaderRaw(key string, Header string) {

	SK.head.Set(key, Header)
}

func (SK *HttpReceiver) ReadHeaderBase64(key string) (keyHeader string, err error) {

	keyHeaderbytes, err := base64.StdEncoding.DecodeString(SK.ReadHeaderRaw(key))
	if err != nil {
		return
	}
	keyHeader = string(keyHeaderbytes)
	return
}

func (SK *HttpReceiver) WriteHeaderBase64(key string, Header string) {

	SK.head.Set(key, base64.StdEncoding.EncodeToString([]byte(Header)))
}
