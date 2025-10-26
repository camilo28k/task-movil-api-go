package main

import (
	"log"
	"net/http"
	"os"

	"crud-backend/database"

	// ====== Usuarios ======
	usuarioEntity "crud-backend/usuarios/entity"
	usuarioHandlerPkg "crud-backend/usuarios/handler"
	usuarioRepoPkg "crud-backend/usuarios/repository"
	usuarioServicePkg "crud-backend/usuarios/service"

	// ====== Tareas ======
	tareaEntity "crud-backend/tareas/entity"
	tareaHandlerPkg "crud-backend/tareas/handler"
	tareaRepoPkg "crud-backend/tareas/repository"
	tareaServicePkg "crud-backend/tareas/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// 1️⃣ Cargar variables del entorno (.env)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró archivo .env, usando variables del entorno")
	}

	// 2️⃣ Conectar a la base de datos PostgreSQL
	database.Connect()

	// 3️⃣ Migraciones automáticas (crea las tablas si no existen)
	if err := database.DB.AutoMigrate(
		&usuarioEntity.Usuario{},
		&tareaEntity.Tarea{},
	); err != nil {
		log.Fatalf("❌ Error en migraciones: %v", err)
	}
	log.Println("✅ Migraciones completadas correctamente")

	// 4️⃣ Inicializar Repositorios, Servicios y Handlers

	// Usuarios
	usuarioRepo := usuarioRepoPkg.NewUsuarioRepository()
	usuarioService := usuarioServicePkg.NewUsuarioService(usuarioRepo)
	usuarioHandler := usuarioHandlerPkg.NewUsuarioHandler(usuarioService)

	// Tareas
	tareaRepo := tareaRepoPkg.NewTareaRepository()
	tareaService := tareaServicePkg.NewTareaService(tareaRepo)
	tareaHandler := tareaHandlerPkg.NewTareaHandler(tareaService)

	// 5️⃣ Configurar Router principal
	r := mux.NewRouter()

	// Registrar rutas
	usuarioHandler.RegisterRoutes(r)
	tareaHandler.RegisterRoutes(r)

	// 6️⃣ Configurar CORS (para permitir peticiones desde el frontend)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8100", "http://localhost:5173"}) // React o Ionic

	// 7️⃣ Puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // puerto interno del contenedor o local
	}

	// 8️⃣ Iniciar servidor
	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(r)))
}
