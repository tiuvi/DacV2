package databaseServer

import (
	. "dacV2"
	. "dacV2/dan"
	. "dacV2/httpReceiver"
	. "dacV2/shell"
	"net/http"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var globalCache *SpaceCacheExpiration

var globalPath string

var Routes = make(map[string]func(http.ResponseWriter, *http.Request))

var RoutesDac = make(map[string]func())

var keyRoutes string = ""

const (
	colNameMap = 0
	colSizeMap = 1
	colMap     = 2
)

type mapDacItem struct {
	Map     map[int64][3]int64
	sizeMap int64
}

var mapsDac map[int64]mapDacItem = make(map[int64]mapDacItem)

var fileMaps *Space

func init() {

	RoutesDac["initMaps"] = func() {

		mapLine, sizeLine, err := CreateMap([]SpaceList{
			{Name: colNameMap, Len: 8},
			{Name: colSizeMap, Len: 8},
			{Name: colMap, Len: 1000},
		})
		if err != nil {
			ErrorFatal(err.Error())
			return
		}

		fileMaps, err = NewSpace(nil, 0, mapLine, sizeLine, globalPath, "maps", "maps.dacByte")
		if err != nil {
			ErrorFatal(err.Error())
			return
		}

		for idLine := int64(0); idLine < fileMaps.CountLines(); idLine++ {

			nameMap, err := fileMaps.GetLineInt64(colNameMap, idLine)
			if err != nil {
				return
			}
	
			sizeMap, err := fileMaps.GetLineInt64(colSizeMap, idLine)
			if err != nil {
				return
			}

			outRef := make(map[int64][3]int64)
			err = fileMaps.GetLineDecoder(colMap, idLine, &outRef)
			if err != nil {
				return
			}

			mapsDac[nameMap] = mapDacItem{
				Map:     outRef,
				sizeMap: sizeMap,
			}

		}

		globalPath = filepath.Join(globalPath , "database")


	}
}

func init() {

	Routes["/"] = func(res http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		
		SK := InitHttpReceiver(res, req)

		if SK.ExistUrl("terminal") {

			SK.Write("html", []byte(
				//Mata el proceso

				`kill `+strconv.Itoa(syscall.Getpid())+` && \<br>`+
					//Limpia la consola
					`clear`+` && \<br>`))
			return
		}

		err := SK.Write("txt", []byte("dacV2"))
		if err != nil {
			println("Error al escribir en la conexion: ", err.Error())
		}
	}
}


func NewServerDacV2() {

	//path keyRoutes publicKey privateKey host port
	pathFlag := InitFlagString("path", "La carpeta donde se alojara la base de datos")

	keyRoutesFlag := InitFlagString("keyRoutes", "Una key que se compara en cada peticion")

	publicKeyFlag := InitFlagString("publicKey", "La carpeta donde se alojara la base de datos")

	privateKeyFlag := InitFlagString("privateKey", "La carpeta donde se alojara la base de datos")

	hostFlag := InitFlagString("host", "Elige un dominio")

	portFlag := InitFlagInt64("port", "Elige el valor del puerto.")

	InitParse()

	globalPath = pathFlag.GetStringReq()

	publicKey := publicKeyFlag.GetStringReq()

	privateKey := privateKeyFlag.GetStringReq()

	keyRoutes = keyRoutesFlag.GetStringReq()

	_ = hostFlag.GetStringReq()

	port := portFlag.GetInt64Req()

	globalCache = NewSpaceCacheExpiration(time.Hour*12, time.Hour*24)

	for _, fileDac := range RoutesDac {

		fileDac()
	}

	dan := NewDanSpace(uint16(port))

	//Cargamos todas las rutas.
	for path, route := range Routes {
		dan.NewRoute(path, route)
	}

	publicKeyPath, err := CreateTempBytes([]byte(publicKey))
	if err != nil {
		ErrorFatal(err.Error())
	}

	privateKeyPath, err := CreateTempBytes([]byte(privateKey))
	if err != nil {
		ErrorFatal(err.Error())
	}

	dan.InitDan(publicKeyPath, privateKeyPath)

}
