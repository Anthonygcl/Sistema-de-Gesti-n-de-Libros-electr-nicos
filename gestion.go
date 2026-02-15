package interfaces

// Gestionable define el comportamiento que cualquier item de la biblioteca debe tener
type Gestionable interface {
	ObtenerDetalles() string // Cualquier struct con este método es un "Gestionable"
}
