project = "simple-go-httpserver"

app "httpserver" {
  labels = {
    "service" = "httpserver",
    "env"     = "dev"
  }

  build {
    use "pack" {}
  }

  deploy {
    use "docker" {}
  }
}