package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserHandler struct {
	sync.Mutex
	store map[string]User
}

func (h *UserHandler) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

}
func (h *UserHandler) post(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	bodyByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var user User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
	}

	h.Lock()
	h.store[user.Email] = user
	defer h.Unlock()
}
func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	users := make([]User, len(h.store))
	i := 0

	h.Lock()
	for _, user := range h.store {
		users[i] = user
		i++
	}

	h.Unlock()

	jsonByte, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	log.Printf("[REST API] Response : %s", jsonByte)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		store: map[string]User{
			"id1": {
				Name:    "Sabituddin bigbang",
				Email:   "sabituddin22@gmail.com",
				Address: "Jl. Merdeka Barat, Jakarta",
			},

			"id2": {
				Name:    "Anonymous",
				Email:   "anon@gmail.com",
				Address: "Jl. Merdeka Timur, Jakarta",
			},
		},
	}
}

func Server() {
	userHandler := NewUserHandler()
	http.HandleFunc("/users", userHandler.users)

	log.Println("[REST API] trying to listen localhost:6969/users ...")
	err := http.ListenAndServe(":6969", nil)

	if err != nil {
		panic(err)
	}
}
