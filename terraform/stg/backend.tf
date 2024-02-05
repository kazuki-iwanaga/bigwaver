terraform {
  required_version = "~> 1.6.6"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.13.0"
    }

    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 5.13.0"
    }
  }

  backend "gcs" {}
}