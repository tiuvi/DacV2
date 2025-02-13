package httpSender

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"time"
)


type HttpRequestBuilder struct {
	*HttpClientBuilder
	Request *http.Request
}

func (httpClientBuilder *HttpClientBuilder)CreateHeadRequest()(MG *HttpRequestBuilder , err error) {

	req, err := http.NewRequest("HEAD",  httpClientBuilder.Url , nil)
	if err != nil {
		return
	}

	return &HttpRequestBuilder{
		HttpClientBuilder: httpClientBuilder,
		Request:           req,
	}, nil
}

func (httpClientBuilder *HttpClientBuilder)CreateGetRequest()(MG *HttpRequestBuilder , err error) {

	req, err := http.NewRequest("GET",  httpClientBuilder.Url + "?" + httpClientBuilder.UrlValues.Encode(), nil)
	if err != nil {
		return
	}

	return &HttpRequestBuilder{
		HttpClientBuilder: httpClientBuilder,
		Request:           req,
	}, nil
}

func (httpClientBuilder *HttpClientBuilder)CreateGeTRequestTimeOut(Timeout time.Duration)(MG *HttpRequestBuilder , err error) {

	httpClientBuilder.Client.Timeout = Timeout

	req, err := http.NewRequest("GET",  httpClientBuilder.Url + "?" + httpClientBuilder.UrlValues.Encode() , nil)
	if err != nil {
		return
	}

	return &HttpRequestBuilder{
		HttpClientBuilder: httpClientBuilder,
		Request:           req,
	}, nil
}

func (httpClientBuilder *HttpClientBuilder)CreatePostRequest(data []byte)(MG *HttpRequestBuilder , err error) {

	req, err := http.NewRequest("POST",  httpClientBuilder.Url + "?" + httpClientBuilder.UrlValues.Encode(), bytes.NewReader(data))
	if err != nil {
		return
	}

	return &HttpRequestBuilder{
		HttpClientBuilder: httpClientBuilder,
		Request:           req,
	}, nil
}

func (httpClientBuilder *HttpClientBuilder)CreatePostRequestTimeOut(data []byte , Timeout time.Duration)(MG *HttpRequestBuilder , err error) {

	httpClientBuilder.Client.Timeout = Timeout

	req, err := http.NewRequest("POST",  httpClientBuilder.Url + "?" + httpClientBuilder.UrlValues.Encode() , bytes.NewReader(data))
	if err != nil {
		return
	}

	return &HttpRequestBuilder{
		HttpClientBuilder: httpClientBuilder,
		Request:           req,
	}, nil
}

func (SK *HttpRequestBuilder) WriteHeaderBase64(key string, message string) {

	SK.Request.Header.Set(key, base64.StdEncoding.EncodeToString([]byte(message)))
}

func (SK *HttpRequestBuilder) WriteHeaderRaw(key string, message string) {

	SK.Request.Header.Set(key, message)
}