package dacv2

import (
	"path/filepath"
	"sync"
	"time"
)

type CacheEntry struct {
	Space   *Space
	CreatedAt time.Time
}

// SpaceCache gestiona la caché de estructuras Space.
type SpaceCacheExpiration struct {
	Cache  map[string]CacheEntry
	Mu     sync.RWMutex // Mutex para asegurar acceso seguro al mapa en entornos concurrentes
	Ticker *time.Ticker
	TimeExpirationFile time.Duration
}

// NewSpaceCacheExpiration crea una nueva instancia de SpaceCache con expiración periódica.
func NewSpaceCacheExpiration(timeCleanCache time.Duration, timeExpirationFile time.Duration) *SpaceCacheExpiration {

	spaceCache := &SpaceCacheExpiration{
		Cache: make(map[string]CacheEntry), // Inicialización del mapa
		TimeExpirationFile:timeExpirationFile,
	}

	// Ticker para ejecutar la limpieza de la caché cada cierto tiempo
	spaceCache.Ticker = time.NewTicker(timeCleanCache)

	// Goroutine para limpiar la caché periódicamente
	go func() {
		for range spaceCache.Ticker.C {
			spaceCache.Clear()
		}
	}()

	return spaceCache
}

// Open obtiene un Space de la caché. Devuelve nil si no existe.
// Open obtiene un Space de la caché. Devuelve nil si no existe.
func (spaceCache *SpaceCacheExpiration) Open(mapFields map[int64][3]int64, sizeField int64, mapLines map[int64][3]int64, sizeLine int64, dirPath ...string) (space *Space, err error) {

	filePath := filepath.Join(dirPath...)

	spaceCache.Mu.RLock()
	entry, exists := spaceCache.Cache[filePath]
	spaceCache.Mu.RUnlock()
	if exists {
		return entry.Space, nil
	}

	spaceCache.Mu.Lock()

	entry, exists = spaceCache.Cache[filePath]
	if exists {
		spaceCache.Mu.Unlock()
		return entry.Space, nil
	}

	space, err = NewSpace(mapFields, sizeField, mapLines, sizeLine, dirPath...)
	if err != nil {
		spaceCache.Mu.Unlock()
		return nil, err
	}

	// Almacenamos el Space junto con la fecha de adición en la caché
	spaceCache.Cache[filePath] = CacheEntry{
		Space:   space,
		CreatedAt: time.Now(),
	}

	spaceCache.Mu.Unlock()

	return space, nil
}
 
func (spaceCache *SpaceCacheExpiration) OpenSpaceRange(dirPath ...string) (space *Space, err error) {

	filePath := filepath.Join(dirPath...)

	spaceCache.Mu.RLock()
	entry, exists := spaceCache.Cache[filePath]
	spaceCache.Mu.RUnlock()
	if exists {
		return entry.Space, nil
	}

	spaceCache.Mu.Lock()

	entry, exists = spaceCache.Cache[filePath]
	if exists {
		spaceCache.Mu.Unlock()
		return entry.Space, nil
	}

	space, err = NewSpaceContent(dirPath...)
	if err != nil {
		return
	}

	// Almacenamos el Space junto con la fecha de adición en la caché
	spaceCache.Cache[filePath] = CacheEntry{
		Space:   space,
		CreatedAt: time.Now(),
	}

	spaceCache.Mu.Unlock()

	return space, nil
}

// Delete elimina un Space de la caché.
func (spaceCache *SpaceCacheExpiration) Delete(space *Space) {

	spaceCache.Mu.Lock()

	delete(spaceCache.Cache, space.FilePath)

	spaceCache.Mu.Unlock()
}


func (spaceCache *SpaceCacheExpiration) Clear() {

	spaceCache.Mu.Lock()

	now := time.Now()

	// Recorremos las entradas de la caché
	for key, entry := range spaceCache.Cache {

		// Eliminamos elementos que hayan estado abiertos por más de un día
		if now.Sub(entry.CreatedAt) > spaceCache.TimeExpirationFile {

			entry.Space.File.Close() // Cerramos el archivo si es necesario

			delete(spaceCache.Cache, key)

		}
	}

	spaceCache.Mu.Unlock()
}

func (spaceCache *SpaceCacheExpiration) Clean(){

	spaceCache.Mu.Lock()

	for _, entry := range spaceCache.Cache {

		entry.Space.File.Close()

	}

    spaceCache.Cache = make(map[string]CacheEntry)

	spaceCache.Ticker.Stop()
	
	spaceCache.Mu.Unlock()
}
