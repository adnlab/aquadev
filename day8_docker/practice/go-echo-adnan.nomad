job "go-echo-adnan" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "go-echo-adnan" {
      driver = "docker"

      config {
        image = "ladnanm/go-echo-adnan:v2"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "go-echo-adnan"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.adnan-echo-http.rule=Host(\"adnanecho.cupang.efishery.ai\")",
      ]
      check {
        port = "http"
        type = "tcp"
        interval = "15s"
        timeout = "14s"
      }
    }

  }
}
