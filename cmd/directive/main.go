package main

import (
	"flag"
	"fmt"
	"github.com/gqlgo/directive"
	"github.com/gqlgo/gqlanalysis"
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

	var analyzers []*gqlanalysis.Analyzer
	for _, c := range config.InputObjectFieldConfig {
		analyzer := directive.InputObjectFieldAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}
	for _, c := range config.ObjectFieldArgumentConfig {
		analyzer := directive.ObjectFieldArgumentAnalyzer(c)
		analyzers = append(analyzers, analyzer)
	}

	multichecker.Main(analyzers...)
}
