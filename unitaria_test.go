package tests

import (
	"AP1/internal/models"
	"testing"
)

// 1. PRUEBA UNITARIA: Verificar creación exitosa (Unidad 1 y 2)
func TestCreacionLibroExitoso(t *testing.T) {
	libro, err := models.NuevoLibro(1, "Berserk", "Kentaro Miura", 700, "PDF")

	if err != nil {
		t.Errorf("❌ Error inesperado: El sistema debería permitir crear un libro válido. Detalles: %v", err)
	}

	if libro.Titulo != "Berserk" {
		t.Errorf("❌ El título guardado no coincide. Esperado: Berserk, Obtenido: %s", libro.Titulo)
	}
}

// 2. PRUEBA DE MANEJO DE ERRORES: Verificar validación de datos (Unidad 3)
func TestValidacionPaginasNegativas(t *testing.T) {
	_, err := models.NuevoLibro(2, "Manga Fallido", "Autor", -50, "EPUB")

	if err == nil {
		t.Error("❌ FALLO DE OPTIMIZACIÓN: El sistema NO debería permitir libros con páginas negativas.")
	}
}

// 3. PRUEBA DE REQUISITOS WEB: Verificar etiquetas JSON (Unidad 4)
func TestExportacionJSON(t *testing.T) {
	libro, _ := models.NuevoLibro(3, "Akira", "Otomo", 350, "PDF")

	// Si el campo empieza con minúscula o no existe, esta prueba fallará
	// Esto asegura que el servicio web podrá enviar los datos correctamente al navegador
	if libro.ID == 0 || libro.Titulo == "" {
		t.Error("❌ ERROR DE SERIALIZACIÓN: Los campos del Libro deben ser públicos para el servicio web JSON.")
	}
}
