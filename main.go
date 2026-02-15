package main

import (
	"AP1/internal/interfaces"
	"AP1/internal/logic"
	"AP1/internal/models"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// 1. INICIALIZACIÓN (POO)
	miBiblioteca := logic.Biblioteca{}
	tituloColor := color.New(color.FgHiMagenta).Add(color.Bold)
	errorColor := color.New(color.FgRed)
	infoColor := color.New(color.FgHiCyan)

	// 2. CARGA DE DATOS CON MANEJO DE ERRORES (Unidad 3)
	// Intentamos cargar libros. Si hay error (como páginas negativas), lo avisamos.
	l1, _ := models.NuevoLibro(1, "Berserk", "Kentaro Miura", 700, "PDF")
	l2, _ := models.NuevoLibro(2, "Vagabond", "Takehiko Inoue", 600, "EPUB")
	l3, _ := models.NuevoLibro(3, "One Piece", "Eiichiro Oda", 1000, "PDF")
	l4, _ := models.NuevoLibro(4, "Akira", "Katsuhiro Otomo", 350, "PDF")

	// Ejemplo de error detectado por el sistema
	_, err := models.NuevoLibro(5, "Error Book", "Desconocido", -10, "MOBI")
	if err != nil {
		errorColor.Printf("ALERTA DE SISTEMA: %v\n", err)
	}

	miBiblioteca.AgregarItem(l1)
	miBiblioteca.AgregarItem(l2)
	miBiblioteca.AgregarItem(l3)
	miBiblioteca.AgregarItem(l4)

	// 3. MENÚ INTERACTIVO (Continuidad del Proyecto)
	for {
		tituloColor.Println("\n--- ZENITH E-BOOK MANAGER V2 (SISTEMA INTERACTIVO) ---")
		fmt.Println("1. Ver catálogo completo")
		fmt.Println("2. Filtrar libros en formato PDF")
		fmt.Println("3. Libros de gran extensión (>400 páginas)")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			infoColor.Println("\n[ CATÁLOGO COMPLETO ]")
			mostrarTabla(miBiblioteca.Items)
		case 2:
			infoColor.Println("\n[ FILTRANDO FORMATO: PDF ]")
			var filtrados []interfaces.Gestionable
			for _, item := range miBiblioteca.Items {
				if libro, ok := item.(*models.Libro); ok {
					if libro.Formato() == "PDF" {
						filtrados = append(filtrados, item)
					}
				}
			}
			mostrarTabla(filtrados)
		case 3:
			infoColor.Println("\n[ LIBROS DE GRAN EXTENSIÓN ]")
			var largos []interfaces.Gestionable
			for _, item := range miBiblioteca.Items {
				if libro, ok := item.(*models.Libro); ok {
					if libro.Paginas() > 400 {
						largos = append(largos, item)
					}
				}
			}
			mostrarTabla(largos)
		case 4:
			color.Green("Saliendo del sistema... ¡Sayonara!")
			return
		default:
			errorColor.Println("Opción no válida, intente de nuevo.")
		}
	}
}

// Función auxiliar para mostrar tablas (Uso de Interfaces)
func mostrarTabla(items []interfaces.Gestionable) {
	if len(items) == 0 {
		color.Yellow("No hay libros para mostrar con ese criterio.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Título", "Autor", "Formato", "Páginas"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)

	for _, item := range items {
		// Polimorfismo: el sistema trata todo como 'interfaces.Gestionable'
		// pero aquí obtenemos los datos específicos del Libro para la tabla
		if libro, ok := item.(*models.Libro); ok {
			table.Append([]string{
				libro.Titulo(),
				libro.Autor(),
				libro.Formato(),
				fmt.Sprintf("%d", libro.Paginas()),
			})
		}
	}
	table.Render()
}

