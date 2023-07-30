terraform {
  required_providers {
    email = {
      version = "0.2"
      source  = "watonomous.ca/tf/email"
    }
  }
}

provider "email" {}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
