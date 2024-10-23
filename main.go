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
	dsn := "root:Anettevira26?@tcp(localhost:3307)/escuela"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Rutas para estudiantes

	r.HandleFunc("/students/{student_id}", GetStudent).Methods("GET")
	r.HandleFunc("/students/{student_id}", GetAllStudents).Methods("PUT")
	r.HandleFunc("/students/{student_id}", DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students", GetAllStudents).Methods("GET")

	// Rutas para calificaciones (grades)
	r.HandleFunc("/grades", InsertGrade).Methods("POST")
	r.HandleFunc("/grades/{grade_id}", UpdateGrade).Methods("PUT")
	r.HandleFunc("/grades/{grade_id}", DeleteGrade).Methods("DELETE")
	r.HandleFunc("/grades/{grade_id}/student/{student_id}", GetGrade).Methods("GET")
	r.HandleFunc("/grades/student/{student_id}", GetAllGrades).Methods("GET")

	// Rutas para materias (subjects)
	r.HandleFunc("/subjects", GetAllSubjects).Methods("POST")
	r.HandleFunc("/subjects/{id}", GetAllSubjects).Methods("PUT")
	r.HandleFunc("/subjects/{id}", DeleteSubject).Methods("DELETE")
	r.HandleFunc("/subjects/{id}", GetSubject).Methods("GET")
	r.HandleFunc("/subjects", GetAllSubjects).Methods("GET")

	log.Println("Servidor corriendo en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
