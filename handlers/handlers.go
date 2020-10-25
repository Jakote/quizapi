package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/EduAR/quizapi/repository"
)

type Handlers struct {
	Repo *repository.Repository
}

func (h *Handlers) All(w http.ResponseWriter, r *http.Request) {
	qqzs, err := h.Repo.FindAll()
	if err != nil {
		error500(w, err)
		return
	}
	jr, err := json.Marshal(qqzs)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) GetByCat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctgry, ok := params["category"]
	if !ok {
		error400(w, "category is Required.")
		return
	}
	qqz, err := h.Repo.FindByCategory(ctgry)
	if err != nil {
		error404(w, "category not found.")
		return
	}
	jr, err := json.Marshal(qqz)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctgry, ok := params["category"]
	if !ok {
		error400(w, "category is Required.")
		return
	}

	qqz := repository.QuizQuestion{
		Category: ctgry,
	}

	err := h.Repo.Delete(qqz)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully deleted.")
}

func (h *Handlers) Insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var qqz repository.QuizQuestion
	err := json.NewDecoder(r.Body).Decode(&qqz)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Insert(qqz)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully inserted.")
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctgry, ok := params["category"]
	if !ok {
		error400(w, "category is Required.")
		return
	}

	var qqz repository.QuizQuestion
	err := json.NewDecoder(r.Body).Decode(&qqz)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Update(ctgry, qqz)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully updated.")
}
