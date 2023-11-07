job "redis" {
	group "group-redis" {
		network {
			port "redis_port" {}
		}
		task "task-redis" {
			driver = "docker"
			config {
				image = "redis:alpine"
				ports = ["redis_port"]
			}
			volume_mount {
				volume = "redis-data:/data"
			}
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}