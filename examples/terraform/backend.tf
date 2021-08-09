terraform {
  required_version = ">= 0.13"

  required_providers {
    ansiblevault = {
      source  = "ronaldslc/ansiblevault"
      version = "~> 2.0"
    }
  }
}
