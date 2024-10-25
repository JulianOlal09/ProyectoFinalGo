package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB // Definición global de db

func main() {
	var err error
	dsn := "root:toor@tcp(localhost:3306)/escuela"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Rutas para estudiantes
	r.HandleFunc("/students", InsertStudent).Methods("POST")                // Crear estudiante
	r.HandleFunc("/students/{student_id}", GetStudent).Methods("GET")       // Obtener un estudiante
	r.HandleFunc("/students/{student_id}", UpdateStudent).Methods("PUT")    // Actualizar estudiante
	r.HandleFunc("/students/{student_id}", DeleteStudent).Methods("DELETE") // Eliminar estudiante
	r.HandleFunc("/students", GetAllStudents).Methods("GET")                // Obtener todos los estudiantes

	log.Println("Servidor corriendo en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
