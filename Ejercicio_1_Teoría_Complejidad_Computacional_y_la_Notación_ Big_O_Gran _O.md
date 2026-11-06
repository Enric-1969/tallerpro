# 🛠️ Ejercicio Práctico de Rendimiento y Estructuras en Go

## ⏱️ 1. El Ejercicio: De Código Ineficiente a Código Pro

Imagina el siguiente problema: te dan una lista de 100.000 números y necesitas saber si hay algún número duplicado.

### 🐢 La Solución Ineficiente (Tiempo: $O(N^2)$ | Memoria: $O(1)$)
El programador principiante mete un bucle dentro de otro para comparar cada número con todos los demás.

```go
// archivo: duplicados.go
package main

// DetectarDuplicadosLento compara todos contra todos.
// Si N = 100.000, ¡hace 10.000.000.000 de operaciones!
func DetectarDuplicadosLento(numeros []int) bool {
	for i := 0; i < len(numeros); i++ {
		for j := i + 1; j < len(numeros); j++ {
			if numeros[i] == numeros[j] {
				return true
			}
		}
	}
	return false
}
```

### 🚀 La Solución Optimizada (Tiempo: $O(N)$ | Memoria: $O(N)$)
Cambiamos la estrategia usando la estructura más rápida de Go: un `map`. Recorremos la lista una sola vez. Si el número ya está en el mapa, encontramos el duplicado al instante gracias al acceso $O(1)$ del mapa.

```go
// DetectarDuplicadosRapido usa un mapa para recordar qué números ya vio.
// Sacrifica un poco de memoria para ganar una velocidad brutal.
func DetectarDuplicadosRapido(numeros []int) bool {
	vistos := make(map[int]bool) // Espacio extra en RAM: O(N)
	for _, num := range numeros { // Un solo bucle: Tiempo O(N)
		if vistos[num] {
			return true
		}
		vistos[num] = true
	}
	return false
}
```

---

## ⏱️ 2. Cómo crear el Test de Rendimiento (Benchmark)

Go incluye una herramienta nativa espectacular para comprobar la velocidad y la memoria exacta que gasta tu código sin necesidad de instalar librerías externas. Para usarla, creamos un archivo de pruebas que termine obligatoriamente en `_test.go`.

```go
// archivo: duplicados_test.go
package main

import (
	"math/rand"
	"testing"
)

// GenerarDatos genera una lista gigante de números aleatorios para la prueba
func GenerarDatos(tamano int) []int {
	datos := make([]int, tamano)
	for i := 0; i < tamano; i++ {
		datos[i] = rand.Intn(tamano * 2)
	}
	return datos
}

// Benchmark de la función lenta
func BenchmarkLento(b *testing.B) {
	datos := GenerarDatos(50000) // Probamos con 50.000 números
	b.ResetTimer()               // Pausamos el cronómetro mientras se generaban los datos
	for i := 0; i < b.N; i++ {
		DetectarDuplicadosLento(datos)
	}
}

// Benchmark de la función rápida
func BenchmarkRapido(b *testing.B) {
	datos := GenerarDatos(50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DetectarDuplicadosRapido(datos)
	}
}
```

---

## 📊 3. Ejecutar y Leer los Resultados

Abre tu terminal en la carpeta donde guardaste los dos archivos anteriores y ejecuta el siguiente comando mágico de Go:

```bash
go test -bench=. -benchmem
```

La consola te arrojará unas líneas parecidas a estas (los números varían según tu procesador):

```text
BenchmarkLento-8           1    1250344100 ns/op          0 B/op          0 allocs/op
BenchmarkRapido-8       2500        421100 ns/op    1422055 B/op        170 allocs/op
```

### 🧠 ¿Cómo interpretar este resultado como un Ingeniero de Software?
*   **`ns/op` (Nanosegundos por operación):** La función lenta tardó 1.250.344.100 ns (¡más de 1,2 segundos!) en procesar la lista una sola vez. La función rápida tardó apenas 421.100 ns (¡0,0004 segundos!). ¡Es casi 3.000 veces más rápida!
*   **`B/op` (Bytes consumidos en memoria RAM):** La función lenta gastó 0 Bytes. Al no crear variables nuevas, su consumo espacial es perfecto ($O(1)$), pero su lentitud no compensa. La función rápida gastó 1.422.055 Bytes (aproximadamente 1.4 Megabytes de RAM). Esto es debido a que tuvo que inicializar el `map` para guardar los registros.
*   **`allocs/op` (Asignaciones de memoria):** Indica cuántas veces el código tuvo que pedirle a la memoria del sistema operativo que reservara espacio. La función rápida lo hizo 170 veces a medida que el mapa iba creciendo en tamaño.

> 💡 **La gran lección de la ingeniería de software:** Casi siempre existe un intercambio (*Trade-off*): para hacer que un programa sea infinitamente más rápido en Tiempo, a menudo tenemos que sacrificar un poco de Espacio (memoria RAM), o viceversa. Poner las cartas sobre la mesa en tu código significa elegir conscientemente qué recurso prefieres priorizar según la máquina en la que corra tu software.

---

## 🔍 4. Análisis profundo bajo el capó de Go

### ¿Qué significa exactamente `datos := make([]int, tamano)`?
En muchos lenguajes (como JavaScript o Python), creas una lista vacía y vas metiendo cosas; el lenguaje se encarga de estirarla en la memoria sobre la marcha. Go no funciona así porque está obsesionado con la velocidad.

Cuando el código necesita meter miles de elementos en una lista (llamada *slice* en Go), ir agrandándola poco a poco obliga a la memoria RAM a hacer mudanzas constantes, lo cual es muy lento. La función `make()` es una orden directa al sistema operativo que dice: *"Oye, sé exactamente cuántas cajas voy a necesitar. Reserva el espacio en la memoria RAM desde este mismo instante"*.

**Desglose paso a paso:**
1.  `datos :=`: Declara una variable nueva llamada datos y le asigna lo que viene a la derecha.
2.  `make()`: Es la función especial de Go para crear e inicializar estructuras vivas en memoria (*slices*, *maps* y *channels*).
3.  `[]int`: Le dice a Go qué tipo de elementos va a guardar la lista. En este caso, números enteros (`int`).
4.  `tamano`: Es la longitud inicial. Si tamano vale 50.000, Go crea inmediatamente en la RAM una fila de 50.000 huecos listos, todos rellenos con el número cero por defecto `[0, 0, 0, ..., 0]`.

### ❓ "En los parámetros solo tenemos 'numeros'"
Fíjate en la diferencia:
*   En la función que resuelve el problema (`DetectarDuplicadosRapido(numeros []int)`), el parámetro que recibe es `numeros`, que es la lista que ya viene del exterior con datos reales.
*   La línea `datos := make([]int, tamano)` pertenece a la función `GenerarDatos`. Esa función no resuelve el problema, solo sirve para fabricar de la nada un saco gigante de números aleatorios para poder hacerle las pruebas de estrés (el Benchmark) a nuestro código.

---

## 💡 5. El Truco: Preasignar la memoria en los Mapas

En la solución rápida original, usamos esta línea para recordar los números que ya habíamos visto:

```go
vistos := make(map[int]bool) // ❌ Inicialización vacía (gasta RAM)
```

Al dejar los paréntesis vacíos, el mapa nace siendo diminuto. A medida que el bucle `for` avanza y mete miles de números, Go se da cuenta de que el mapa no cabe en su sitio original de la RAM. Entonces detiene el programa un milisegundo, busca un espacio más grande, muda todos los datos y borra el espacio viejo. Si procesas 100.000 datos, esa "mudanza" ocurre muchas veces (las llamadas `allocs/op` del Benchmark).

### 🚀 El código con el truco aplicado:
Como la función recibe la lista `numeros`, tú ya sabes cuál es el tamaño máximo que podría llegar a ocupar tu mapa: el largo total de la lista (`len(numeros)`). Si le pasas ese dato como segundo parámetro al `make` del mapa, Go reservará todo el espacio del tirón y eliminará todas las mudanzas de memoria.

```go
func DetectarDuplicadosPro(numeros []int) bool {
	// 🎁 EL TRUCO: Le indicamos al mapa el tamaño estimado desde el inicio
	vistos := make(map[int]bool, len(numeros)) 
	
	for _, num := range numeros {
		if vistos[num] {
			return true
		}
		vistos[num] = true
	}
	return false
}
```

### 📊 ¿Qué pasa en la consola con este truco?
Si ejecutáramos el Benchmark comparando la versión rápida sin truco contra la versión "Pro", la diferencia en la consola es drástica:

```text
BenchmarkRapido-8       2500        421100 ns/op    1422055 B/op        170 allocs/op
BenchmarkPro-8          3800        295000 ns/op     711020 B/op          1 allocs/op
```

*   **`allocs/op` bajó de 170 a 1:** El sistema operativo solo tuvo que otorgar memoria una sola vez. No hubo mudanzas intermedias.
*   **`B/op` (Bytes por operación) cayó a la mitad:** Pasó de gastar 1.4 Megabytes a solo unos 700 Kilobytes de RAM porque no hubo espacio desperdiciado en las expansiones automáticas.
*   **Velocidad:** Al ahorrarse el trabajo de mover datos en la memoria, el código se volvió todavía más rápido.

---

## 🗑️ 6. El Garbage Collector (GC) y la Memoria RAM

El *Garbage Collector* (GC) o Recolector de Basura es el encargado de mantener limpia la memoria RAM de tu servidor mientras tu programa en Go está corriendo.

En lenguajes antiguos como C o C++, el programador tenía que liberar la memoria manualmente con código. Si se olvidaba, el programa se quedaba con la RAM para siempre (un error gravísimo llamado *Memory Leak* o fuga de memoria). Go te quita ese dolor de cabeza y lo hace por ti de forma automática, pero esa comodidad tiene un precio en rendimiento.

### ¿Cómo funciona el Garbage Collector en Go?
Imagina que la memoria RAM de tu aplicación es un lienzo en blanco.
1. **Creación:** Cuando tu código ejecuta una función y creas un mapa con `vistos := make(map[int]bool)`, Go reserva un pedazo de la RAM para ese mapa (esto se llama asignación en el *Heap*).
2. **Fin del uso:** Cuando la función termina y encuentra el duplicado, tu programa ya no necesita ese mapa. El mapa se queda ahí flotando, inservible. Se convierte en "basura".
3. **La limpieza:** Periódicamente, el Garbage Collector se despierta, escanea la memoria RAM, busca qué variables ya nadie está usando y las borra para dejar el espacio libre para el resto del programa.

### 🛑 El gran problema: El impacto "Stop the World"
El recolector de basura de Go es de los más rápidos del mundo, pero para poder limpiar la RAM de forma segura sin romper nada, a veces tiene que congelar o ralentizar por unos microsegundos la ejecución de tu código. Si tu aplicación en Go está procesando miles de peticiones de usuarios por segundo en una tienda online o una aplicación financiera, cada vez que el GC se activa, la aplicación puede sufrir pequeños tirones (latencia).

### 🔄 La conexión con nuestro truco de memoria
¿Te acuerdas de que la función rápida sin optimizar hacía 170 asignaciones de memoria porque el mapa iba creciendo y mudándose de sitio? [10]
* **Sin el truco:** Esas 170 mudanzas dejaban tras de sí trozos de mapas viejos y pequeños esparcidos por la RAM. El Garbage Collector tenía que trabajar el triple para limpiar todo ese rastro.
* **Con el truco (`allocs/op = 1`):** Go reserva un único bloque limpio. Cuando la función termina, solo hay una sola cosa que limpiar. El Garbage Collector entra, borra ese bloque en un instante y se va.

> Al preasignar el tamaño con `make()`, le estás haciendo el trabajo sucio al recolector de basura. Tu código corre más fluido, consume menos procesador y tus servidores estables no sufren picos de lag.
