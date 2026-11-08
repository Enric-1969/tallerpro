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

import (
	"gorm.io/gorm"
)

// 👤 Modelo de Cliente (Persistencia SQLite)
type Cliente struct {
	gorm.Model            // ID (uint/8 bytes)
	Nombre     string     `json:"nombre" gorm:"not null"` // 16 bytes (puntero + longitud)
	Telefono   string     `json:"telefono"`               // 16 bytes (puntero + longitud)
	Vehiculos  []Vehiculo `json:"vehiculos"`              // Slice de vehículos asociados
}

// 🚗 Modelo de Vehículo (Persistencia SQLite)
type Vehiculo struct {
	gorm.Model        // ID (uint/8 bytes)
	ClienteID  uint   `json:"clienteId"`                        // 8 bytes (Relación con el Cliente)
	Matricula  string `json:"matricula" gorm:"unique;not null"` // 16 bytes (puntero + longitud)
	Bastidor   string `json:"bastidor" gorm:"unique;not null"`  // 16 bytes (puntero + longitud)
	Marca      string `json:"marca"`                            // 16 bytes (puntero + longitud)
	Modelo     string `json:"modelo"`                           // 16 bytes (puntero + longitud)
}

// 📋 Orden de Reparación / Presupuesto (Persistencia SQLite + Negocio)
type Presupuesto struct {
	gorm.Model           // ID (uint/8 bytes), CreatedAt, UpdatedAt, DeletedAt
	VehiculoID  uint     `json:"vehiculoId"`  // 8 bytes (Relación con el vehículo)
	Vehiculo    Vehiculo `json:"vehiculo"`    // Estructura embebida / auto-carga con Preload
	Descripcion string   `json:"descripcion"` // 16 bytes (8 bytes para el puntero + 8 bytes para la longitud)
	Total       float64  `json:"total"`       // 8 bytes
	Estado      string   `json:"estado"`      // 16 bytes (8 bytes para el puntero + 8 bytes para la longitud) ("Pendiente", "Aprobado", "Rechazado")
	FirmaSHA    string   `json:"firmaSha"`    // 16 bytes (8 bytes para el puntero + 8 bytes para la longitud) - Firma criptográfica
}
