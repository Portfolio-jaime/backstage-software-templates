# ArgoCD configuration container
FROM argoproj/argocd:v2.8.0

WORKDIR /argocd

# Copy ArgoCD configurations
COPY argocd-apps/ ./argocd-apps/
COPY README.md ./

# Install kubectl for deployment
USER root
RUN apt-get update && apt-get install -y curl && \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

USER argocd

# Default command
CMD ["argocd", "version"]

