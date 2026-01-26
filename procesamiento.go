package logic

import (
	"AP1/internal/models"
	"strings"
)

// PredicadoLibro es el tipo de función para la lógica funcional
type PredicadoLibro func(models.Libro) bool

// FiltrarLibros aplica Programación Funcional: recibe datos y una función
func FiltrarLibros(libros []models.Libro, criterio PredicadoLibro) []models.Libro {
	var resultado []models.Libro
	for _, l := range libros {
		if criterio(l) {
			resultado = append(resultado, l)
		}
	}
	return resultado
}

// CriterioPorFormato genera una función de filtrado
func CriterioPorFormato(formato string) PredicadoLibro {
	return func(l models.Libro) bool {
		return strings.EqualFold(l.Formato, formato)
	}
}

// CriterioLibrosLargos filtra por número de páginas
func CriterioLibrosLargos() PredicadoLibro {
	return func(l models.Libro) bool {
		return l.Paginas > 400
	}
}
