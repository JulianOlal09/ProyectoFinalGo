package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Estructura para representar un estudiante
type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

// Actualizar un estudiante
func GetAlltStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO students SET name=?, age=? WHERE student_id=?"
	_, err := db.Exec(query, student.Name, student.Age, studentID)
	if err != nil {
		http.Error(w, "Error al actualizar el estudiante", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Estudiante actualizado correctamente"))
}

// Eliminar un estudiante
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	query := "DELETE FROM students WHERE student_id=?"
	_, err := db.Exec(query, studentID)
	if err != nil {
		http.Error(w, "Error al eliminar el estudiante", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Estudiante eliminado correctamente"))
}

// Obtener un estudiante específico
func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	var student Student
	query := "SELECT student_id, name, age FROM students WHERE student_id=?"
	err := db.QueryRow(query, studentID).Scan(&student.StudentID, &student.Name, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener el estudiante", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}
