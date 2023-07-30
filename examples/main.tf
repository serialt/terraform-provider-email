terraform {
  required_providers {
    email = {
      version = "0.2.1"
      source  = "watonomous.ca/tf/email"
    }
  }
}

provider "email" {}

resource "email_email" "example" {
  to = "infra-outreach@watonomous.ca"
  from = "sentry-outgoing@watonomous.ca"
  reply_to = "infrastructure@watonomous.ca"
  subject = "Hello from Terraform"
  body = "This is a test email sent from Terraform using a custom email provider."
  smtp_server = "smtp.gmail.com"
  smtp_port = "587"
  smtp_username = "mailbot@watonomous.ca"
  smtp_password = "<replace_me>"
}
