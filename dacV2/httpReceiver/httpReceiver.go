package httpReceiver

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

var (
	ErrTokenEmpty    = errors.New("el campo token esta vacio")
	ErrNotTokenAuth2 = errors.New("no es un token auth2")
)

type HttpReceiver struct {
	head     http.Header
	url      url.Values
	response http.ResponseWriter
	request  *http.Request
}
 
func InitHttpReceiverCorsDisabled(response http.ResponseWriter, request *http.Request) *HttpReceiver {
 
	SK := new(HttpReceiver)
	SK.head = response.Header()
	SK.url = SK.request.URL.Query()
	SK.response = response
	SK.request = request

	//Permite solicidtudes de cualquier origen
	SK.WriteHeaderRaw("Access-Control-Allow-Origin", "*")

	//Se permiten todos los metodos de solicitud GET, POST, PUT, DELETE..
	SK.WriteHeaderRaw("Access-Control-Allow-Methods", "*")

	//Se permite el envio de cualquier tipo de headers desde el cliente al servidor.
	SK.WriteHeaderRaw("Access-Control-Allow-Headers", "*")

	//Se eligen qué cabeceras pueden ser leídas desde la respuesta por el código JavaScript del
	//cliente cuando realiza una solicitud de CORS
	SK.WriteHeaderRaw("Access-Control-Expose-Headers", "Content-Range , Bandwidth")

	//Determina cómo los recursos en tu servidor pueden ser compartidos con otras páginas web
	SK.WriteHeaderRaw("Cross-Origin-Resource-Policy", "cross-origin")

	return SK
}

func InitHttpReceiver(response http.ResponseWriter, request *http.Request) *HttpReceiver {

	SK := new(HttpReceiver)
	SK.head = response.Header()
	SK.url = request.URL.Query()
	SK.response = response
	SK.request = request

	return SK
}

func (SK *HttpReceiver) IsOptions() bool {

	return SK.request.Method == "OPTIONS"
}

func (SK *HttpReceiver) WriteCacheControlNoStore() {

	SK.response.Header().Set("Cache-Control", "no-store, max-age=0")
}

func (SK *HttpReceiver) WriteContentType(extension string) {

	if value, found := IsExtensionContent(extension); found {

		SK.response.Header().Set("Content-Type", value)

		//if SK.contenType == "glb" {}
		//SK.response.Header().Set("Access-Control-Allow-Origin", "*")

		if extension == "pdf" {
			SK.response.Header().Set("Content-Disposition", "inline")
		}
	}
}

func (SK *HttpReceiver) WriteContentLength(length int64) {

	SK.response.Header().Set("Content-Length", strconv.FormatInt(length, 10))
}

func (SK *HttpReceiver) WriteBandwidth(length int64) {

	SK.response.Header().Set("Bandwidth", strconv.FormatInt(length, 10))
}
