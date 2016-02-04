package main

import (
	"encoding/json"
	//"html"
	"io"
	"io/ioutil"
	"log"
	//"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var dal DAL

func init() {
	dal = DAL{InMemStorage{make(map[string]APIUser)}}
}

func Index(w http.ResponseWriter, r *http.Request) {
	printResponse(w, http.StatusOK, APIMessage{http.StatusOK, "Welcome!"})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user APIUser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		printResponse(w, 422, APIMessage{422, "Unprocessable entity"})
	}

	if !isRegistered(user) {
		dal.CreateUser(user)
		log.Printf("New user registered: %q\n", user)
		printResponse(w, http.StatusOK, APIMessage{http.StatusOK, user})
	} else {
		log.Printf("User already registered: %q\n", user)
		printResponse(w, http.StatusConflict, APIMessage{http.StatusConflict, "User already registered"})
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	if user, err := dal.FindUser(APIUser{"", "", token}); err == nil {
		log.Printf("User[%q] found: %q\n", token, user)
		printResponse(w, http.StatusOK, APIMessage{http.StatusOK, user})
	} else {
		log.Printf("User[%q] not found\n", token)
		printResponse(w, http.StatusNotFound, APIMessage{http.StatusNotFound, "Not found"})
	}
}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {
	var user APIUser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		printResponse(w, 422, APIMessage{422, "Unprocessable entity"})
	}
	printResponse(w, http.StatusOK, APIMessage{http.StatusOK, user})
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	//printResponse(w, http.StatusOK, APIMessage{http.StatusOK, dal.FindUser(APIUser{})})
}

func printResponse(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

func isRegistered(user APIUser) bool {
	if _, err := dal.FindUser(user); err != nil {
		return false
	}
	return true
}

func sleep(tts int64) {
	log.Print("sleep period is about to begin ... \n")
	timeToSleep := time.Duration(tts) * time.Millisecond
	time.Sleep(timeToSleep)
	log.Print("sleep period is over ... \n")
}
