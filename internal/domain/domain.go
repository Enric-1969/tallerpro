package domain

import (
	"time"
)

// Cliente representa a particulares o autónomos sin sobrecarga de ORM oculta.
type Cliente struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Nombre      string     `json:"nombre" gorm:"not null"`
	Telefono    string     `json:"telefono"`
	CifDni      string     `json:"cifDni" gorm:"unique;not null"`
	RazonSocial string     `json:"razonSocial"`
	Vehiculos   []Vehiculo `json:"vehiculos" gorm:"foreignKey:ClienteID"`
	CreatedAt   time.Time  `json:"createdAt"`
}

// Vehiculo representa el automóvil vinculado al cliente.
type Vehiculo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ClienteID uint   `json:"clienteId" gorm:"index;not null"`
	Matricula string `json:"matricula" gorm:"unique;not null"`
	Bastidor  string `json:"bastidor" gorm:"unique;not null"`
	Marca     string `json:"marca"`
	Modelo    string `json:"modelo"`
}

// Presupuesto gestiona la orden de reparación, su estado inmutable y firma criptográfica.
type Presupuesto struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	VehiculoID  uint      `json:"vehiculoId" gorm:"index;not null"`
	Vehiculo    *Vehiculo `json:"vehiculo,omitempty" gorm:"foreignKey:VehiculoID"`
	Descripcion string    `json:"descripcion" gorm:"type:text"`
	Total       float64   `json:"total" gorm:"not null"`
	Estado      string    `json:"estado" gorm:"not null"` // "Pendiente", "Aprobado", "Rechazado", "BloqueadoFiscal"
	FirmaSHA    string    `json:"firmaSha"`
	CreatedAt   time.Time `json:"createdAt"`
}
