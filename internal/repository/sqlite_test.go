package repository

import (
	"tallermecanico/internal/domain"
	"testing"
)

func TestInitDB_YPersistencia(t *testing.T) {
	// Usamos SQLite en memoria RAM para no ensuciar el disco duro con archivos .db
	t.Setenv("DB_PATH", ":memory:")

	// 1. Probamos que InitDB conecta, configura el pool y ejecuta AutoMigrate sin errores
	db, err := InitDB()
	if err != nil {
		t.Fatalf("InitDB falló inesperadamente: %v", err)
	}

	// 2. Probamos que las tablas migradas funcionan guardando y leyendo un Cliente real
	cliente := domain.Cliente{
		Nombre:   "Taller Mecánico Prueba S.L.",
		Telefono: "600123456",
		CifDni:   "B12345678",
	}

	if err := db.Create(&cliente).Error; err != nil {
		t.Fatalf("Error al insertar cliente de prueba en SQLite: %v", err)
	}

	var recuperado domain.Cliente
	if err := db.First(&recuperado, cliente.ID).Error; err != nil {
		t.Fatalf("Error al recuperar el cliente insertado: %v", err)
	}

	if recuperado.CifDni != "B12345678" {
		t.Errorf("Se esperaba CIF 'B12345678', se obtuvo '%s'", recuperado.CifDni)
	}
}
