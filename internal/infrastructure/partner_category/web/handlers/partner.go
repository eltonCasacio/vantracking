package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
	deleteUsecase "github.com/eltoncasacio/vantracking/internal/usecase/category/delete"
	findbyid "github.com/eltoncasacio/vantracking/internal/usecase/category/findbyid"
	listUsecase "github.com/eltoncasacio/vantracking/internal/usecase/category/list"
	registerUsecase "github.com/eltoncasacio/vantracking/internal/usecase/category/register"
	updateUsecase "github.com/eltoncasacio/vantracking/internal/usecase/category/update"
	"github.com/go-chi/chi"
)

type categoryHandler struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewPartnerHandler(repo repo.PartnerCategoryRepositoryInterface) *categoryHandler {
	return &categoryHandler{repository: repo}
}

func (h *categoryHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerUsecase.CategoryInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecaseInput := registerUsecase.CategoryInput{
		Name: input.Name,
	}

	err = registerUsecase.NewUseCase(h.repository).Register(usecaseInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *categoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	partner, err := findbyid.NewUseCase(dh.repository).FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(partner)
	w.WriteHeader(http.StatusOK)
}

func (dh *categoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input updateUsecase.CategoryInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = updateUsecase.NewUseCase(dh.repository).Update(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *categoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := deleteUsecase.NewUseCase(dh.repository).Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *categoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := listUsecase.NewUseCase(h.repository).ListAll()

	output := []listUsecase.CategoryOutpu{}
	for _, partner := range usecaseOutput {
		d := listUsecase.CategoryOutpu{
			ID:   partner.ID,
			Name: partner.Name,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}
