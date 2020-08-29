package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	db "LoginServer/db"
	model "LoginServer/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetPeople is an httpHandler for route GET /people
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)

	res := db.Register(user)
	if err != nil {
		res.Error = err.Error()
	}
	json.NewEncoder(w).Encode(res)
}

// GetPerson is an httpHandler for route GET /people/{id}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	res, result := db.Login(user)
	if res.Error == "" {
		json.NewEncoder(w).Encode(result)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	tokenString := r.Header.Get("Authorization")

	result, _ := db.Profile(tokenString)

	json.NewEncoder(w).Encode(bson.M{"data": result})

}

func GetHostID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := r.Header.Get("Authorization")
	token := strings.Split(tokenString, " ")
	result, _ := db.GetHosts(token[1])
	data := bson.M{"data": result}
	json.NewEncoder(w).Encode(data)

}
