module "gh_actions" {
  source         = "../modules/workload_identity"
  gcp_project_id = var.gcp_project_id
  sa_email_id    = "gh-actions"
  gh_repository  = var.gh_repository
}

resource "google_project_iam_member" "github_actions" {
  for_each = toset([
    "roles/viewer"
  ])

  project = var.gcp_project_id
  role    = each.value
  member  = "serviceAccount:${module.gh_actions.email}"
}