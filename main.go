package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var (
	inputFile string
	outputDir string
)

func main() {
	flag.StringVar(&inputFile, "input", "site.json", "The main site JSON file")
	flag.StringVar(&outputDir, "output", "/public", "The directory where")
	flag.Parse()

	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic("Could not read file: " + inputFile)
	}

	var site Site
	err = json.Unmarshal([]byte(file), &site)

	if err != nil {
		panic("Could not parse site")
	}

	err = site.Render()

	if err != nil {
		panic("Could not render site")
	}

	err = site.CopyAssets(outputDir)

	if err != nil {
		panic("Could not copy assets.")
	}
}
