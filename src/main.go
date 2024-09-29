package main

import (
	"embed"
	"fmt"

	"golearn/src/features"
	"golearn/src/internal/handlers"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

//go:embed .\files\*.json
var files embed.FS

func main() {

	features.Files = files

	// Command line arguments
	args := os.Args

	// Printing command line arguments
	fmt.Println(args)

	features.DemonstrateVariables()

	var result, remainder, err = features.DemonstrateFunctions()

	features.DemonstrateErrorHandling(err, remainder, result)

	features.DemonstrateSwitchStatements(result, remainder)

	features.DemonstrateArrays()

	features.DemonstratePointers()

	features.DemonstrateMaps()

	features.DemonstrateSlicePerformance()

	features.DemonstrateStrings()

	features.DemonstrateTypes()

	features.DemonstrateSlices()

	features.DemonstrateGoroutines()

	features.DemonstrateChannels()

	features.DemonstrateGenerics()

	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	srvErr := http.ListenAndServe("localhost:9276", router)

	if err != nil {
		log.Error(srvErr)
	}

}

// go build -o bin/main.exe ./src/  && ./bin/main --hello
// go build -o ../bin/main.exe ./  && ../bin/main --hello
