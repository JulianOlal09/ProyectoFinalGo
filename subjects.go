package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Subject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
}

func GetAllSubjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["id"]

	var subject Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO subjects SET name=?, teacher=? WHERE id=?"
	_, err := db.Exec(query, subject.Name, subject.Teacher, subjectID)
	if err != nil {
		http.Error(w, "Error al actualizar la materia", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Materia actualizada correctamente"))
}

// Eliminar una materia
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["id"]

	query := "DELETE FROM subjects WHERE id=?"
	_, err := db.Exec(query, subjectID)
	if err != nil {
		http.Error(w, "Error al eliminar la materia", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Materia eliminada correctamente"))
}

// Obtener una materia específica
func GetSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["id"]

	var subject Subject
	query := "SELECT id, name, teacher FROM subjects WHERE id=?"
	err := db.QueryRow(query, subjectID).Scan(&subject.ID, &subject.Name, &subject.Teacher)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Materia no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener la materia", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subject)
}
