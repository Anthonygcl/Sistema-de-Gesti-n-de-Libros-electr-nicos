# Sistema-de-Gesti-n-de-Libros-electr-nicos
¡Bienvenido a Libroteca. Este es un sistema de gestión de libros electrónicos desarrollado en **Go**, diseñado para demostrar el poder de la **programación funcional** aplicada a la organización de contenidos digitales, utilizando como base títulos icónicos del manga japonés.

## 🚀 Objetivo del Proyecto
Desarrollar una herramienta CLI (Command Line Interface) modular que permita gestionar inventarios de e-books, aplicando los conocimientos de la **Unidad 1** de Go, tales como:
- Sintaxis y tipado.
- Estructuras y control de flujo.
- Manejo de paquetes.
- Funciones de orden superior (Programación Funcional).

## 🛠️ Tecnologías y Paquetes
Para este proyecto se utilizaron las siguientes herramientas:
- **Lenguaje:** [Go (Golang)](https://go.dev/)
- **Paquetes de Terceros:**
  - `fatih/color`: Para jerarquización visual mediante colores en la terminal.
  - `olekukonko/tablewriter`: Para la generación de tablas estructuradas de datos.

## 🧩 Estructura del Proyecto
```text
AP1/
├── main.go             # Punto de entrada y menú interactivo
├── internal/
│   ├── logic/          # Motor funcional (filtros y predicados)
│   └── models/         # Definición de estructuras de datos (Libro)
└── docs/               # Documentación PDF y alcance
