# TallerPro - Gestión de Presupuestos Cifrados

Solución de software modular y eficiente para talleres mecánicos basada en **Clean Architecture**.

## Stack Tecnológico
- **Frontend:** React + Tailwind CSS + shadcn/ui (Diseñado en v0)
- **Backend:** Go (Golang) + Wails (Comunicación nativa en memoria)
- **Base de Datos:** SQLite (Costo de infraestructura: 0€)

## Características principales
- Firma electrónica avanzada criptográfica (SHA-256).
- Gestión de anexos de ampliación/reducción de orden.
- Cumplimiento estricto del RGPD.

## 🚀 Guía de Instalación Rápida
Abre la terminal integrada de VS Code en la raíz de tu proyecto y ejecuta los siguientes comandos para descargar las dependencias:

```bash
go get gorm.io/gorm                     # Para descargar GORM     
go get [github.com/glebarez/sqlite](https://github.com/glebarez/sqlite)       # Para descargar el driver de SQLite puro
```

## Estructura de proyecto
internal/domain/domain.go        
# Definición centralizada de tipos y estructuras de datos

internal/repository/sqlite.go   
    # Conexión SQLite pura sin CGO y migraciones automáticas de GORM.
    # Usa el módulo correcto: tallermecanico/internal/domain coincide exactamente con la línea module tallermecanico que tienes definida en tu archivo go.mod.
    # SQLite puro sin CGO: Utiliza [github.com/glebarez/sqlite](https://github.com/glebarez/sqlite) para que no tengas problemas de compilación en Windows.
    # Migraciones automáticas: Con db.AutoMigrate, GORM creará de forma inteligente las tablas para Cliente, Vehiculo y Presupuesto en el archivo local taller.db la primera vez que ejecutes el programa.

internal/usecase/presupuesto.go        
# Implementar lógica de firma criptográfica y mapa optimizado O(N) en el caso de uso.

.env            # Archivo almacena variables de entorno de forma local.

.gitignore      # Archivo configuración para Git que indica qué archivos
                # y carpetas no deben ser rastreados ni subidos al 
                # control de versiones.

go.mod          # Archivo principal de gestión de módulos de Go.
                # Define el nombre del proyecto, la versión de Go que 
                # utiliza y lista las dependencias externas.

go.sum          # Archivo de huellas digitales (hashes criptográficos)
                # generado automáticamente por Go. Su función es 
                # garantizar la integridad y seguridad de las 
                # dependencias descargadas.

Guarda el archivo y ejecuta tu comando de commit:

```bash
git add README.md
git commit -m "docs: actualizar README completo con arquitectura, estructura detallada e instrucciones"
```