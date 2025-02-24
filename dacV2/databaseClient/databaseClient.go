package databaseClient

type SpaceDB struct {
	BaseURL          string
	Port             uint16
	FieldsMapName    int64
	IndexSizeFields  map[int64][3]int64
	LinesMapName     int64
	IndexSizeColumns map[int64][3]int64
	DirPath          []string
}

func NewSpaceDB(baseURL string, port uint16,
	fieldsMapName int64, IndexSizeFields map[int64][3]int64,
	linesMapName int64, IndexSizeColumns map[int64][3]int64,
	dirPath ...string) *SpaceDB {
		
	return &SpaceDB{
		BaseURL:          baseURL,
		Port:             port,
		FieldsMapName:    fieldsMapName,
		IndexSizeFields:  IndexSizeFields,
		LinesMapName:     linesMapName,
		IndexSizeColumns: IndexSizeColumns,
		DirPath:          dirPath,
	}
}

func NewSpaceDBContent(baseURL string, port uint16,dirPath ...string) *SpaceDB {
		
	return &SpaceDB{
		BaseURL:          baseURL,
		Port:             port,
		DirPath:          dirPath,
	}
}