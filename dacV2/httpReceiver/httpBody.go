package httpReceiver

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)
 

func (SK *HttpReceiver) WriteOk() (err error) {

	err = SK.Write("txt", []byte("ok"))
	if err != nil {
		return
	}
	return
}

func (SK *HttpReceiver) Write(extension string, content []byte) error {

	extensionOk, found := IsExtensionContent(extension)
	if !found {
		return errors.New("extension de archivo no valida")
	}

	SK.WriteHeaderRaw("Content-Type", extensionOk)
	SK.WriteHeaderRaw("X-Content-Type-Options", "nosniff")

	_, err := SK.response.Write(content)
	if err != nil {
		return err
	}

	return nil
}
func (SK *HttpReceiver) WriteBytes(content []byte) error {

	_, err := SK.response.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func (SK *HttpReceiver) WriteTxtString(content string) error {

	SK.WriteHeaderRaw("Content-Type", "text/plain; charset=UTF-8")

	_, err := SK.response.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func (SK *HttpReceiver) WriteString(extension string, content string) error {

	return SK.Write(extension , []byte(content))
}

func (SK *HttpReceiver) WriteBool(boolean bool) error {

	content := strconv.FormatBool(boolean)

	return SK.WriteTxtString(content)
}

func (SK *HttpReceiver) WriteInt64(value int64) error {

	content := strconv.FormatInt(value, 10)

	return SK.WriteTxtString(content)
}

func (SK *HttpReceiver) WriteJson(jsonValue interface{}) error {

	jsonByte, err := json.Marshal(jsonValue)
	if err != nil {
		SK.ErrorStatusInternalServerError(err.Error())
		return nil
	}

	return SK.Write("json", jsonByte)
}

func (SK *HttpReceiver) WriteJsonBytes(content []byte) error {

	SK.WriteHeaderRaw("Content-Type", "application/json; charset=utf-8")

	_, err := SK.response.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func (SK *HttpReceiver) ReadBodyBytes() (body []byte, err error) {

	body, errRead := io.ReadAll(SK.request.Body)

	errClose := SK.request.Body.Close()

	if errRead != nil {
		return
	}

	if errClose != nil {
		return
	}
	
	return
}

func (SK *HttpReceiver) ReadBodyMaxBytes(maxSize int64) (body []byte, err error) {

	limitBody := http.MaxBytesReader(SK.response, SK.request.Body, maxSize)

	body, errRead := io.ReadAll(limitBody)

	errLimitBody := limitBody.Close()

	errBody := SK.request.Body.Close()

	if errRead != nil {
		return
	}

	if errLimitBody != nil {
		return
	}

	if errBody != nil {
		return
	}

	return
}
