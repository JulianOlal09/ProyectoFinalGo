package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Subject struct {
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
}

// Crear una nueva materia
func CreateSubject(w http.ResponseWriter, r *http.Request) {
	var subject Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO subjects (name) VALUES (?)"
	_, err := db.Exec(query, subject.Name)
	if err != nil {
		http.Error(w, "Error al crear la materia: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Materia creada correctamente"))
}

// Actualizar materia
func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subject_id"]

	var subject Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "UPDATE subjects SET name=? WHERE subject_id=?"
	_, err := db.Exec(query, subject.Name, subjectID)
	if err != nil {
		http.Error(w, "Error al actualizar la materia: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Materia actualizada correctamente"))
}

// Obtener materia por ID
func GetSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subject_id"]

	var subject Subject
	query := "SELECT subject_id, name FROM subjects WHERE subject_id=?"
	err := db.QueryRow(query, subjectID).Scan(&subject.SubjectID, &subject.Name)
	if err != nil {
		http.Error(w, "Materia no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subject)
}

// Obtener todas las materias
func GetAllSubjects(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT subject_id, name FROM subjects")
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var subjects []Subject
	for rows.Next() {
		var subject Subject
		if err := rows.Scan(&subject.SubjectID, &subject.Name); err != nil {
			http.Error(w, "Error al procesar las materias", http.StatusInternalServerError)
			return
		}
		subjects = append(subjects, subject)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subjects)
}

// Eliminar materia
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subject_id"]

	query := "DELETE FROM subjects WHERE subject_id=?"
	_, err := db.Exec(query, subjectID)
	if err != nil {
		http.Error(w, "Error al eliminar la materia", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Materia eliminada correctamente"))
}
