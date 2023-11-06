package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type ComposeFile struct {
	Version  string
	Services map[string]ServiceConfig
}

type ServiceConfig struct {
	Image       string
	Ports       []string
	Environment []string
	Volumes     []string
	DependsOn   []string `yaml:"depends_on"`
}

const nomadJobTemplate = `job "{{.Name}}" {
	group "group-{{.Name}}" {
		network {
			{{- range .Ports}}
			port "{{.Label}}" {}
			{{- end}}
		}
		task "task-{{.Name}}" {
			driver = "docker"
			config {
				image = "{{.Image}}"
				ports = [{{range $index, $p := .Ports}}{{if $index}},{{end}}"{{$p.Label}}"{{end}}]
			}
			{{- if .Environment}}
			env {
				{{- range .Environment}}
				{{ $keyval := splitN . "=" 2 }}{{index $keyval 0}} = "{{index $keyval 1}}"
				{{- end}}
			}
			{{- end}}
			{{- if .Volumes}}
			volume_mount {
				{{- range .Volumes}}
				volume = "{{.}}"
				{{- end}}
			}
			{{- end}}
			resources {
				network {
					{{- range .Ports}}
					mbits = 10
					port "{{.Label}}" {
						static = {{.ContainerPort}}
					}
					{{- end}}
				}
			}
		}
	}
}`

func main() {
	composeFilePath := flag.String("compose-file", "", "Path to the Docker Compose YAML file")
	outputDirPath := flag.String("output-dir", ".", "Directory to output the generated Nomad job files")
	flag.Parse()

	if *composeFilePath == "" {
		log.Fatal("You must specify a path to a Docker Compose file using the -compose-file flag.")
	}

	data, err := os.ReadFile(*composeFilePath)
	if err != nil {
		log.Fatalf("Failed to read Docker Compose file: %v", err)
	}

	var composeFile ComposeFile
	err = yaml.Unmarshal(data, &composeFile)
	if err != nil {
		log.Fatalf("Failed to unmarshal Docker Compose YAML: %v", err)
	}

	funcMap := template.FuncMap{
		"splitN": strings.SplitN,
	}
	tmpl, err := template.New("nomad").Funcs(funcMap).Parse(nomadJobTemplate)
	if err != nil {
		log.Fatalf("Failed to parse Nomad template: %v", err)
	}

	err = os.MkdirAll(*outputDirPath, 0755)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	for name, service := range composeFile.Services {
		var ports []map[string]interface{}
		for _, portMapping := range service.Ports {
			parts := strings.Split(portMapping, ":")
			if len(parts) == 2 {
				hostPort, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("Invalid host port '%s' for service '%s': %v", parts[0], name, err)
				}
				ports = append(ports, map[string]interface{}{
					"HostPort":      hostPort,
					"ContainerPort": parts[1],
					"Label":         fmt.Sprintf("%s_port", name),
				})
			}
		}

		jobData := map[string]interface{}{
			"Name":        name,
			"Image":       service.Image,
			"Ports":       ports,
			"Environment": service.Environment,
			"Volumes":     service.Volumes,
		}

		outputFilePath := fmt.Sprintf("%s/%s.hcl", *outputDirPath, name)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			log.Fatalf("Failed to create output file for service '%s': %v", name, err)
		}
		defer outputFile.Close()

		err = tmpl.Execute(outputFile, jobData)
		if err != nil {
			log.Fatalf("Failed to execute template for service '%s': %v", name, err)
		}

		fmt.Printf("Nomad job specification for service '%s' has been written to %s\n", name, outputFilePath)
	}
}
