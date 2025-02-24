package dan

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DanSpace struct {
	Mux    *http.ServeMux
	Server *http.Server
}

func NewDanSpace(port uint16) *DanSpace {

	mux := http.NewServeMux()
	
	dan:=  &DanSpace{

		Mux: mux,
		Server: &http.Server{

			Addr: ":" + strconv.FormatUint(uint64(port), 10),

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
			
			// TLSNextProto especifica opcionalmente una función para asumir el control propiedad de la
			// conexión TLS proporcionada cuando un ALPN se ha producido una actualización del protocolo.
			// En caso de llamar a esta funcion http2 no funcionara
			// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),

			//ConnState especifica una función de devolución de llamada opcional que se llama cuando
			//la conexión de un cliente cambia de estado.
			// ConnState func(net.Conn, ConnState)

			// ErrorLog especifica un registrador opcional para errores al aceptar conexiones, comportamiento inesperado de los controladores y errores subyacentes del sistema de archivos.
			//ErrorLog *log.Logger
		},
	}

	return dan
}

func (dan *DanSpace) NewBaseRoute(handler func(http.ResponseWriter, *http.Request)) {
	
	// Concatenar la base con el host (en este caso, solo se usa el host como referencia)
	dan.Mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// Verificar que la ruta solicitada es exactamente "/"
		if req.URL.Path == "/" {
			handler(res, req) // Usar el handler proporcionado
			return
		}

		// Si no es "/", devolver un 404
		http.NotFound(res, req)
	})
}

func (dan *DanSpace) NewRoute(pattern string, handler func(http.ResponseWriter, *http.Request))(err error){

	if strings.HasSuffix(pattern, "/") {
		return errors.New("la ruta no debe terminar con un slash")
	}

	dan.Mux.HandleFunc(pattern, handler)

	return
}

func (dan *DanSpace) NewRouteWildcard(pattern string, handler func(http.ResponseWriter, *http.Request))(err error){

	if !strings.HasSuffix(pattern, "/") {
		return errors.New("la ruta debe terminar con una slash")
	}

	dan.Mux.HandleFunc(pattern, handler)
	
	return
}

func (dan *DanSpace) NewMidleware(pattern string,
	midleware func(handler http.HandlerFunc) http.HandlerFunc,
	handler func(http.ResponseWriter, *http.Request)) {

	dan.Mux.HandleFunc(pattern, midleware(handler))

}

func (dan *DanSpace) NewChainMidleware(
	pattern string,
	middlewares []func(http.HandlerFunc) http.HandlerFunc,
	handler http.HandlerFunc) {

	var chainFunc http.HandlerFunc = handler

	// Iterate through middlewares in reverse order to apply them in the correct order
	for i := len(middlewares) - 1; i >= 0; i-- {
		chainFunc = middlewares[i](chainFunc)
	}

	dan.Mux.HandleFunc(pattern, chainFunc)
}

func RedirectHttps(domain string) func(response http.ResponseWriter, request *http.Request) {

	return func(response http.ResponseWriter, request *http.Request) {

		target := "https://" + domain + request.URL.Path
		if len(request.URL.RawQuery) > 0 {
			target += "?" + request.URL.RawQuery
		}

		http.Redirect(response, request, target, http.StatusTemporaryRedirect)
	}
}

func (dan *DanSpace) InitDan(publicKey string, privateKey string) {

	err := dan.Server.ListenAndServeTLS(publicKey, privateKey)
	if err != nil {
		println(err.Error())
	}
}

func (dan DanSpace) InitDanInsecure() {

	err := dan.Server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}
}
