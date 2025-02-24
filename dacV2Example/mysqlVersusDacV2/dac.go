package main

import (
	. "dacV2"
	"fmt"
	"log"
	"time"
)

//

const (
	col1 int64 = 1
)

const totalOperation = 1000000

func testDac(){

	println("total de operacion dacv2 de escritura y lectura: " , totalOperation * 2)

	mapLine , sizeLine , err := CreateMap([]SpaceList{
		{Name:col1, Len:1},
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	sf , err := NewSpace(nil , 0 , mapLine , sizeLine , "/home/franky/dacTest/" , "test.dacByte")
	if err != nil {
		log.Fatalln(err.Error())
	}

	start := time.Now()

	for i := 0; i < totalOperation; i++ {

		sf.NewLineBool(col1, i%2 == 0)

	}

	duration := time.Since(start)
	fmt.Printf("✅ Insercion completa dacv2. Tiempo transcurrido: %s\n", duration)

	start = time.Now()

	for i := 0; i < totalOperation; i++ {

		sf.GetLineBool(col1, int64(i))
	}

	duration = time.Since(start)
	fmt.Printf("✅ Lectura completa dacv2. Tiempo transcurrido: %s\n", duration)

	sf.TruncateZeroFile()

}
