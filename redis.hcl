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
				network {
					mbits = 10
					port "redis_port" {
						static = 6379
					}
				}
			}
		}
	}
}