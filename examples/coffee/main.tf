terraform {
  required_providers {
    email = {
      version = "0.2"
      source  = "watonomous.ca/tf/email"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "email_coffees" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.email_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.email_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
