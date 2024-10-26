package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	dsn := "root:abigail30@tcp(127.0.0.1:3306)/escuela?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/subjects", CreateSubject).Methods("POST")
	r.HandleFunc("/api/subjects/{subject_id}", UpdateSubject).Methods("PUT")
	r.HandleFunc("/api/subjects/{subject_id}", DeleteSubject).Methods("DELETE")
	r.HandleFunc("/api/subjects/{subject_id}", GetSubject).Methods("GET")
	r.HandleFunc("/api/subjects", GetAllSubjects).Methods("GET")

	// Ejecutar servidor
	log.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
