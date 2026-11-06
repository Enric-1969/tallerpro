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

## Estructura del Proyecto

```text
internal/domain/domain.go        
# Definición centralizada de tipos y estructuras de datos

internal/usecase/presupuesto.go        
# Implementar lógica de firma criptográfica y mapa optimizado O(N) en 
# el caso de uso.

.env            # Archivo almacena variables de entorno de forma local.

.gitignore      # Archivo configuración para Git que indica qué archivos
                # y carpetas no deben ser rastreados ni subidos al 
                # control de versiones.

go.mod          # Archivo principal de gestión de módulos de Go.
                # Define el nombre del proyecto, la versión de Go que 
                # utiliza y lista las dependencias externas

go.sum          # Archivo de huellas digitales (hashes criptográficos)
                # generado automáticamente por Go. Su función es 
                # garantizar la integridad y seguridad de las 
                # dependencias descargadas

```