# Sistema-de-Gestion-de-Libros-electronicos
Este proyecto representa la culminación de la asignatura **Programación en Go**. Es una API REST concurrente diseñada para la gestión de libros electrónicos temáticos, integrando los paradigmas funcional y orientado a objetos, junto con servicios web y pruebas unitarias.

## 🛠️ Tecnologías Utilizadas
- **Lenguaje:** Go (Golang) 1.21+
- **Protocolo:** HTTP / JSON
- **Concurrencia:** Goroutines, WaitGroups y RWMutex
- **Testing:** Paquete nativo `testing` de Go
- **Librerías:** `fatih/color`, `olekukonko/tablewriter`

## 📂 Estructura del Proyecto
- `internal/models`: Entidades core y encapsulación.
- `internal/interfaces`: Contratos de polimorfismo.
- `internal/logic`: Lógica de biblioteca y control de concurrencia.
- `tests`: Pruebas automatizadas de calidad.
- `main.go`: Punto de entrada y definición de los 8 servicios web.

## 🚀 Instalación y Ejecución
1. Clonar el repositorio: `git clone [URL-DE-TU-REPOSITO]`
2. Instalar dependencias: `go mod tidy`
3. Iniciar el servidor: `go run main.go`

## 🔗 Servicios Web Disponibles (Endpoints)
1. `GET /api/libros`: Lista completa en JSON.
2. `GET /api/stats/paginas`: Cálculo concurrente de páginas (Goroutines).
3. `GET /api/libros/pdf`: Filtrado por formato.
4. `GET /api/libros/largos`: Libros de más de 400 páginas.
5. `GET /api/info`: Metadatos del sistema.
6. `GET /api/health`: Estado del servidor.
7. `GET /api/autores`: Listado de mangakas.
8. `GET /api/ultimo`: Último registro ingresado.

## 🧩 Estructura del Proyecto
AP1/
├── main.go                 # Servidor Web, 8 Endpoints y Concurrencia
├── go.mod                  # Definición del módulo (module AP1)
├── go.sum                  # Registro de librerías (color, tablewriter)
├── internal/
│   ├── models/
│   │   └── libro.go        # Structs con JSON, Encapsulación y Constructor
│   ├── logic/
│   │   ├── biblioteca.go   # Clase Gestora con RWMutex (Concurrencia segura)
│   │   └── procesamiento.go # Lógica de filtrado (Continuidad Unidad 1 y 2)
│   └── interfaces/
│       └── gestion.go      # Contrato de la Interfaz
├── tests/
│   └── unitaria_test.go    # Pruebas Unitarias de calidad
