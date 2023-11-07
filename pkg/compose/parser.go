package compose

import (
	"github.com/kocierik/compose-to-nomad/pkg/types"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadComposeFile(path string) types.ComposeFile {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read Docker Compose file: %v", err)
	}

	var composeFile types.ComposeFile
	err = yaml.Unmarshal(data, &composeFile)
	if err != nil {
		log.Fatalf("Failed to unmarshal Docker Compose YAML: %v", err)
	}
	return composeFile
}
