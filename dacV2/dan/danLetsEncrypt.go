package dan

import (
	"net/http"
	"strconv"
	"time"
	"golang.org/x/crypto/acme/autocert"
)

func NewDanSpaceLetsEncrypt(dirCache string, email string , domain ...string) (dan *DanSpace, err error) {

	mux := http.NewServeMux()

	cert := &autocert.Manager{
		Cache:      autocert.DirCache(dirCache), 
		Prompt:     autocert.AcceptTOS,                   
		Email:      email,
		HostPolicy: autocert.HostWhitelist(domain...),
	}

	dan = &DanSpace{

		Mux: mux,
		Server: &http.Server{

			Addr: ":" + strconv.FormatUint(443, 10),

					//TimeoutHandler devuelve un Handler que ejecuta h con el límite de tiempo dado.
			Handler: http.TimeoutHandler(mux, 30*time.Second, "Timeout!\n"),

			//ReadTimeout es la duración máxima para leer la solicitud completa,
			ReadTimeout: 30 * time.Second,

			//ReadHeaderTimeout es la cantidad de tiempo permitido para leer los encabezados de solicitud.
			ReadHeaderTimeout: 30 * time.Second,

			//WriteTimeout es la duración máxima antes de que se agote el tiempo de escritura de la respuesta.
			WriteTimeout: 30 * time.Second,

			// IdleTimeout es la cantidad máxima de tiempo para esperar la próxima solicitud cuando se habilita la función Keep-Alives. Si IdleTimeout es cero, se utiliza el valor de ReadTimeout.
			IdleTimeout: 60 * time.Second,

			// MaxHeaderBytes controla la cantidad máxima de bytes que el servidor leerá al analizar las claves y valores del encabezado de la solicitud, incluida la línea de solicitud.
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
			
			TLSConfig:   cert.TLSConfig(),
		},
	}

	return dan , nil
}

func (dan *DanSpace) InitDanLetsEncrypt() {

	err := dan.Server.ListenAndServeTLS("", "")
	if err != nil {
		println(err.Error())
	}
}