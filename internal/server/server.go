package server

import (
	"encoding/json"
	"html/template"
	"io"
	"l0/internal/models"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	JSON string
}

func Run() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/get-json", handleGetJSON)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		JSON: "{}",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetJSON(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	log.Println("ID:", id)

	jsonFile, err := os.Open("cmd/sender/model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var order models.Order
	_ = json.Unmarshal(byteJson, &order)

	jsonData, err := json.Marshal(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
