package main

import (
	"flag"
	"fmt"
	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis/multichecker"
	"log"
	"os"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "directive config yaml file path")
	flag.Parse()

	fmt.Println("configFilePath: ", configFilePath)

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	config, err := directive.ParseConfigFile(configFile)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	analyzers := directive.NewAnalyzers(config)

	multichecker.Main(analyzers...)
}
