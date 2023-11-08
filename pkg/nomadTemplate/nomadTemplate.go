package nomadTemplate

const NomadJobTemplate = `job "{{.Name}}" {
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
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}`
