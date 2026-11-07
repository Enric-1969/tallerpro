package usecase

import (
	"crypto/sha256"
	"fmt"
	"tallermecanico/internal/domain"
)

type PresupuestoUsecase struct{}

func NewPresupuestoUsecase() *PresupuestoUsecase {
	return &PresupuestoUsecase{}
}

// FirmarPresupuesto genera un hash criptográfico SHA-256 adaptado al nuevo modelo
func (u *PresupuestoUsecase) FirmarPresupuesto(p *domain.Presupuesto) string {
	// Usamos la matrícula del vehículo o datos disponibles en el struct actualizado
	datosParaFirmar := fmt.Sprintf("%d-%s-%.2f", p.ID, p.Vehiculo.Matricula, p.Total)

	hash := sha256.Sum256([]byte(datosParaFirmar))
	return fmt.Sprintf("%x", hash)
}

// BuscarDuplicados ajustado al ID de tipo uint de gorm.Model
func (u *PresupuestoUsecase) BuscarDuplicados(lista []domain.Presupuesto) bool {
	// Cambiamos la clave del mapa a uint para que coincida con gorm.Model
	vistos := make(map[uint]bool, len(lista))

	for _, p := range lista {
		if vistos[p.ID] {
			return true // ID duplicado detectado
		}
		vistos[p.ID] = true
	}
	return false
}
