variable "project_id" {
  description = "The GCP Project ID."
  type        = string
}

variable "region" {
  type    = string
}

variable "credentials" {
  description = "Path to your service account key json file."
  type = string
}