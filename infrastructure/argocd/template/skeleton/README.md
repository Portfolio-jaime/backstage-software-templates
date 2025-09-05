# ${{values.app_name}} ArgoCD Configuration

This repository contains ArgoCD application and project configurations for deploying ${{values.app_name}} to the ${{values.app_env}} environment.

## Structure

```
argocd-apps/
├── application.yaml    # ArgoCD Application definition
├── appproject.yaml     # ArgoCD AppProject definition
└── kustomization.yaml  # Kustomize configuration
```

## Deployment

### Prerequisites

- ArgoCD installed in the cluster
- Proper RBAC permissions configured
- Source repository with Kubernetes manifests

### Apply Configuration

```bash
# Apply the ArgoCD configurations
kubectl apply -k argocd-apps/
```

### Sync Application

```bash
# Sync the application manually
argocd app sync ${{values.app_name}}-${{values.app_env}}

# Get application status
argocd app get ${{values.app_name}}-${{values.app_env}}
```

## Configuration Details

### Application
- **Name**: ${{values.app_name}}-${{values.app_env}}
- **Project**: ${{values.app_name}}-project
- **Source**: https://github.com/Portfolio-jaime/${{values.app_name}}
- **Target Namespace**: ${{values.app_name}}-${{values.app_env}}
- **Sync Policy**: Automated with self-heal and prune

### Project
- **Name**: ${{values.app_name}}-project
- **Source Repos**: Portfolio-jaime organization repositories
- **Destinations**: ${{values.app_name}}-* namespaces
- **RBAC**: Admin and Developer roles configured

## Monitoring

Access ArgoCD UI to monitor the application:
- Application health status
- Sync status and history
- Resource tree view
- Events and logs

## Troubleshooting

Common issues and solutions:

1. **Sync Failed**: Check source repository and manifest syntax
2. **Permission Denied**: Verify RBAC and AppProject permissions
3. **Resource Not Found**: Ensure target namespace exists
4. **Health Check Failed**: Review application pod logs

## Security

- RBAC configured with least privilege principle
- AppProject restricts allowed resources and destinations
- Source repository access controlled via GitHub permissions