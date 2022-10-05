project = "izaakdemoapp"

app "izaakdemoapp" {

  build {
    use "docker" {}

    #registry {
    #   use "aws-ecr" {
    #    region     = "us-east-2"
    #    repository = "izaakdemoapp"
    #    tag        = "latest"
    #  }
    #}

    registry {
     use "docker" {
       image = "ttl.sh/kubernetes-nodejs-web"
       tag   = "1h"
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

 # deploy {
 #   use "aws-ecs" {
 #     region = "us-east-2"
 #     memory = "512"
 #   }
 # }
}


