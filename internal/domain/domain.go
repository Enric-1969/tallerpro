/*
Taller mecánico
Don Ramiro tiene un taller con cuatro mecánicos que solo usan el celular,
casi siempre con las manos sucias. Los clientes llaman varias veces al
día para preguntar si su coche está listo, y cada llamada interrumpe el
trabajo.
Cuando al revisar un coche aparece una reparación adicional, el trabajo
se detiene hasta que el cliente aprueba el costo por teléfono, lo que
puede tardar días, y a veces el cliente luego niega haberla aprobado.
Necesita que los clientes conozcan el estado de su coche sin llamar y que
puedan aprobar cotizaciones de forma rápida y con registro.
No necesita control de inventario, que ya se lleva en Excel.
*/

package domain

type Presupuesto struct {
	ID       int     `json:"id"`       //8 bytes
	Total    float64 `json:"total"`    //8 bytes
	Cliente  string  `json:"cliente"`  //16 bytes (8 bytes para el puntero + 8 bytes para la longitud)
	Vehiculo string  `json:"vehiculo"` //16 bytes (8 bytes para el puntero + 8 bytes para la longitud)
	FirmaSHA string  `json:"firmaSha"` //16 bytes (8 bytes para el puntero + 8 bytes para la longitud)
}
