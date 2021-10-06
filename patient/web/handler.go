package web

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"patients-api/internal/web"
	"patients-api/patient/model"
	"patients-api/patient/store"

	"github.com/go-chi/chi"
)

type Service struct {
	gtw store.PatientGateway
}

func New(db *sql.DB) *Service {
	return &Service{
		store.NewPatientGateway(db),
	}
}

func (s *Service) GetPatientsHandler(w http.ResponseWriter, r *http.Request) {
	p := s.gtw.GetPatients()
	if p == nil || len(p) == 0 {
		p = []model.Patient{}
	}
	web.Success(&p, http.StatusOK).Send(w)
}

func (s *Service) GetPatientsByIDHandler(w http.ResponseWriter, r *http.Request) {
	patientID := chi.URLParam(r, "patientID")
	id, _ := strconv.ParseInt(patientID, 10, 64)
	patient, err := s.gtw.GetPatientByID(id)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&patient, http.StatusOK).Send(w)
}

func (s *Service) CreatePatientsHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	var patient model.PatientInsert
	err := json.NewDecoder(body).Decode(&patient)

	if err != nil {
		web.ErrInvalidJSON.Send(w)
		return
	}

	err = s.gtw.CreatePatient(patient)
	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&patient, http.StatusOK).Send(w)
}
