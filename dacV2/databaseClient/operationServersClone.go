package databaseClient

import (
	"dacV2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type fileJobClone struct {
	isDir       bool
	fullpath    string
	dirTruncate string
	bandwidth   int64
}
 
func fileWorkerCloneFile(domain string, port uint16, jobs <-chan fileJobClone, wg *sync.WaitGroup) {

	// Siempre que haya un trabajo en el canal, el worker lo ejecuta
	for fileJobClone := range jobs {

		partialPath := strings.TrimPrefix(fileJobClone.fullpath, fileJobClone.dirTruncate)

		if fileJobClone.isDir {

			err := CreateDirectory(domain, port, partialPath)
			if err != nil {
				wg.Done()
				log.Println(err.Error())
				continue
			}

			wg.Done()
			continue
		}

		space, err := dacv2.NewSpaceContent(fileJobClone.fullpath)
		if err != nil {
			wg.Done()
			log.Println(err.Error())
			continue
		}

		spaceDB := NewSpaceDBContent(domain, port, partialPath)

		nRange, err := space.CalcRange(fileJobClone.bandwidth)
		if err != nil {
			wg.Done()
			log.Println(err.Error())
			continue
		}

		if nRange == 0 {
			err = spaceDB.SetAtRange([]byte{}, 0, fileJobClone.bandwidth)
			if err != nil {
				wg.Done()
				log.Println(err.Error())
				continue
			}
			wg.Done()
			continue
		}

		for x := int64(0); x < nRange; x++ {

			data, err := space.GetAtRange(x, fileJobClone.bandwidth)
			if err != nil {
				wg.Done()
				log.Println(err.Error())
				continue
			}

			err = spaceDB.SetAtRange(data, x, fileJobClone.bandwidth)
			if err != nil {
				wg.Done()
				log.Println(err.Error())
				continue
			}
		}

		wg.Done()
	}
}


func CloneDirectoryToServer(domain string, port uint16, dir string, dirTruncate string, bandwidth int64, nWorkers int, nJobs int) (err error) {

	var jobsItems []fileJobClone // Array para guardar las rutas completas

	// Función recursiva para listar directorios y subdirectorios
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Println("Error al acceder a la ruta:", err)
			return err
		}

		jobsItems = append(jobsItems, fileJobClone{
			isDir:       info.IsDir(),
			fullpath:    path,
			dirTruncate: dirTruncate,
			bandwidth:   bandwidth,
		})
		return nil
	})
	if err != nil {
		return
	}

	jobs := make(chan fileJobClone, nJobs)

	var wg sync.WaitGroup
	for i := 1; i <= nWorkers; i++ {
		go fileWorkerCloneFile(domain, port, jobs, &wg)
	}

	for _, itemJob := range jobsItems {
		wg.Add(1)
		jobs <- itemJob

	}

	close(jobs)

	wg.Wait()

	return
}
