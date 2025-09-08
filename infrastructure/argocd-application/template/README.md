# ${{values.name}} - ArgoCD GitOps Application

${{values.description}}

## Architecture

This repository follows the GitOps pattern using ArgoCD for continuous deployment:

```
├── ${{values.manifest_path}}/          # Kubernetes manifests
│   ├── deployment.yaml    # Application deployment
│   └── service.yaml       # Service configuration
├── .github/workflows/     # CI/CD pipelines
│   └── cd.yml            # Continuous deployment workflow
├── application.yaml       # ArgoCD application configuration
└── README.md             # This file
```

## ArgoCD Configuration

- **Application Name**: `${{values.name}}`
- **Project**: `${{values.project}}`
- **Target Namespace**: `${{values.namespace}}`
- **Manifest Path**: `${{values.manifest_path}}/`
- **Auto Sync**: {%- if values.auto_sync %}`Enabled`{%- else %}`Disabled`{%- endif %}
- **Self Heal**: {%- if values.self_heal %}`Enabled`{%- else %}`Disabled`{%- endif %}

## Deployment Process

1. **Code Changes**: Developers push changes to the main branch
2. **CI Pipeline**: GitHub Actions validates and scans the manifests
3. **GitOps**: ArgoCD detects changes and syncs to Kubernetes
4. **Monitoring**: ArgoCD tracks application health and status

## Getting Started

### Prerequisites

- ArgoCD installed and configured
- Kubernetes cluster access
- GitHub repository permissions

### Deployment

1. Apply the ArgoCD application:
```bash
kubectl apply -f application.yaml
```

2. Monitor the deployment:
```bash
argocd app get ${{values.name}}
argocd app sync ${{values.name}}
```

## Manifest Structure

### Deployment
- **Replicas**: 2 (for high availability)
- **Resources**: CPU and memory limits configured
- **Health Checks**: Liveness and readiness probes
- **Image**: nginx:latest (update as needed)

### Service
- **Type**: ClusterIP
- **Port**: 80
- **Target Port**: 80

## CI/CD Pipeline

The GitHub Actions workflow includes:

1. **Manifest Validation**: Using kubeval and kubeconform
2. **Security Scanning**: Trivy vulnerability scanner
3. **ArgoCD Notification**: Automatic sync trigger

## Monitoring

Monitor your application using:

```bash
# ArgoCD UI
kubectl port-forward svc/argocd-server -n argocd 8080:443

# Application status
kubectl get pods -n ${{values.namespace}} -l app=${{values.name}}
kubectl describe deployment ${{values.name}} -n ${{values.namespace}}
```

## Troubleshooting

Common issues and solutions:

1. **Sync Failures**: Check ArgoCD application events
2. **Pod Issues**: Verify resource limits and image availability
3. **Network Issues**: Ensure service and ingress configurations

## Contributing

1. Make changes to Kubernetes manifests in `${{values.manifest_path}}/`
2. Test changes in development environment
3. Create pull request with detailed description
4. CI pipeline will validate changes automatically
5. Merge to main triggers ArgoCD sync

## Links

- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
- [GitOps Guide](https://www.weave.works/technologies/gitops/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)