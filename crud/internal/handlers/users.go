package handlers

import (
	"crud/internal/models"
	"encoding/json"
	"net/http"
)

func ListUsers(db map[int64]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{
			Name: "Rafael",
			Age:  37,
		}

		res, _ := json.Marshal(user)

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

}
