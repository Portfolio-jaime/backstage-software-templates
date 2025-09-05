# Kubernetes manifests container for GitOps
FROM alpine/k8s:1.27.3

WORKDIR /manifests

# Copy Kubernetes manifests
COPY manifests/ ./manifests/
COPY kustomization.yaml ./

# Install kubectl and kustomize if not available
RUN apk add --no-cache curl && \
    curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash && \
    mv kustomize /usr/local/bin/

# Default command
CMD ["kubectl", "version", "--client"]

