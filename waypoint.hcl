project = "izaakdemoapp"

app "izaakdemoapp" {

  build {
    use "docker" {}
    registry {
      use "docker" {
        image = "ttl.sh/izaakprodapp-longobscurestring"
        tag        = "1h"
      }
    }
  }

  deploy {
    use "kubernetes" {
      probe_path = "/"
    }
  }

  release {
    use "kubernetes" {
      // Sets up a load balancer to access released application
      load_balancer = true
      port          = 3000
    }
  }
}
