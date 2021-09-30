# main.tf

terraform {
  required_version = ">= 0.14"

  required_providers {
    # Cloud Run support was added on 3.3.0
    google = ">= 3.3"
  }
}

provider "google" {
  credentials = file(var.credentials)
  project = var.project_id
  region  = var.region
}

resource "google_project_service" "run_api" {
  service = "run.googleapis.com"
  disable_on_destroy = true
}

resource "google_project_service" "firestore" {
  service = "firestore.googleapis.com"
  disable_on_destroy = false
}

resource "google_cloud_run_service" "run_service" {
  name = "nastia-go-pets"
  location = var.region

  template {
    spec {
      containers {
        env {
          name = "GOOGLE_CLOUD_PROJECT"
          value = var.project_id
        }
        env {
          name = "GOOGLE_APPLICATION_CREDENTIALS"
          value = var.credentials
        }
        image = "gcr.io/${var.project_id}/go-pets:latest"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Waits for the Cloud Run API to be enabled
  depends_on = [google_project_service.run_api]
}

# Allow unauthenticated users to invoke the service
resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# set up database
resource "google_firestore_document" "pet_1" {
  collection  = "Pet"
  document_id = "pet-1"
  fields      = "{\"petname\":{\"stringValue\":\"Loki\"}, \"owner\":{\"stringValue\":\"James\"}, \"likes\":{\"integerValue\":\"180\"}, \"image\":{\"stringValue\":\"https://i.pinimg.com/originals/73/35/45/733545a67e5fb354181a6516a2e410ba.jpg\"}}"
}

# Display the service URL
output "service_url" {
  value = google_cloud_run_service.run_service.status[0].url
}
