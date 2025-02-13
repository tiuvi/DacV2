package httpReceiver

import(
	"net/http"
)

/*
	Error por fallo en los parametros aceptados.
	Por ejemplo si se pide unos valores del 1 al 10, y se envia un 0 o un 11.
*/
func (SK HttpReceiver) ErrorStatusBadRequest(message string){

	http.Error(SK.response, message, http.StatusBadRequest)
}

/*
	Error por fallos en funciones nativas no previstos en el codigo.

	El caso seria si se pide un archivo pero se desconecta el ssd que esta conectado al pc,
	la funcion de os de golang lanzara un error no previsto pero perfectamente manejable.

	Imagina que tenemos un mapa sincronizado y un archivo. Añadimos un valor al mapa pero
	cuando queremos añadirlo al archivo falla, en este caso abria que borrarlo del mapa para
	evitar corrupcion de datos futura
	
*/
func (SK HttpReceiver) ErrorStatusInternalServerError(message string){

	http.Error(SK.response, message, http.StatusInternalServerError)

}


