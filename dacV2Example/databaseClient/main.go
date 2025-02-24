package main

import (
	"dacV2"
	"dacV2/databaseClient"
	"dacV2/shell"
)

const (
	mapTest1 int64 = iota + 1
)
const (
	col1 int64 = iota + 1
	col2
	col3
	col4
	col5
	col6
	col7
)

var mapDac map[int64][3]int64

func main() {

	var err error
	mapDac, err = databaseClient.CreateMap("cell1.tiuvi.com", 3001, mapTest1, []dacv2.SpaceList{
		{col1, 1},
		{col2, 2},
		{col3, 4},
		{col4, 8},
		{col5, 16},
		{col6, 32},
		{col7, 64},
	})
	if err != nil {
		shell.ErrorFatal(err.Error())
	}

	err = databaseClient.CloneDirectoryToServer("cell1.tiuvi.com", 3001,
		"/media/franky/tiuviweb/test/dacv2/testnode/node",
		"/media/franky/tiuviweb/test/dacv2/testnode",10485760 , 50 , 100)
		if err != nil {
			shell.ErrorFatal(err.Error())
		}

}
