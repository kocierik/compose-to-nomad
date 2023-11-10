package nomadTemplate

const NomadJobTemplate = `job "{{.Name}}" {
  datacenters = ["dc1"]

	group "group-{{.Name}}" {

    count = 1

  {{- range .Volumes}} 
    volume "volume-{{.Label}}" {
      type = "host"
      source = "{{.Source}}"
      read_only = false

    }
  {{- end}}

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

    {{- range .Volumes}}
      volume_mount {
        volume = "volume-{{.Label}}"
        destination = "/var/lib/{{.Label}}"
        read_only = false
      }
    {{- end}}

			{{- if .Environment}}
			env {
				{{- range .Environment}}
				{{ $keyval := splitN . "=" 2 }}{{index $keyval 0}} = "{{index $keyval 1}}"
				{{- end}}
			}

			{{- end}}
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}`
