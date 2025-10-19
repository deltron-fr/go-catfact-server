package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/deltron-fr/hng-go/internal/catapi"
)

const url = "https://catfact.ninja/fact"

var (
	emailName = "youremail@gmail.com"
	fullname = "Dave"
)

type User struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Stack string `json:"stack"`
}

type Response struct {
	Status string `json:"status"`
	User `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact string `json:"fact"`
}


func HandleMe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	catFact, err := catapi.GetCatFacts(url)
	if err != nil {
		http.Error(w, "error getting cat facts", http.StatusInternalServerError)
		log.Printf("error getting cat facts: %v", err)
		return
	}

	currentTime := time.Now().UTC()
	timeISO := currentTime.Format(time.RFC3339)

	r := Response{
		Status: "success",
		User: User{
			Email: emailName,
			Name: fullname,
			Stack: "golang",
		},
		Timestamp: timeISO,
		Fact: catFact.Fact,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		http.Error(w, "error marshaling json", http.StatusInternalServerError)
		log.Printf("error marshaling json: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}