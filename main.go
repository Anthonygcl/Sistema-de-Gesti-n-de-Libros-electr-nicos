package main

import (
	"AP1/internal/interfaces"
	"AP1/internal/logic"
	"AP1/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/fatih/color"
)

var repo = logic.Biblioteca{}

func main() {
	// 1. CARGA INICIAL
	l1, _ := models.NuevoLibro(1, "Berserk", "Kentaro Miura", 700, "PDF")
	l2, _ := models.NuevoLibro(2, "Vagabond", "Takehiko Inoue", 600, "EPUB")
	l3, _ := models.NuevoLibro(3, "One Piece", "Eiichiro Oda", 1000, "PDF")
	l4, _ := models.NuevoLibro(4, "Akira", "Katsuhiro Otomo", 350, "PDF")

	repo.AgregarItem(l1)
	repo.AgregarItem(l2)
	repo.AgregarItem(l3)
	repo.AgregarItem(l4)

	// 2. REGISTRO DE LOS 8 SERVICIOS WEB
	http.HandleFunc("/api/libros", listarTodos)                 // 1
	http.HandleFunc("/api/libros/pdf", listarPDFs)              // 2
	http.HandleFunc("/api/libros/largos", listarLargos)         // 3
	http.HandleFunc("/api/stats/paginas", calcularTotalPaginas) // 4 (CONCURRENCIA)
	http.HandleFunc("/api/info", infoSistema)                   // 5
	http.HandleFunc("/api/health", healthCheck)                 // 6
	http.HandleFunc("/api/autores", listarAutores)              // 7
	http.HandleFunc("/api/ultimo", verUltimo)                   // 8

	color.HiMagenta("🚀 ZENITH MANGA VAULT - SERVICIO WEB INICIADO")
	color.Cyan("Servidor corriendo en http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// --- HANDLERS (SERVICIOS) ---

func listarTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// USAMOS EL PAQUETE INTERFACES AQUÍ PARA QUE NO DE ERROR
	var items []interfaces.Gestionable = repo.ObtenerItems()

	var libros []models.Libro
	for _, item := range items {
		if l, ok := item.(*models.Libro); ok {
			libros = append(libros, *l)
		}
	}
	json.NewEncoder(w).Encode(libros)
}

func calcularTotalPaginas(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0
	// TAMBIÉN LO USAMOS AQUÍ
	var items []interfaces.Gestionable = repo.ObtenerItems()

	for _, item := range items {
		if l, ok := item.(*models.Libro); ok {
			wg.Add(1)
			go func(p int) {
				defer wg.Done()
				mu.Lock()
				total += p
				mu.Unlock()
			}(l.Paginas)
		}
	}
	wg.Wait()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"total_paginas_concurrente": total})
}

func listarPDFs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libros []models.Libro
	for _, item := range repo.ObtenerItems() {
		if l, ok := item.(*models.Libro); ok && l.Formato == "PDF" {
			libros = append(libros, *l)
		}
	}
	json.NewEncoder(w).Encode(libros)
}

func listarLargos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libros []models.Libro
	for _, item := range repo.ObtenerItems() {
		if l, ok := item.(*models.Libro); ok && l.Paginas > 400 {
			libros = append(libros, *l)
		}
	}
	json.NewEncoder(w).Encode(libros)
}

func infoSistema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"proyecto": "Zenith Manga", "unidad": "4 - Web & Concurrencia"})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor Zenith: Online")
}

func listarAutores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	autores := []string{"Kentaro Miura", "Takehiko Inoue", "Eiichiro Oda", "Katsuhiro Otomo"}
	json.NewEncoder(w).Encode(autores)
}

func verUltimo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := repo.ObtenerItems()
	if len(items) > 0 {
		json.NewEncoder(w).Encode(items[len(items)-1])
	}
}
