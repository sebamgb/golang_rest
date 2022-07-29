package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-go/models"
	"rest-go/repository"
	"rest-go/server"

	"github.com/segmentio/ksuid"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		client := &http.Client{}
		url := "http://20.106.99.61" + r.URL.String()
		requestJson, err := json.Marshal(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req, err := http.NewRequest(r.Method, url, bytes.NewBuffer(requestJson))
		if err != nil {
			log.Fatalf("request failed when it creating with error: %v", err)
		}
		req.Header.Add("Content-Type", "application/json")
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("request failed when it doing with error: %v", err)
		}
		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("error when reading response: %v", err)
		}
		stringResponse := string(responseBody)
		log.Printf("response code: %d\n", response.StatusCode)
		log.Printf("header: '%q'\n", response.Header)
		log.Println(stringResponse)
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(req.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
