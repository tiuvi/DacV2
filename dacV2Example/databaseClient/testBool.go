package main

import (
	"dacV2/databaseClient"
	"dacV2/shell"
	"sync"
)

func testBool(nLines int64){

	space := databaseClient.NewSpaceDB("cell1.tiuvi.com", 
	3001, mapTest1 , mapDac ,mapTest1 , mapDac, "testBool.bin")

	err := space.TruncateZeroFile()
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	dataBool ,err := space.GetFieldBool(col1)
	if err != nil {
		shell.ErrorFatal(err.Error())
	}	

	if dataBool != false {
		shell.ErrorFatal("Error en las funciones de field")
	}

	err = space.SetFieldBool(col1 , true)
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	dataBool ,err = space.GetFieldBool(col1)
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	if dataBool != true {
		shell.ErrorFatal("Error en las funciones fields")
	}

	for x:= nLines ; x >= 0; x--{
		
		dataBool ,err = space.GetLineBool(col1 , x)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}
		
		if dataBool != false {
			shell.ErrorFatal("Error en las funciones de lineas")
		}
	}

	for x:= int64(0) ; x <= nLines; x++{
		
		line , err := space.NewLineBool(col1 , true)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}
		
		if line != x {
			println("Error al crear lineas de manera sincrona" ,line , x)
			shell.ErrorFatal("")
		}
	}


	for x:= nLines ; x >= 0; x--{
		
		dataBool ,err = space.GetLineBool(col1 , x)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}

		if dataBool != true {
			shell.ErrorFatal("Error en las funciones de lineas")
		}
	}

	err = space.TruncateZeroFile()
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	for x:= nLines ; x >= 0; x--{
		
		err := space.SetLineBool(col1 , x , true)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}
		
	}

	for x:= nLines ; x >= 0; x--{
		
		dataBool ,err = space.GetLineBool(col1 , x)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}

		if dataBool != true {
			shell.ErrorFatal("Error en las funciones de lineas")
		}
	}

	err = space.TruncateZeroFile()
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	var wg sync.WaitGroup

	for x := nLines; x >= 0; x-- {

		wg.Add(1) // Agregamos una goroutine al WaitGroup

		go func(x int64) {
			defer wg.Done() // Marcamos la goroutine como terminada al salir

			err := space.SetLineBool(col1, x, true)
			if err != nil {
				shell.ErrorFatal(err.Error())
			}
		}(x) // Pasamos `x` como argumento para evitar problemas con la concurrencia
	}

	wg.Wait() 

	for x:= nLines ; x >= 0; x--{
		
		dataBool ,err = space.GetLineBool(col1 , x)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}

		if dataBool != true {
			shell.ErrorFatal("Error en las funciones de lineas asincronas")
		}
	}

	err = space.DeleteFile()
	if err != nil {
		shell.ErrorFatal(err.Error())
	}


}