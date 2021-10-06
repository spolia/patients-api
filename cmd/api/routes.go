package api

import (
	"patients-api/patient/web"

	"github.com/go-chi/chi"
)

func routes(s web.Service) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", s.GetPatientsHandler)
	r.Post("/patients", s.CreatePatientsHandler)
	r.Get("/patients/{patientID}", s.GetPatientsByIDHandler)

	return r
}
