package repository

import (
	"log"
	"tallermecanico/internal/domain"

	"github.com/glebarez/sqlite" // 👈 Usamos SQLite puro sin CGO
	"gorm.io/gorm"
)

// InitDB inicializa la conexión con SQLite y ejecuta las migraciones automáticas
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("taller.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Auto-migración de las tablas del dominio
	db.AutoMigrate(
		&domain.Cliente{},
		&domain.Vehiculo{},
		&domain.Presupuesto{},
	)

	return db
}
