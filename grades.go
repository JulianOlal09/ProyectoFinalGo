package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Estructura para representar una calificación
type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	IDSubject int     `json:"id_subject"`
	Grade     float32 `json:"grade"`
}

// Crear una nueva calificación
func InsertGrade(w http.ResponseWriter, r *http.Request) {
	var grade Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO grades (grade_id, student_id, id_subject, grade) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, grade.GradeID, grade.StudentID, grade.IDSubject, grade.Grade)
	if err != nil {
		http.Error(w, "Error al insertar la calificación", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Calificación insertada correctamente"))
}

// Actualizar una calificación
func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gradeID := vars["grade_id"]

	var grade Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "UPDATE grades SET student_id=?, id_subject=?, grade=? WHERE grade_id=?"
	_, err := db.Exec(query, grade.StudentID, grade.IDSubject, grade.Grade, gradeID)
	if err != nil {
		http.Error(w, "Error al actualizar la calificación", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Calificación actualizada correctamente"))
}

// Eliminar una calificación
func DeleteGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gradeID := vars["grade_id"]

	query := "DELETE FROM grades WHERE grade_id=?"
	_, err := db.Exec(query, gradeID)
	if err != nil {
		http.Error(w, "Error al eliminar la calificación", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Calificación eliminada correctamente"))
}

// Obtener una calificación específica
func GetGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gradeID := vars["grade_id"]
	studentID := vars["student_id"]

	var grade Grade
	query := "SELECT grade_id, student_id, id_subject, grade FROM grades WHERE grade_id=? AND student_id=?"
	err := db.QueryRow(query, gradeID, studentID).Scan(&grade.GradeID, &grade.StudentID, &grade.IDSubject, &grade.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Calificación no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener la calificación", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(grade)
}

// Obtener todas las calificaciones de un estudiante
func GetAllGrades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["student_id"]

	rows, err := db.Query("SELECT grade_id, student_id, id_subject, grade FROM grades WHERE student_id=?", studentID)
	if err != nil {
		http.Error(w, "Error al obtener las calificaciones", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var grades []Grade
	for rows.Next() {
		var grade Grade
		if err := rows.Scan(&grade.GradeID, &grade.StudentID, &grade.IDSubject, &grade.Grade); err != nil {
			http.Error(w, "Error al procesar los datos", http.StatusInternalServerError)
			return
		}
		grades = append(grades, grade)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(grades)
}
