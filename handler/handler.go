package handler

import (
	"authentication_microservice/database"
	"authentication_microservice/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}
type data map[string]interface{}

func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		fmt.Fprintf(w, "Welcome!")
		return
	}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		JSONWriter(w, data{
			"Error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		JSONWriter(w, data{
			"Error": "No Json Found on Request",
		}, http.StatusUnprocessableEntity)
		return
	}
	err = database.CreateUser(&user, h.DB)
	if err != nil {
		JSONWriter(w, data{
			"Error": err,
		}, http.StatusInternalServerError)
		return
	}

	JSONWriter(w, data{"Success": "Account Created Successfully"}, 200)
}
func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		JSONWriter(w, data{
			"Error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	user := model.UserRequest{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		JSONWriter(w, data{
			"Error": "No Json Found on Request",
		}, http.StatusUnprocessableEntity)
		return
	}
	if user.Username == "" {
		JSONWriter(w, data{
			"Error": "Fields Can't be empty",
		}, http.StatusUnprocessableEntity)
		return
	}
	getuser, err := database.FindUserByUsername(user.Username, h.DB)
	if err != nil {
		JSONWriter(w, data{
			"Error": "internal server error",
		}, http.StatusNotFound)
		return
	}

	if getuser.Username == "" {
		JSONWriter(w, data{
			"Response": "true",
		}, http.StatusOK)
		return
	} else {
		JSONWriter(w, data{
			"Response": "false",
		}, http.StatusOK)
		return
	}

}
