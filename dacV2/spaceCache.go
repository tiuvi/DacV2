package dacv2

import (
	"path/filepath"
	"sync"

)

// SpaceCache gestiona la caché de estructuras Space.
type SpaceCache struct {
	Cache map[string]*Space
	Mu    sync.RWMutex // Mutex para asegurar acceso seguro al mapa en entornos concurrentes
}

// NewSpaceCache crea una nueva instancia de SpaceCache.
func NewSpaceCache() *SpaceCache {
	
	return &SpaceCache{
		Cache: make(map[string]*Space), // Inicialización del mapa
	}
}

// Open obtiene un Space de la caché. Devuelve nil si no existe.
func (spaceCache *SpaceCache) Open(mapFields map[int64][3]int64, sizeField int64, mapLines map[int64][3]int64, sizeLine int64, dirPath ...string) (space *Space, err error) {

	filePath := filepath.Join(dirPath...)

	spaceCache.Mu.RLock()
	value, exists := spaceCache.Cache[filePath]
	spaceCache.Mu.RUnlock()
	if exists {

		return value, nil
	}

	spaceCache.Mu.Lock()

	value, exists = spaceCache.Cache[filePath]
	if exists {
		spaceCache.Mu.Unlock()
		return value, nil
	}

	space, err = NewSpace(mapFields, sizeField, mapLines, sizeLine, dirPath...)
	if err != nil {
		spaceCache.Mu.Unlock()
		return nil, err
	}

	// Almacenamos el Space en la caché
	spaceCache.Cache[filePath] = space

	spaceCache.Mu.Unlock()

	return space, nil
}

// Delete elimina un Space de la caché.
func (spaceCache *SpaceCache) Delete(space *Space) {

	spaceCache.Mu.Lock()

	delete(spaceCache.Cache, space.FilePath)

	spaceCache.Mu.Unlock()
}

func (spaceCache *SpaceCache) Clean(){

	spaceCache.Mu.Lock()

	for _, space := range spaceCache.Cache {

		space.File.Close()

	}

    spaceCache.Cache = make(map[string]*Space)

	spaceCache.Mu.Unlock()
}