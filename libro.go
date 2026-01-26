package models

// Libro representa la estructura de un e-book
type Libro struct {
	ID      int
	Titulo  string // Usaremos nombres de mangas
	Autor   string
	Paginas int
	Formato string
	Estado  string
}

// ObtenerInventario carga los libros iniciales
func ObtenerInventario() []Libro {
	return []Libro{
		{1, "Berserk: Deluxe Edition", "Kentaro Miura", 700, "PDF", "Leído"},
		{2, "One Piece: Saga de Wano", "Eiichiro Oda", 200, "EPUB", "Pendiente"},
		{3, "Oyasumi Punpun: Vol 1", "Inio Asano", 150, "PDF", "Descargado"},
		{4, "Vagabond: VizBig Edition", "Takehiko Inoue", 600, "MOBI", "Leído"},
		{5, "Akira: 35th Anniversary", "Katsuhiro Otomo", 350, "PDF", "Pendiente"},
	}
}
