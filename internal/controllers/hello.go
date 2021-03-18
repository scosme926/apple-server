package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/scosme926/apple-server/internal/models"
)

func (c *Controller) postHello(w http.ResponseWriter, r *http.Request) {
	var requestData models.HelloRequest

	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := requestData.Name
	var responseData = &models.HelloResponse{
		Message: "Hello" + name,
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
