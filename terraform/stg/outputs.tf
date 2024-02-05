output "gh_actions_email" {
  description = "The service account email for GitHub Actions"
  value       = module.gh_actions.email
}

output "gh_actions_provider" {
  description = "The Workload Identity Pool Provider for GitHub Actions"
  value       = module.gh_actions.provider
}