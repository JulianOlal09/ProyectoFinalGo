package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Email     string `json:"email"`
}

// Nuevo estudiante
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO students (name, `group`, email) VALUES (?, ?, ?)"
	_, err := db.Exec(query, student.Name, student.Group, student.Email)
	if err != nil {
		http.Error(w, "Error al insertar el estudiante: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Estudiante insertado correctamente"))
}

// Actualizar estudiante
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "UPDATE students SET name=?, `group`=?, email=? WHERE student_id=?"
	_, err := db.Exec(query, student.Name, student.Group, student.Email, studentID)
	if err != nil {
		http.Error(w, "Error al actualizar el estudiante: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Estudiante actualizado correctamente"))
}

// Eliminar estudiante
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	query := "DELETE FROM students WHERE student_id=?"
	_, err := db.Exec(query, studentID)
	if err != nil {
		http.Error(w, "Error al eliminar el estudiante: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Estudiante eliminado correctamente"))
}

// Obtener estudiante específico
func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	var student Student
	query := "SELECT student_id, name, `group`, email FROM students WHERE student_id=?"
	err := db.QueryRow(query, studentID).Scan(&student.StudentID, &student.Name, &student.Group, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener el estudiante: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

// Obtener todos los estudiantes
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT student_id, name, `group`, email FROM students")
	if err != nil {
		http.Error(w, "Error al obtener los estudiantes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.StudentID, &student.Name, &student.Group, &student.Email); err != nil {
			http.Error(w, "Error al procesar los datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}
