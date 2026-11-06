# 🏗️ Estructuras Fundamentales de Eficiencia en Go y Frontend

Para dominar la eficiencia en Go, no necesitas aprender cientos de herramientas, sino entender a fondo las 4 estructuras de datos fundamentales que el propio lenguaje trae integradas. Cada una fue diseñada específicamente para resolver un problema de Big O (tiempo o memoria) de la forma más rápida posible a nivel de hardware.

---

## 🗺️ 1. Las 4 Estructuras Clave de Go

### 🗂️ Arrays y Slices (La base de la velocidad lineal)
Aunque parecen lo mismo, en Go hay una diferencia técnica crucial que afecta a la memoria RAM.

*   **El Array (Estático):** Una lista con un tamaño fijo que nunca puede cambiar. Ejemplo: `var miArray [5]int`.
    *   *Eficiencia:* Tiempo $O(1)$ para leer o escribir si conoces el índice.
    *   *Secreto de hardware:* Los elementos se guardan en la memoria RAM uno exactamente al lado del otro (de forma contigua). Tu procesador lee esto a la velocidad de la luz porque no tiene que saltar de un punto a otro de la memoria.
*   **El Slice (Dinámico):** Es lo que usamos el 99% del tiempo (`[]int`). Es una ventana flexible que apunta a un array interno.
    *   *Eficiencia:* Tiempo $O(1)$ para añadir elementos al final (`append`), siempre y cuando tenga espacio. Si se llena, ocurre la "mudanza en la RAM" para duplicar su tamaño de forma automática.

### 🗺️ Maps (Tablas Hash: El rey del tiempo constante)
El `map` es la estructura que usamos en nuestro ejercicio para ganarle a los bucles anidados.
*   **¿Qué es?:** Una estructura de Clave-Valor (`map[string]int`).
*   **Eficiencia:** Tiempo $O(1)$ tanto para guardar un dato, como para buscarlo o borrarlo. Da igual si el mapa tiene 10 datos o 10 millones; tarda exactamente un paso.
*   **¿Cómo logra esa magia?:** Cuando le pides buscar la clave `"usuario_100"`, Go no lee el mapa desde el principio. Pasa esa palabra por una función matemática interna (función *Hash*) que calcula instantáneamente la posición exacta de la memoria RAM donde está guardada. Va directo al grano.

### 🏗️ Structs (Estructuras personalizadas: Cero desperdicio)
En Go no existen las "Clases" de la programación tradicional orientada a objetos; usamos `struct`.
*   **¿Qué es?:** Un molde para agrupar diferentes tipos de datos bajo un mismo nombre.
    ```go
    type Usuario struct {
        ID     int    // Gasta 8 bytes
        Edad   int8   // Gasta 1 byte
        Nombre string // Gasta 16 bytes
    }
    ```
*   **Eficiencia:** Uso de memoria ultra optimizado. El compilador de Go calcula exactamente cuántos bytes va a pesar cada variable del struct.
*   **Secreto de ingeniería:** Go empaqueta estos datos de forma tan compacta que muchas veces se quedan en el *Stack* (una zona de la memoria RAM que se limpia sola inmediatamente al terminar la función), evitando por completo tener que cansar al *Garbage Collector*.

### 🔀 Channels (Canales: Eficiencia en Concurrencia)
Go es famoso mundialmente por las *Goroutines* (funciones que se ejecutan simultáneamente en paralelo aprovechando todos los núcleos de tu procesador). Los canales son las tuberías que las conectan.
*   **¿Qué es?:** Un conducto seguro para enviar datos de una Goroutine a otra (`chan int`).
*   **Eficiencia:** En otros lenguajes, para compartir datos entre hilos en paralelo, tienes que bloquear la memoria manualmente (usando *Locks* o *Mutex*) para que dos procesos no pisen el mismo dato a la vez, lo cual genera colas y lentitud. Los canales de Go gestionan este tráfico de datos a nivel interno de forma matemática y optimizada, permitiendo procesar millones de operaciones concurrentes con un consumo de RAM ridículo (apenas 2 KB por Goroutine).

> 💡 **El resumen definitivo para tu día a día:**
> 1. Si necesitas almacenar una lista ordenada y procesarla en orden: Usa un **Slice** (y si sabes el tamaño, créalo con `make`).
> 2. Si necesitas buscar elementos constantemente por un identificador único: Usa un **Map** (evitarás los bucles anidados $O(N^2)$).
> 3. Si necesitas modelar datos complejos de tu negocio: Usa un **Struct**.
> 4. Si vas a hacer tareas en paralelo: Comunícalas mediante **Channels**.

---

## 🔒 2. Diseño del Frontend y Flujo de Datos API

En Go, ni los slices ni los maps son inmutables por defecto. Si un código modifica un slice o un map, los datos cambian en la memoria. Sin embargo, para alimentar un frontend mediante una API (por ejemplo, enviando un JSON), la estructura que elijas en el backend depende 100% de **cómo va a consumir los datos el frontend** y de la eficiencia Big O.

### 🗂️ Cuándo usar un Slice (`[]Estructura`) para el Frontend
Usa un slice (una lista ordenada) cuando el frontend necesite mostrar la información en un orden específico o cuando tenga que recorrer todos los elementos a la vez.
*   **Ejemplos en el Frontend:** Un *feed* de noticias (ordenado por fecha), una tabla de productos con paginación o un ranking.
*   **Complejidad en el Frontend:** El frontend recibe un Array de JSON `[{...}, {...}]`. Recorrer un array para pintarlo en la pantalla es una operación $O(N)$ muy natural para el navegador.
*   **Desventaja de seguridad:** Si el frontend recibe un slice gigante y necesita buscar un ID, hará un bucle en JavaScript. Un atacante podría enviar peticiones maliciosas que fuercen al navegador a realizar millones de operaciones, provocando un ataque tipo *Client-side DoS* (congelar la pestaña).

### 🗺️ Cuándo usar un Map (`map[Key]Estructura`) para el Frontend
Usa un map (diccionario clave-valor) cuando el frontend no necesite un orden, sino acceder directamente a un elemento específico de forma instantánea usando un ID.
*   **Ejemplos en el Frontend:** Un catálogo de productos donde haces clic en un ID específico o un sistema de traducciones (`traducciones["boton_guardar"]`).
*   **Complejidad y Seguridad:** El frontend recibe un Objeto JSON `{"id_123": {...}}`. Para renderizarlo accede directamente en tiempo $O(1)$ sin bucles, blindando al cliente contra ataques de ineficiencia.

### ⚖️ Tabla de Decisión Rápida

| Si el Frontend necesita... | Estructura en Go | Complejidad | Formato JSON resultante |
| :--- | :--- | :--- | :--- |
| **Mostrar un orden, fechas o listas paginadas** | Slice (Lista) | $O(N)$ | `[{"id": 1, "nombre": "A"}, ...]` |
| **Buscar un elemento por su ID directamente** | Map (Clave-Valor) | $O(1)$ | `{"1": {"nombre": "A"}, "2": ...}` |

> 🔒 **Un truco de arquitectura Pro:** En aplicaciones profesionales, el backend en Go suele almacenar los datos en *Maps* internamente para que las búsquedas en el servidor sean $O(1)$. Luego, justo antes de enviar la respuesta, Go transforma ese mapa en un *Slice* ordenado si el diseño de la pantalla lo requiere. Así tienes lo mejor de los dos mundos.

---

## 🏗️ 3. Inmutabilidad en JavaScript y Reparto de Carga

Hoy en día, el backend delega gran parte del peso de la renderización y la memoria al cliente (el navegador). El servidor en Go solo procesa los datos limpios de forma ultraeficiente, y la máquina del usuario se encarga de estructurarlos y pintarlos en pantalla.

### 🛑 ¿Con qué estructura inmutable llenar la memoria en JavaScript?
Para blindar la información contra manipulaciones y ciberataques en el navegador, tienes dos caminos:

#### Opción A: Si Go te envió un Objeto / Mapa ➔ `Object.freeze()`
Si necesitas acceso instantáneo por ID ($O(1)$) y máxima seguridad, usas un objeto plano de JavaScript y lo "congelas":

```javascript
// El JSON que llegó desde tu backend en Go
const datosDelBackend = {
    "usr_99": { nombre: "Carlos", rol: "admin" },
    "usr_100": { nombre: "Ana", rol: "user" }
};

// 🔒 TRUCO PRO: Congelar la estructura en el navegador
const usuariosSeguros = Object.freeze(datosDelBackend);

// 🛡️ Intento de ataque o modificación maliciosa:
usuariosSeguros["usr_99"].rol = "hacker"; // ❌ Da error o se ignora en "strict mode"
```
`Object.freeze()` hace que el objeto sea 100% inmutable: no se pueden añadir propiedades, ni borrar, ni modificar las existentes.

#### Opción B: Si necesitas una Lista Ordenada ➔ `Map` de JS + Métodos no mutables
Si necesitas recorrer datos en orden, la regla de oro en el frontend es **nunca usar métodos mutables como `push()` o `splice()`**. Se crean nuevas copias usando el operador de propagación (`...` o *spread operator*):

```javascript
// Crear una lista inmutable
const listaOriginal = Object.freeze([{ id: 1, nombre: "Producto A" }]);

// Si el usuario añade un producto, NO hacemos push. Creamos una lista NUEVA:
const nuevaLista = [...listaOriginal, { id: 2, nombre: "Producto B" }];
```

### 💻 El reparto de tareas actual (Servidor vs Cliente)
*   **El Servidor (Go):** Se encarga del trabajo pesado de lógica de negocio y seguridad. Al usar estructuras optimizadas y preasignación con `make()`, atiende a millones de usuarios simultáneamente consumiendo poquísima RAM, despachando en milisegundos.
*   **El Cliente (Navegador/Android):** Descarga los datos crudos y asume el coste de memoria RAM para mantener el estado de la aplicación, ordenar tablas y dibujar la interfaz.

### ⚠️ El peligro del exceso de memoria en el cliente
Liberar al servidor es genial, pero no debes saturar al cliente. Si tu backend envía un *slice* con 50.000 registros de golpe a un móvil de gama baja, el motor de JavaScript se quedará sin RAM y la pestaña sufrirá un *Crash*. Por eso se utiliza la **paginación** (entregar bloques de 50 en 50), manteniendo feliz al servidor y seguro al navegador.

---

## 🔍 4. Auditoría de los Comentarios de Memoria (Arquitectura de 64 bits)

### 1. En la estructura Presupuesto (`domain.go`)
Los cálculos sobre las variables individuales son matemáticamente perfectos en una máquina de 64 bits:
*   **`ID` (`int`):** 8 bytes.
*   **`Total` (`float64`):** 8 bytes.
*   **`Cliente`, `Vehiculo`, `FirmaSHA` (`string`):** 16 bytes cada uno (8 bytes para el puntero al array de bytes + 8 bytes para la longitud).

> ⚡ **El secreto de hardware:** Sumando todos los campos, tu estructura pesa exactamente **64 bytes** (8 + 8 + 16 + 16 + 16). Al pesar exactamente 64 bytes, un registro entero **cabe perfectamente en una sola línea de caché L1/L2 del procesador de tu máquina Windows**. Cuando Go procese un presupuesto, la CPU lo leerá en un único ciclo de reloj sin necesidad de dar saltos de memoria. ¡Alineación y rendimiento de nivel experto!

### 2. En el struct PresupuestoUsecase (`usecase.go`)
*   **Matiz técnico de Go:** Como tu struct actual está vacío (`type PresupuestoUsecase struct{}`), esto se conoce en Go como una **estructura vacía**. Las estructuras vacías ocupan **exactamente 0 bytes** en la memoria RAM. 
*   Go es tan inteligente que si creas 1.000 instancias de un struct vacío, todas apuntarán a la misma dirección de memoria base del sistema para ahorrar recursos. Ocupará espacio real en la RAM únicamente cuando le inyectemos dependencias operativas (como el puntero de conexión de la base de datos de SQLite).
