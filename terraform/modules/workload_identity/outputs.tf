output "email" {
  description = "The email ID of the service account"
  value       = google_service_account.this.email
}

output "provider" {
  description = "The Workload Identity Pool Provider"
  value       = google_iam_workload_identity_pool_provider.this.id
}