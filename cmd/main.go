package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kocierik/compose-to-nomad/pkg/compose"
	"github.com/kocierik/compose-to-nomad/pkg/nomad"
	"github.com/kocierik/compose-to-nomad/pkg/utils"
)

func printUsageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Fprintln(os.Stderr, "Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	composeFilePath := flag.String("compose-file", "", "Path to the Docker Compose YAML file")
	outputDirPath := flag.String("output-dir", ".", "Directory to output the generated Nomad job files")
	flag.Parse()

	if *composeFilePath == "" {
		printUsageAndExit(" You must specify a path to a Docker Compose file using the -compose-file flag.\n")
	}

	composeFile := compose.ReadComposeFile(*composeFilePath)

	tmpl := nomad.ParseTemplate()
	utils.CreateOutputDirectory(*outputDirPath)

	for name, service := range composeFile.Services {
		nomad.GenerateNomadJob(name, service, tmpl, *outputDirPath)
	}

}
