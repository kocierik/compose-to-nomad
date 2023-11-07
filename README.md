# Compose to Nomad

## Usage

To use this tool, you need to specify the path to your Docker Compose YAML file and the output directory where you want the generated Nomad job specification files to be stored.

```shell
go run cmd/main.go -compose-file docker-compose.yml  -output-dir example
```

---

##  Currently supported

- Service Conversion
- Image Specification 
- Port Mapping 
- Environment Variables
- Volume Mapping
- Dependency Management




