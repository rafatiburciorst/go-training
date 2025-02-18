package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Erro ao fazer marchal de json", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Erro ao enviar a resposta", "error", err)
		return
	}
}

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password string `json:"-"`
}

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	l.Info("Serviço iniciado", "time", time.Now(), "version", "1.0.0")
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/horario", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		fmt.Fprintln(w, now)
	})

	db := map[int64]User{
		1: {
			Username: "admin",
			Password: "admin",
			Role:     "admin",
			ID:       1,
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "Usuário não encontrado"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {
				slog.Error("body too large", "error", err)
				sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
				return
			}
			slog.Error("body too large", "error", err)
			sendJSON(w, Response{Error: "body too large"}, http.StatusInternalServerError)
			return
		}

		var user User

		if err := json.Unmarshal(data, &user); err != nil {
			slog.Error("body too large", "error", err)
			sendJSON(w, Response{Error: "body too large"}, http.StatusUnprocessableEntity)
			return
		}

		db[user.ID] = user
		w.WriteHeader(http.StatusCreated)
	}
}
