package usecase

import (
	"crypto/sha256"
	"fmt"
	"tallermecanico/internal/domain"
)

// PresupuestoUsecase contiene las dependencias (por ahora pura lógica)
type PresupuestoUsecase struct{} //0 bytes (por estar vacia)

// NewPresupuestoUsecase inicializa el caso de uso
func NewPresupuestoUsecase() *PresupuestoUsecase {
	return &PresupuestoUsecase{}
}

// FirmarPresupuesto genera un hash criptográfico SHA-256 de forma ultraeficiente
func (u *PresupuestoUsecase) FirmarPresupuesto(p *domain.Presupuesto) string {
	datosParaFirmar := fmt.Sprintf("%d-%s-%s-%.2f", p.ID, p.Cliente, p.Vehiculo, p.Total)

	// Criptografía nativa en Go (Seguridad y velocidad en el Stack de memoria)
	hash := sha256.Sum256([]byte(datosParaFirmar))

	return fmt.Sprintf("%x", hash)
}

// BuscarDuplicados usa tu mapa optimizado O(N) para evitar fraudes en la lista de presupuestos
func (u *PresupuestoUsecase) BuscarDuplicados(lista []domain.Presupuesto) bool {
	// 🎁 Aplicamos tu truco de optimización: preasignamos el mapa en la RAM
	vistos := make(map[int]bool, len(lista))

	for _, p := range lista {
		if vistos[p.ID] {
			return true // Encontró un ID duplicado (Alerta de fraude)
		}
		vistos[p.ID] = true
	}
	return false
}
