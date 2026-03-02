package logic

import (
	"AP1/internal/models"
	"strings"
)

type PredicadoLibro func(models.Libro) bool

func Filtrar(libros []models.Libro, p PredicadoLibro) []models.Libro {
	var res []models.Libro
	for _, l := range libros {
		if p(l) { res = append(res, l) }
	}
	return res
}

func EsPDF(l models.Libro) bool { return strings.EqualFold(l.Formato, "PDF") }



