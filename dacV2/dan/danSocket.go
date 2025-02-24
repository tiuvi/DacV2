package dan

import (
	"net/http"
	"strconv"
)

func NewDanSpaceSocket(port uint16) *DanSpace {

	mux := http.NewServeMux()
	
	dan := &DanSpace{

		Mux: mux,
		Server: &http.Server{

			Addr: ":" + strconv.FormatUint(uint64(port), 10),

			//TimeoutHandler devuelve un Handler que ejecuta h con el límite de tiempo dado.
			Handler: mux,

			//ReadTimeout es la duración máxima para leer la solicitud completa,
			ReadTimeout: 0,

			//ReadHeaderTimeout es la cantidad de tiempo permitido para leer los encabezados de solicitud.
			ReadHeaderTimeout: 0,

			//WriteTimeout es la duración máxima antes de que se agote el tiempo de escritura de la respuesta.
			WriteTimeout: 0,

			// IdleTimeout es la cantidad máxima de tiempo para esperar la próxima solicitud cuando se habilita la función Keep-Alives. Si IdleTimeout es cero, se utiliza el valor de ReadTimeout.
			IdleTimeout: 0,

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

	dan.Server.SetKeepAlivesEnabled(false)

	return dan
}
