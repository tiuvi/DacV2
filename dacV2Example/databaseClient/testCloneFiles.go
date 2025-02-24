package main

import (
	"dacV2"
	"dacV2/databaseClient"
	"dacV2/shell"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)


type fileJob struct {
	fullpath string
	dirTruncate string 
	bandwidth int64
	ind int
	itemsTotal int64
}

func worker(jobs <-chan fileJob, wg *sync.WaitGroup) {

	// Siempre que haya un trabajo en el canal, el worker lo ejecuta
	for fileJob := range jobs {

		defer wg.Done()

		partialPath := strings.TrimPrefix(fileJob.fullpath, fileJob.dirTruncate)

		space, err := dacv2.NewSpaceContent(fileJob.fullpath)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}
	
		spaceDB := databaseClient.NewSpaceDBContent("cell1.tiuvi.com", 3001, partialPath)
	
		nRange, err := space.CalcRange(fileJob.bandwidth)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}

		if nRange == 0 {
			err = spaceDB.SetAtRange([]byte{}, 0, fileJob.bandwidth)
			if err != nil {
				shell.ErrorFatal(err.Error())
			}
			continue
		}

		for x := int64(0); x < nRange; x++ {
	
			data, err := space.GetAtRange(x, fileJob.bandwidth)
			if err != nil {
				shell.ErrorFatal(err.Error())
			}
	
			err = spaceDB.SetAtRange(data, x, fileJob.bandwidth)
			if err != nil {
				shell.ErrorFatal(err.Error())
			}
		}

	}
}

/*dir := "/media/franky/tiuviweb/test/dacv2/testnode/node" 
dirTruncate := "/media/franky/tiuviweb/test/dacv2/testnode"
 bandwidth 10485760
testCloneFiles("/media/franky/tiuviweb/test/dacv2/testnode/node" ,
"/media/franky/tiuviweb/test/dacv2/testnode" , 10485760)
*/
func testCloneFiles(dir string , dirTruncate string ,bandwidth int64){

	var fullPaths []string // Array para guardar las rutas completas

	// Función recursiva para listar directorios y subdirectorios
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Println("Error al acceder a la ruta:", err)
			return err
		}
		if !info.IsDir() {
			fullPaths = append(fullPaths, path)
			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatal("Error al recorrer directorios:", err)
	}

	jobs := make(chan fileJob, 50)

	var wg sync.WaitGroup

	// Creamos 3 workers (puedes cambiar el número según lo necesites)
	numWorkers := 50
	for i := 1; i <= numWorkers; i++ {
		go worker(jobs, &wg)
	}



	for  ind  , fullpath := range fullPaths {
		wg.Add(1)
		jobs <- fileJob{
			fullpath: fullpath,
			dirTruncate: dirTruncate,
			bandwidth: bandwidth,
			ind:ind,
			itemsTotal:int64(len(fullPaths)),
		}

	}

	close(jobs)

	wg.Wait()
}

