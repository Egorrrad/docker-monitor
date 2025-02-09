package main

import (
	"Backend/internal/model"
	"encoding/json"
	"net/http"
)

func (app *application) getStatus(w http.ResponseWriter, r *http.Request) {
	statuses, err := app.Store.GetAllStatuses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (app *application) addStatus(w http.ResponseWriter, r *http.Request) {
	var status model.ContainerStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.Store.AddStatus(status.IP, status.Alive, status.Checked, status.LastSuccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
