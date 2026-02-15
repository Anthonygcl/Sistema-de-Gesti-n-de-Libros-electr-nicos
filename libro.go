package models

import (
	"errors"
	"fmt"
)

// Libro con campos privados (Encapsulación)
type Libro struct {
	id      int
	titulo  string
	autor   string
	paginas int
	formato string
}

// NuevoLibro es el constructor con manejo de errores
func NuevoLibro(id int, titulo string, autor string, paginas int, formato string) (*Libro, error) {
	if paginas <= 0 {
		return nil, errors.New("el libro debe tener al menos 1 página")
	}
	if titulo == "" {
		return nil, errors.New("el título no puede estar vacío")
	}
	return &Libro{id, titulo, autor, paginas, formato}, nil
}

// MÉTODOS GETTERS (Para que otros paquetes puedan leer los datos)
func (l *Libro) Titulo() string  { return l.titulo }
func (l *Libro) Autor() string   { return l.autor }
func (l *Libro) Paginas() int    { return l.paginas }
func (l *Libro) Formato() string { return l.formato }

// ObtenerDetalles implementa la interfaz Gestionable
func (l *Libro) ObtenerDetalles() string {
	return fmt.Sprintf("[%s] %s - %d pág. (%s)", l.autor, l.titulo, l.paginas, l.formato)
}


