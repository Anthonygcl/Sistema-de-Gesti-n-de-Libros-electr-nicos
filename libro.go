package models

import (
	"errors"
	"fmt"
)

// Libro: Campos públicos para JSON.
type Libro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Paginas int    `json:"paginas"`
	Formato string `json:"formato"`
}

// NuevoLibro: Mantiene la validación de la Unidad 3
func NuevoLibro(id int, titulo string, autor string, paginas int, formato string) (*Libro, error) {
	if paginas <= 0 {
		return nil, errors.New("las páginas deben ser un número positivo")
	}
	if titulo == "" {
		return nil, errors.New("el título es obligatorio")
	}
	return &Libro{
		ID:      id,
		Titulo:  titulo,
		Autor:   autor,
		Paginas: paginas,
		Formato: formato,
	}, nil
}

func (l *Libro) ObtenerDetalles() string {
	return fmt.Sprintf("[%s] %s - %d pág.", l.Autor, l.Titulo, l.Paginas)
}


