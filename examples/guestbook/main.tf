terraform {
  required_providers {
    carvel = {
      // Local provider
      source  = "carvel.dev/carvel/k14s"
      version = "0.7.0"
    }
  }
}
