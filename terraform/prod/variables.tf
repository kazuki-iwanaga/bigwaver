variable "gcp_project_id" {
  type        = string
  description = "GCP Project ID"
}

variable "gcp_region" {
  type        = string
  description = "GCP Region"
  default     = "asia-northeast1"
}

variable "gh_repository" {
  type        = string
  description = "The GitHub repository name like OWNER/REPO"
}