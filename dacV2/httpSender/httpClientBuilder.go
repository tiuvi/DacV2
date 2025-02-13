package httpSender

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"encoding/base64"
)

type HttpClientBuilder struct {
	Client *http.Client
	Url string
	UrlValues url.Values
}


func validateUrl(baseURL string ,port uint16)error{

	if strings.HasPrefix(baseURL, "http://") {
    	return  errors.New("no se admiten conexiones inseguras")
    }

	if port == 80 {
		return  errors.New("el puerto 80 no admite conexiones seguras o https")
	}

	return nil
}

func NewBuildURL(baseURL string, port uint16, path string) (*HttpClientBuilder , error) {
	
	MG := HttpClientBuilder{}

	MG.Client = &http.Client{}
	MG.UrlValues = make(url.Values)

	// Validar la URL base
	err := validateUrl(baseURL, port)
	if err != nil {
		return &HttpClientBuilder{} , err
	}

	// Eliminar el prefijo "https://" si está presente
	baseURL = strings.TrimPrefix(baseURL, "https://")
	
	// Eliminar el prefijo "/" en la ruta
	path = strings.TrimPrefix(path, "/")
	

	// Construir la URL completa usando strings.Join
	MG.Url = strings.Join([]string{
		"https://",
		baseURL,
		":",
		strconv.Itoa(int(port)),
		"/",
		path,
	}, "")

	return &MG, nil
}

func (b *HttpClientBuilder) WriteUrlRaw(key, value string) {

	b.UrlValues.Set(key, value) // Agregar el par clave-valor
}

func (b *HttpClientBuilder) WriteUrlInt64(key string, value int64) {

	b.UrlValues.Set(key, strconv.FormatInt(value , 10) ) // Agregar el par clave-valor
}

func (b *HttpClientBuilder) WriteUrlMultipleRaw(key string, value []string) {

	for _ , moreValues := range value{
		b.UrlValues.Add(key, moreValues)
	}
}

func (SK *HttpClientBuilder) WriteUrlBase64(key string, message string) {

	SK.UrlValues.Set(key, base64.StdEncoding.EncodeToString([]byte(message)))
}
