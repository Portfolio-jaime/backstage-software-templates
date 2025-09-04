# Backstage Software Templates

Este repositorio contiene las **plantillas de Backstage** organizadas por tipo de componente.

## 📁 Estructura del Repositorio

### 🚀 Services (`services/`)
Templates para aplicaciones y microservicios:
- **go-app/** - Aplicación Go básica
- **go-service/** - Microservicio Go con mejores prácticas
- **java-app/** - Aplicación Spring Boot
- **java-service/** - Microservicio Java empresarial  
- **python-app/** - Aplicación Python moderna

### 🏗️ Infrastructure (`infrastructure/`)
Templates para infraestructura y DevOps:
- **terraform/** - Infraestructura base con Terraform
- **terraform-module/** - Módulos reutilizables de Terraform
- **kubernetes/** - Recursos básicos de Kubernetes
- **kubernetes-deployment/** - Deployments avanzados de K8s
- **argocd/** - Configuración de ArgoCD
- **argocd-application/** - Aplicaciones de ArgoCD

### 📚 Libraries (`libraries/`)
Templates para librerías y herramientas:
- **cli-go/** - Herramientas CLI en Go con Cobra

### 🌟 Golden Path (`golden-path/`)
Templates completas end-to-end:
- **go-service/** - Stack completo (Go + Terraform + ArgoCD + CI/CD)

## 🚀 Uso

1. Registra este repositorio en Backstage mediante `catalog-info.yaml`
2. Ve a **"Create..."** en la interfaz de Backstage
3. Selecciona la plantilla deseada por categoría
4. Completa los parámetros requeridos
5. ¡Genera tu componente!

## 🔧 Para Desarrolladores de Templates

- Cada directorio contiene `template.yaml` y carpeta `skeleton/`
- Sigue las convenciones establecidas en plantillas existentes
- Actualiza esta documentación al agregar nuevas templates

---
**Organización:** British Airways DevOps Team  
**Contacto:** jaime.andres.henao.arbelaez@ba.com
