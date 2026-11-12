package usecase

import (
	"tallermecanico/internal/domain"
	"testing"
)

func TestFirmarPresupuesto(t *testing.T) {
	uc := NewPresupuestoUsecase()

	presupuesto := &domain.Presupuesto{
		ID:    101,
		Total: 250.50,
		Vehiculo: &domain.Vehiculo{
			Matricula: "1234-ABC",
		},
	}

	firma1 := uc.FirmarPresupuesto(presupuesto)
	firma2 := uc.FirmarPresupuesto(presupuesto)

	// 1. La firma SHA-256 en hexadecimal siempre debe medir 64 caracteres
	if len(firma1) != 64 {
		t.Fatalf("Se esperaba un hash SHA-256 de 64 caracteres, se obtuvo %d", len(firma1))
	}

	// 2. Determinismo: mismos datos deben producir exactamente la misma firma
	if firma1 != firma2 {
		t.Errorf("La firma debería ser determinista: '%s' != '%s'", firma1, firma2)
	}

	// 3. Integridad: si cambia el importe en 1 céntimo, la firma debe cambiar por completo
	presupuesto.Total = 250.51
	firmaAlterada := uc.FirmarPresupuesto(presupuesto)
	if firma1 == firmaAlterada {
		t.Error("¡Fallo de seguridad! El presupuesto fue alterado pero la firma SHA-256 no cambió")
	}
}

func TestBuscarDuplicados(t *testing.T) {
	uc := NewPresupuestoUsecase()

	// Caso sin duplicados
	listaLimpia := []domain.Presupuesto{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	if uc.BuscarDuplicados(listaLimpia) {
		t.Error("No debería detectar duplicados en una lista con IDs únicos")
	}

	// Caso con ID duplicado
	listaConDuplicado := []domain.Presupuesto{
		{ID: 1},
		{ID: 2},
		{ID: 1},
	}
	if !uc.BuscarDuplicados(listaConDuplicado) {
		t.Error("Debería haber detectado el ID 1 duplicado")
	}
}
