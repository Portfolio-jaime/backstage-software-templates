# 🚀 Plan de Reorganización - Backstage Templates

**Fecha:** 3 de Septiembre, 2025  
**Estado:** Pendiente de ejecución  
**Responsable:** Jaime Henao

## 📋 Resumen del Plan

Reorganizar el repositorio de plantillas Backstage segmentando por **tipo de componente** para mejorar la organización y mantenibilidad.

## 🎯 Objetivos

- ✅ **Análisis completado:** Estructura actual identificada
- ✅ **Diseño completado:** Nueva estructura definida  
- ⏳ **Pendiente:** Ejecución de migración
- ⏳ **Pendiente:** Actualización de documentación
- ⏳ **Pendiente:** Testing y validación

## 📁 Nueva Estructura Propuesta

```
├── services/                    # 🚀 Aplicaciones y microservicios
│   ├── go-app/                 # Aplicación Go básica
│   ├── go-service/             # Microservicio Go
│   ├── java-app/               # Aplicación Spring Boot
│   ├── java-service/           # Microservicio Java
│   └── python-app/             # Aplicación Python
├── infrastructure/             # 🏗️ Infraestructura y DevOps
│   ├── terraform/              # Infraestructura base
│   ├── terraform-module/       # Módulos reutilizables
│   ├── kubernetes/             # Recursos K8s
│   ├── kubernetes-deployment/  # Deployments K8s
│   ├── argocd/                 # Configuración ArgoCD
│   └── argocd-application/     # Apps ArgoCD
├── libraries/                  # 📚 Librerías y herramientas
│   └── cli-go/                 # Herramientas CLI en Go
├── golden-path/                # 🌟 Templates completas
│   └── go-service/             # Stack completo
└── README.md
```

## 🔧 Comandos de Migración

### Paso 1: Crear estructura
```bash
# Actualizar repositorio
git pull origin main
git checkout -b feat/reorganize-templates-by-type

# Crear directorios por tipo
mkdir -p services infrastructure libraries
```

### Paso 2: Mover plantillas
```bash
# Mover servicios
mv go-app go-service java-app java-service python-app services/

# Mover infraestructura
mv terraform terraform-module kubernetes kubernetes-deployment argocd argocd-application infrastructure/

# Mover librerías
mv cli-go libraries/

# golden-path permanece en raíz (ya está bien ubicado)
```

### Paso 3: Actualizar catalog-info.yaml
```yaml
# Reemplazar content actual con:
apiVersion: backstage.io/v1alpha1
kind: Location
metadata:
  name: backstage-templates
  description: Plantillas de software organizadas por tipo
spec:
  targets:
    - ./services/*/template.yaml
    - ./infrastructure/*/template.yaml
    - ./libraries/*/template.yaml
    - ./golden-path/*/template.yaml
```

### Paso 4: Actualizar README.md
```bash
# Actualizar con nueva estructura y descripción de categorías
# Ver detalles en sección "Nuevo README" abajo
```

### Paso 5: Commit y Push
```bash
git add .
git commit -m "feat: reorganize templates by type (services, infrastructure, libraries)

- Move service templates to services/ directory
- Move infrastructure templates to infrastructure/ directory  
- Move library templates to libraries/ directory
- Update catalog-info.yaml with new paths
- Update README.md with new structure documentation

Author: Jaime Henao <jaime.andres.henao.arbelaez@ba.com>"

git push -u origin feat/reorganize-templates-by-type
```

## 📖 Nuevo README.md

```markdown
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
```

## ✅ Checklist de Validación

Después de la migración, verificar:

- [ ] Todas las plantillas aparecen en Backstage UI
- [ ] Cada template genera código correctamente
- [ ] No hay rutas rotas en catalog-info.yaml
- [ ] README.md refleja nueva estructura
- [ ] Commit sigue convenciones de git
- [ ] Branch está listo para PR

## 🚨 Impacto en Backstage

**✅ No afecta a usuarios finales:**
- Mismas plantillas disponibles
- Mismo proceso de creación
- Misma funcionalidad

**✅ Solo cambia organización interna:**
- Mejor mantenibilidad del repo
- Categorización clara por tipo
- Escalabilidad para futuras templates

## 📞 Contacto

**Jaime Henao**  
DevOps Engineer - British Airways  
📧 jaime.andres.henao.arbelaez@ba.com

---

**⏰ Próximos pasos:** Ejecutar comandos de migración mañana y crear PR para revisión del equipo.