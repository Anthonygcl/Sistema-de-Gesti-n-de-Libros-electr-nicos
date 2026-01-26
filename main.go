package main

import (
	"AP1/internal/logic"
	"AP1/internal/models"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// 1. Cargar datos
	biblioteca := models.ObtenerInventario()

	// 2. Estilos visuales
	tituloColor := color.New(color.FgHiMagenta).Add(color.Bold)
	errorColor := color.New(color.FgRed)

	for {
		tituloColor.Println("\n--- E-BOOK MANAGER ---")
		fmt.Println("1. Ver catálogo completo")
		fmt.Println("2. Filtrar libros PDF")
		fmt.Println("3. Ver libros de más de 400 páginas")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opción: ")

		var opc int
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			imprimirTabla(biblioteca)
		case 2:
			// Uso de programación funcional
			pdfs := logic.FiltrarLibros(biblioteca, logic.CriterioPorFormato("PDF"))
			imprimirTabla(pdfs)
		case 3:
			largos := logic.FiltrarLibros(biblioteca, logic.CriterioLibrosLargos())
			imprimirTabla(largos)
		case 4:
			color.Cyan("¡Gracias por usar E-BOOK Manager!")
			return
		default:
			errorColor.Println("Error: Opción no válida.")
		}
	}
}

func imprimirTabla(libros []models.Libro) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Título del Manga", "Autor", "Páginas", "Formato", "Estado"})

	// Estética de la tabla
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)

	for _, l := range libros {
		table.Append([]string{
			fmt.Sprintf("%d", l.ID),
			l.Titulo,
			l.Autor,
			fmt.Sprintf("%d", l.Paginas),
			l.Formato,
			l.Estado,
		})
	}
	table.Render()
}
