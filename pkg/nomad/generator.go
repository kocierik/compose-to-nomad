package nomad

import (
	"fmt"
	"github.com/kocierik/compose-to-nomad/pkg/nomadTemplate"
	"github.com/kocierik/compose-to-nomad/pkg/types"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func ParseTemplate() *template.Template {
	funcMap := template.FuncMap{
		"splitN": strings.SplitN,
	}
	tmpl, err := template.New("nomad").Funcs(funcMap).Parse(nomadTemplate.NomadJobTemplate)
	if err != nil {
		log.Fatalf("Failed to parse Nomad template: %v", err)
	}
	return tmpl
}

func GenerateNomadJob(name string, service types.ServiceConfig, tmpl *template.Template, outputDir string) {
	jobData := prepareJobData(name, service)
	outputFilePath := fmt.Sprintf("%s/%s.hcl", outputDir, name)
	fmt.Printf("Nomad job specification for service '%s' has been written to %s/%s.hcl\n", name, outputDir, name)
	writeNomadJobFile(outputFilePath, jobData, tmpl)
}

func prepareJobData(name string, service types.ServiceConfig) map[string]interface{} {
	ports := make([]map[string]interface{}, 0)
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
	return map[string]interface{}{
		"Name":        name,
		"Image":       service.Image,
		"Ports":       ports,
		"Environment": service.Environment,
	}
}

func writeNomadJobFile(filePath string, jobData map[string]interface{}, tmpl *template.Template) {
	outputFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create output file for service '%s': %v", jobData["Name"], err)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, jobData)
	if err != nil {
		log.Fatalf("Failed to execute template for service '%s': %v", jobData["Name"], err)
	}
}
