job "webapp" {
  datacenters = ["dc1"]

	group "group-webapp" {

    count = 1 
    volume "volume-webapp" {
      type = "host"
      source = "webapp-data"
      read_only = false

    }

		network {
			port "webapp_port" {}
		}

		task "task-webapp" {
			driver = "docker"
			config {
				image = "my-web-app:latest"
				ports = ["webapp_port"]
			}
      volume_mount {
        volume = "volume-webapp"
        destination = "/var/lib/webapp"
        read_only = false
      }
			env {
				DEBUG = "true"
				REDIS_HOST = "redis"
			}
			resources {
        cpu = 1000
        memory = 1000
			}
		}
	}
}