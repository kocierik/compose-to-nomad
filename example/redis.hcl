job "redis" {
  datacenters = ["dc1"]

	group "group-redis" {

    count = 1 
    volume "volume-redis" {
      type = "host"
      source = "redis-data"
      read_only = false

    }

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
        volume = "volume-redis"
        destination = "/var/lib/redis"
        read_only = false
      }
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}