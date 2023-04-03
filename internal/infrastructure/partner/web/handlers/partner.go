package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
	deleteUsecase "github.com/eltoncasacio/vantracking/internal/usecase/partner/delete"
	findbyid "github.com/eltoncasacio/vantracking/internal/usecase/partner/findbyid"
	listUsecase "github.com/eltoncasacio/vantracking/internal/usecase/partner/list"
	listByCityUsecase "github.com/eltoncasacio/vantracking/internal/usecase/partner/list_by_city"
	registerUsecase "github.com/eltoncasacio/vantracking/internal/usecase/partner/register"
	updateUsecase "github.com/eltoncasacio/vantracking/internal/usecase/partner/update"
	"github.com/go-chi/chi"
)

type partnerHandler struct {
	repository repo.PartnerRepositoryInterface
}

func NewPartnerHandler(repo repo.PartnerRepositoryInterface) *partnerHandler {
	return &partnerHandler{repository: repo}
}

func (h *partnerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerUsecase.PartnerInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecaseInput := registerUsecase.PartnerInput{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
		Complement:  input.Complement,
	}

	err = registerUsecase.NewUseCase(h.repository).Register(usecaseInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *partnerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	partner, err := findbyid.NewUseCase(dh.repository).FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(partner)
	w.WriteHeader(http.StatusOK)
}

func (dh *partnerHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input updateUsecase.PartnerInputDTO
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

func (dh *partnerHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *partnerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := listUsecase.NewUseCase(h.repository).ListAll()

	output := []listUsecase.PartnerOutputDTO{}
	for _, partner := range usecaseOutput {
		d := listUsecase.PartnerOutputDTO{
			ID:          partner.ID,
			Name:        partner.Name,
			Description: partner.Description,
			Price:       partner.Price,
			PhoneNumber: partner.PhoneNumber,
			UF:          partner.UF,
			City:        partner.City,
			Street:      partner.Street,
			Number:      partner.Number,
			CEP:         partner.CEP,
			Complement:  partner.Complement,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *partnerHandler) FindByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecaseOutput, _ := listByCityUsecase.NewUseCase(h.repository).ListByCity(city)

	output := []listUsecase.PartnerOutputDTO{}
	for _, partner := range usecaseOutput {
		d := listUsecase.PartnerOutputDTO{
			ID:          partner.ID,
			Name:        partner.Name,
			Description: partner.Description,
			Price:       partner.Price,
			PhoneNumber: partner.PhoneNumber,
			UF:          partner.UF,
			City:        partner.City,
			Street:      partner.Street,
			Number:      partner.Number,
			CEP:         partner.CEP,
			Complement:  partner.Complement,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}
