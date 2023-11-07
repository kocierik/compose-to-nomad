# Compose to Nomad

## Usage

To use this tool, you need to specify the path to your Docker Compose YAML file and the output directory where you want the generated Nomad job specification files to be stored.

```shell
./nomad-job-generator -compose-file=<path-to-docker-compose-file> -output-dir=<path-to-output-directory>
```

---

##  Currently supported

- Service Conversion
- Image Specification 
- Port Mapping 
- Environment Variables
- Volume Mapping
- Dependency Management




