project = "izaakdemoapp"

app "izaakdemoapp" {

  build {
    use "docker" {}
    registry {
       use "aws-ecr" {
        region     = "us-east-1"
        repository = "waypoint-example"
        tag        = "latest"
      }
    }
  }

  deploy {
    use "aws-ecs" {
      region = "us-east-1"
      memory = "512"
    }
  }
}
