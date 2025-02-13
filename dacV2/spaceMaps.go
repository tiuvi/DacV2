package dacv2

// Funciones para crear columnas y fields
type SpaceList struct {
	Name int64
	Len  int64
}

type SpaceListContainer struct {
	Spaces []SpaceList
}

func NewSpaceListContainer() *SpaceListContainer {
	return &SpaceListContainer{
		Spaces: make([]SpaceList, 0), // Inicializa el slice vacío
	}
}

/*
Ejemplo de nombres de variables
const (

	colTest1 int64 = iota + 1  // 1
	colTest2             // 2

)
*/
func (slc *SpaceListContainer) AppendSpaceList(name int64, length int64) {
	slc.Spaces = append(slc.Spaces, SpaceList{Name: name, Len: length})
}

const (
	//Este campo equivaldria a una linea completa o todos los campos juntos
	FullLine   int64 = 9223372036854775807
	FullFields int64 = 9223372036854775807
	
	//Para espacios de contenido
	FieldSizeFile int64 = 9223372036854775806
	FieldContentFile int64 = 9223372036854775805
)

// CreateMapSpaceList crea un mapa basado en una lista de espacios (SpaceList).
// El mapa contiene el rango inicial y final de cada entrada en la lista.
func (slc *SpaceListContainer) CreateMapSpaceList() (spaceListMap map[int64][3]int64, lenSpace int64,  err error) {

	return CreateMap(slc.Spaces)
}

/*
Example:

	sizeField, mapFields, err := CreateMap([]SpaceList{
		{Name:colTest1, Len:10},
	})
*/
func CreateMap(spaceList []SpaceList) (mapDac map[int64][3]int64, sizeMap int64,  err error) {

	// Inicializa el mapa para almacenar los resultados.
	mapDac = make(map[int64][3]int64)

	for ind, value := range spaceList {
		// Validación: no se permiten longitudes negativas.
		if value.Len < 0 {
			return nil, 0,  ErrNegativeLength
		}

		// Validación: evitar nombres duplicados en el mapa.
		if _, exists := mapDac[value.Name]; exists {
			return  nil, 0, ErrDuplicateName
		}

		// Determina el rango inicial y final para cada entrada.
		if ind == 0 {
			mapDac[value.Name] = [3]int64{0, value.Len, value.Len}
		} else {
			mapDac[value.Name] = [3]int64{sizeMap, value.Len + sizeMap, value.Len}
		}

		// Actualiza el puntero para el próximo inicio.
		sizeMap += value.Len
	}

	//Obtener todas columnas o fields
	mapDac[FullLine] = [3]int64{0, sizeMap, sizeMap}

	return mapDac, sizeMap,  nil
}

func CreateMapFieldsLines(fields []SpaceList , lines []SpaceList)(
	 mapField map[int64][3]int64, sizeField int64,
	 mapLine map[int64][3]int64, sizeLine int64, err error) {

		mapField , sizeField ,  err = CreateMap(fields)
	if err != nil {
		return
	}

	mapLine , sizeLine ,  err = CreateMap(lines)
	if err != nil {
		return
	}
	
	return
}