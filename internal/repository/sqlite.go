package repository

import (
	"errors"
	"log"
	"os"
	"tallermecanico/internal/domain"
	"time"

	"github.com/glebarez/sqlite" // 👈 Conserva SIEMPRE tu importación original intacta
	"gorm.io/gorm"
)

// InitDB inicializa la conexión devolviendo el error en lugar de matar la app
func InitDB() (*gorm.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "taller.db"
	}

	var db *gorm.DB
	var err error

	maxReintentos := 3
	espera := 500 * time.Millisecond // Reintento al medio segundo

	// Bucle de reintentos robusto
	for i := 0; i < maxReintentos; i++ {
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err == nil {
			// Conexión exitosa, rompemos el bucle
			break
		}

		log.Printf("[⚠️ BBDD] Intento %d de %d fallido: %v. Reintentando en %v...", i+1, maxReintentos, err, espera)
		time.Sleep(espera)

		// Estrategia Senior (Backoff): Duplicamos el tiempo de espera para el siguiente intento
		espera *= 2
	}

	// Si tras los reintentos sigue fallando, devolvemos el error limpiamente
	if err != nil {
		return nil, errors.New("no se pudo establecer conexión con la base de datos local tras varios intentos")
	}

	// Configuración de optimización de concurrencia interna de SQLite
	sqlDB, errDB := db.DB()
	if errDB == nil {
		// Evita que SQLite se bloquee limitando el pool a 1 conexión de escritura abierta activa
		sqlDB.SetMaxOpenConns(1)
	}

	// Auto-migración
	err = db.AutoMigrate(
		&domain.Cliente{},
		&domain.Vehiculo{},
		&domain.Presupuesto{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
