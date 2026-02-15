package logic

import (
	"AP1/internal/interfaces"
	"errors"
)

// Biblioteca actúa como nuestra "Clase" administradora
type Biblioteca struct {
	Items []interfaces.Gestionable
}

// AgregarItem usa Polimorfismo: recibe cualquier objeto que sea Gestionable
func (b *Biblioteca) AgregarItem(item interfaces.Gestionable) error {
	if item == nil {
		return errors.New("no se puede agregar un item vacío")
	}
	b.Items = append(b.Items, item)
	return nil
}
