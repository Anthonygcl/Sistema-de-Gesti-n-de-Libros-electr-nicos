package logic

import (
	"AP1/internal/interfaces"
	"sync"
)

type Biblioteca struct {
	mu    sync.RWMutex // Mutex para seguridad en Servicios Web
	Items []interfaces.Gestionable
}

func (b *Biblioteca) AgregarItem(item interfaces.Gestionable) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Items = append(b.Items, item)
}

func (b *Biblioteca) ObtenerItems() []interfaces.Gestionable {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Items
}
