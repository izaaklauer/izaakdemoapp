project = "izaakdemoapp"

app "izaakdemoapp" {

  build {
    use "docker" {}
    registry {
       use "aws-ecr" {
        region     = "us-east-2"
        repository = "izaakdemoapp"
        tag        = "latest"
      }
    }
  }

  deploy {
    use "aws-ecs" {
      region = "us-east-2"
      memory = "512"
    }
  }
}
