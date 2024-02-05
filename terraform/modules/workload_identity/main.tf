resource "google_iam_workload_identity_pool" "this" {
  project = var.gcp_project_id

  workload_identity_pool_id = var.sa_email_id
}

resource "google_iam_workload_identity_pool_provider" "this" {
  project = var.gcp_project_id

  workload_identity_pool_provider_id = var.sa_email_id

  workload_identity_pool_id = google_iam_workload_identity_pool.this.workload_identity_pool_id

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
  }

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

resource "google_service_account" "this" {
  project = var.gcp_project_id

  account_id = var.sa_email_id
}

resource "google_service_account_iam_member" "this" {
  service_account_id = google_service_account.this.id

  role   = "roles/iam.workloadIdentityUser"
  member = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.this.name}/attribute.repository/${var.gh_repository}"
}