# 🏗️ Ejemplo Práctico Completo: Arquitectura Limpia, Docker y Despliegue en la Nube

En Arquitectura Limpia (*Clean Architecture*), la interfaz de usuario (ya sea una API web, una pantalla móvil o plantillas HTML) se considera un detalle de infraestructura o entorno exterior. Pertenece por completo a la capa de **UI / Infraestructura**.

El backend en Go procesa el caso de uso de forma pura y desacoplada, expone los datos y los archivos del Frontend (HTML/JS/CSS), y finalmente Docker empaqueta absolutamente todo en un único contenedor ligero y aislado, listo para producción.

---

## 📂 1. Estructura de Carpetas del Proyecto

```text
mi-proyecto/
├── core/
│   └── usecase.go        # Lógica pura (Go): Detectar duplicados de forma eficiente
├── infrastructure/
│   ├── web/
│   │   ├── handler.go    # Controlador HTTP en Go (UI/API)
│   │   └── public/       # El Frontend (Detalle de Infraestructura)
│   │       ├── index.html
│   │       ├── style.css
│   │       └── app.js
├── main.go               # Punto de entrada de la aplicación
├── Dockerfile            # Configuración de Docker
```

---

## 🏗️ 2. El Backend en Go

### core/usecase.go (Lógica pura de negocio)
Aquí aplicamos el truco de la preasignación de memoria con mapas para lograr eficiencia $O(N)$ en tiempo y $O(N)$ en memoria. 

```go
package core

// AnalizarDatos gasta O(N) tiempo gracias al mapa con tamaño preasignado.
func AnalizarDatos(numeros []int) (bool, map[int]bool) {
	// Preasignamos el mapa para ayudar al Garbage Collector
	vistos := make(map[int]bool, len(numeros))
	tieneDuplicados := false

	for _, num := range numeros {
		if vistos[num] {
			tieneDuplicados = true
		}
		vistos[num] = true
	}
	// Devolvemos el resultado y el mapa inmutable para el frontend
	return tieneDuplicados, vistos
}
```

### infrastructure/web/handler.go (Capa de Entrada / API Web)
Este manejador recibe los datos del frontend, llama a la lógica de negocio y responde con un JSON.

```go
package web

import (
	"encoding/json"
	"mi-proyecto/core"
	"net/http"
)

type RequestData struct {
	Numeros []int `json:"numeros"`
}

type ResponseData struct {
	Duplicados bool         `json:"duplicados"`
	MapaVistos map[int]bool `json:"mapaVistos"`
}

func ProcesarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req RequestData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Llamamos al caso de uso puro del core
	tieneDuplicados, mapaVistos := core.AnalizarDatos(req.Numeros)

	res := ResponseData{Duplicados: tieneDuplicados, MapaVistos: mapaVistos}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
```

### main.go (Arrancador del Servidor)
Une las piezas y sirve tanto la API como los archivos del Frontend.

```go
package main

import (
	"log"
	"mi-proyecto/infrastructure/web"
	"net/http"
)

func main() {
	// Servir los archivos estáticos del frontend (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./infrastructure/web/public"))
	http.Handle("/", fs)

	// Ruta de la API
	http.HandleFunc("/api/procesar", web.ProcesarHandler)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## 🎨 3. El Frontend (UI / Detalle de Infraestructura)

Ubicado dentro de la carpeta `infrastructure/web/public/`. Usamos **Tailwind CSS** vía CDN para estructurarlo de forma visual y moderna.

### infrastructure/web/public/index.html
```html
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Analizador Eficiente Big O</title>
    <!-- Tailwind CSS v4 para un diseño limpio y rápido (Ruta CDN absoluta corregida) -->
    <script src="https://jsdelivr.net"></script>
    <!-- Enlace al archivo de estilos local -->
    <link rel="stylesheet" href="style.css">
</head>
<body class="bg-slate-900 text-slate-100 flex items-center justify-center min-h-screen p-4">
    <div class="bg-slate-800 p-6 rounded-2xl shadow-xl w-full max-w-md border border-slate-700">
        <h1 class="text-xl font-bold mb-4 text-emerald-400">🧠 Analizador de Memoria O(1)</h1>
        <p class="text-sm text-slate-400 mb-6">Introduce números separados por comas para enviarlos al backend en Go.</p>
        
        <input id="inputNumeros" type="text" value="10, 25, 42, 10, 88" 
               class="w-full p-3 bg-slate-950 rounded-lg border border-slate-600 focus:outline-none focus:border-emerald-500 mb-4 text-emerald-300 font-mono">
        
        <button id="btnEnviar" class="w-full bg-emerald-500 hover:bg-emerald-600 text-slate-950 font-bold p-3 rounded-lg transition duration-200 cursor-pointer">
            Enviar a Go y Analizar
        </button>

        <div id="resultado" class="mt-6 p-4 rounded-lg bg-slate-950 border border-slate-700 hidden font-mono text-sm">
            <!-- Los datos inmutables se pintarán aquí -->
        </div>
    </div>
    <script src="app.js"></script>
</body>
</html>
```

### infrastructure/web/public/app.js (JavaScript con Inmutabilidad)
Aquí el navegador recibe el mapa JSON de Go, lo congela en su RAM para evitar manipulaciones maliciosas y realiza consultas instantáneas en $O(1)$.

```javascript
document.getElementById('btnEnviar').addEventListener('click', async () => {
    const input = document.getElementById('inputNumeros').value;
    const numeros = input.split(',').map(n => parseInt(n.trim())).filter(n => !isNaN(n));

    try {
        const respuesta = await fetch('/api/procesar', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ numeros: numeros })
        });

        const datos = await respuesta.json();

        // 🔒 BLINDAJE INMUTABLE EN EL NAVEGADOR
        // Congelamos el mapa que devolvió Go para evitar alteraciones maliciosas en el cliente
        const mapaVistosInmutable = Object.freeze(datos.mapaVistos);

        const divResultado = document.getElementById('resultado');
        divResultado.classList.remove('hidden');
        
        divResultado.innerHTML = `
            <p class="font-bold ${datos.duplicados ? 'text-red-400' : 'text-emerald-400'}">
                ¿Tiene duplicados?: ${datos.duplicados ? 'SÍ 🚨' : 'NO ✅'}
            </p>
            <div class="mt-3 text-xs text-slate-400">
                <p class="mb-1 text-slate-300 font-bold">Mapa Inmutable en RAM del cliente (O(1)):</p>
                <pre class="bg-slate-900 p-2 rounded text-emerald-500">${JSON.stringify(mapaVistosInmutable, null, 2)}</pre>
            </div>
        `;
    } catch (error) {
        console.error("Error en la petición:", error);
    }
});
```

---

## 🐳 4. El Contenedor: Dockerfile

Para empaquetar de forma eficiente, usamos la estrategia **Multi-stage build**. Compila en una imagen de Go, pero el contenedor de producción final es de apenas **20 MB** porque solo incluye el ejecutable estático y los archivos del Frontend.

```dockerfile
# --- Fase 1: Compilación ---
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
# Compilamos el binario de Go optimizado y estático
RUN CGO_ENABLED=0 GOOS=linux go build -o mi-servidor main.go

# --- Fase 2: Producción ---
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mi-servidor .
COPY --from=builder /app/infrastructure/web/public ./infrastructure/web/public

EXPOSE 8080
CMD ["./mi-servidor"]
```

---

## 🛠️ 5. Cheat Sheet de Consola para Proyectos en Go

Cuando arranques este proyecto en local desde tu terminal de Windows, ejecuta estos comandos en orden:

1.  **Inicializar el sistema de módulos de Go:**
    ```bash
    go mod init mi-proyecto
    ```
2.  **Ejecutar las pruebas de rendimiento (Benchmarks) viendo la RAM:**
    ```bash
    go test -bench=. -benchmem ./...
    ```
3.  **Ver qué variables "escapan" a la RAM pesada (Análisis de Escape):**
    ```bash
    go build -gcflags="-m" main.go
    ```
4.  **Construir y probar el contenedor Docker en tu máquina local:**
    ```bash
    docker build -t mi-app-go .
    docker run -p 8080:8080 mi-app-go
    ```
    *Abre tu navegador en `http://localhost:8080` para ver la aplicación corriendo localmente.*

---

## ☁️ 6. Despliegue en Servidores Gratuitos de Prueba

Existen plataformas en la nube excelentes con planes gratuitos para desarrolladores que leen directamente tu `Dockerfile`.

### Las 3 mejores opciones actuales:
1.  **Render (`render.com`):** La más fácil y visual. Conectas tu GitHub, detecta el `Dockerfile` y te da una URL pública segura. *Nota: Tras 15 minutos de inactividad, entra en suspensión y tarda unos 30 segundos en reactivarse con la siguiente visita [1].*
2.  **Railway (`railway.app`):** Interfaz muy intuitiva y despliegues automáticos a la velocidad de la luz. Te otorga créditos mensuales gratuitos.
3.  **Fly.io (`fly.io`):** Despliega contenedores muy cerca de los usuarios. Se controla desde la terminal instalando su herramienta `flyctl`.

### 🛠️ Flujo paso a paso para desplegar en Render:
*   **Paso 1: Sube tu proyecto a GitHub.** Crea un repositorio y sube todos los archivos (`main.go`, `Dockerfile`, la carpeta `infrastructure`, etc.).
*   **Paso 2: Configura Render.** Regístrate gratis usando tu cuenta de GitHub. En el panel haz clic en **"New +"** y selecciona **"Web Service"**. Conecta el repositorio del proyecto.
