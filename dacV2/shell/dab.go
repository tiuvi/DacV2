package shell

import (
	"flag"
	"os"
)

var hasFlagParsingBeenExecutedSuccessfully bool = false

func InitParse() {

	help := flag.Bool("help", false, "Muestra la ayuda del programa")

	flag.Parse()

	if *help {
		// Mostrar el uso del programa
		flag.Usage()
		// Terminar el programa después de mostrar la ayuda
		os.Exit(0)
	}
	
	hasFlagParsingBeenExecutedSuccessfully = true
}

type FlagString struct {
	name        string
	description string
	Value       *string
}

// Funciones de inicialización para cada tipo de flag
func InitFlagString(name string, description string) *FlagString {
	flagValue := flag.String(name, "", description)
	return &FlagString{
		name:        name,
		description: description,
		Value:       flagValue,
	}
}

func (FL *FlagString) GetString() string {

	if !hasFlagParsingBeenExecutedSuccessfully {
		InitParse()
	}

	return *FL.Value
}

func (FL *FlagString) GetStringReq() string {

	if !hasFlagParsingBeenExecutedSuccessfully {
		InitParse()
	}

	if *FL.Value == "" {
		ErrorFatal("Este valor " + FL.name + " es requerido")
	}

	return *FL.Value
}

type FlagInt struct {
	name        string
	description string
	Value       *int
}

func InitFlagInt(name string, description string) *FlagInt {

	flagValue := flag.Int(name, 0, description)

	return &FlagInt{
		name:        name,
		description: description,
		Value:       flagValue,
	}
}

// Métodos para obtener los valores de los flags
func (FL *FlagInt) GetInt() int {
	return *FL.Value
}

func (FL *FlagInt) GetIntReq() int {
	if !hasFlagParsingBeenExecutedSuccessfully {
		InitParse()
	}

	if *FL.Value == 0 {
		ErrorFatal("El valor para el flag '" + FL.name + "' es requerido")
	}

	return *FL.Value
}

type FlagInt64 struct {
	name        string
	description string
	Value       *int64
}

func InitFlagInt64(name string, description string) *FlagInt64 {

	flagValue := flag.Int64(name, 0, description)

	return &FlagInt64{
		name:        name,
		description: description,
		Value:       flagValue,
	}
}

// Métodos para obtener los valores de los flags
func (FL *FlagInt64) GetInt64() int64 {
	return *FL.Value
}

func (FL *FlagInt64) GetInt64Req() int64 {
	
	if !hasFlagParsingBeenExecutedSuccessfully {
		InitParse()
	}

	if *FL.Value == 0 {
		ErrorFatal("El valor para el flag '" + FL.name + "' es requerido")
	}

	return *FL.Value
}


type FlagBool struct {
	name        string
	description string
	Value       *bool
}

func InitFlagBool(name string, description string) *FlagBool {
	flagValue := flag.Bool(name, false, description)
	return &FlagBool{
		name:        name,
		description: description,
		Value:       flagValue,
	}
}

func (flagBool *FlagBool) GetBool() bool {
	return *flagBool.Value
}

type FlagFloat64 struct {
	name        string
	description string
	Value       *float64
}

func InitFlagFloat64(name string, description string) *FlagFloat64 {
	flagValue := flag.Float64(name, 0.0, description)
	return &FlagFloat64{
		name:        name,
		description: description,
		Value:       flagValue,
	}
}

func (FL *FlagFloat64) GetFloat64() float64 {
	return *FL.Value
}

func (FL *FlagFloat64) GetFloat64Req() float64 {
	if !hasFlagParsingBeenExecutedSuccessfully {
		InitParse()
	}

	if *FL.Value == 0.0 {
		ErrorFatal("El valor para el flag '" + FL.name + "' es requerido")
	}

	return *FL.Value
}

func Write(message string) {
	os.Stdout.WriteString(message)
	os.Exit(0)
}

func ErrorFatal(message string) {
	os.Stderr.WriteString(message)
	os.Exit(1)
}
