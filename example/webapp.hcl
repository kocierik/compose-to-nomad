job "webapp" {
	group "group-webapp" {
		network {
			port "webapp_port" {}
		}
		task "task-webapp" {
			driver = "docker"
			config {
				image = "my-web-app:latest"
				ports = ["webapp_port"]
			}
			env {
				DEBUG = "true"
				REDIS_HOST = "redis"
			}
			volume_mount {
				volume = "webapp-data:/var/lib/webapp"
			}
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}