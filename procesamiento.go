package logic

import (
	"AP1/internal/models"
	"strings"
)

type PredicadoLibro func(models.Libro) bool

func FiltrarLibros(libros []models.Libro, criterio PredicadoLibro) []models.Libro {
	var resultado []models.Libro
	for _, l := range libros {
		if criterio(l) {
			resultado = append(resultado, l)
		}
	}
	return resultado
}

// CORRECCIÓN: Usamos l.Formato (el campo) y la variable correcta 'l'
func CriterioPorFormato(formato string) PredicadoLibro {
	return func(l models.Libro) bool {
		return strings.EqualFold(l.Formato, formato)
	}
}

// CORRECCIÓN: Usamos l.Paginas y la variable correcta 'l'
func CriterioLibrosLargos() PredicadoLibro {
	return func(l models.Libro) bool {
		return l.Paginas > 400
	}
}
