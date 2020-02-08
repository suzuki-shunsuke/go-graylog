provider "graylog" {
  web_endpoint_uri = var.web_endpoint_uri
  auth_name        = var.auth_name
  auth_password    = var.auth_password
  api_version      = var.api_version
}

variable "web_endpoint_uri" {
}

variable "auth_name" {
}

variable "auth_password" {
}

variable "api_version" {
}

