package store

import (
	"database/sql"
	"patients-api/patient/model"
)

type PatientGateway interface {
	CreatePatient(p model.PatientInsert) error
	GetPatients() []model.Patient
	GetPatientByID(id int64) (model.Patient, error)
}

type CreatePatientInDB struct {
	Store
}

func NewPatientGateway(db *sql.DB) PatientGateway {
	return &CreatePatientInDB{New(db)}
}

func (c *CreatePatientInDB) CreatePatient(p model.PatientInsert) error {
	return c.insert(p)
}

func (c *CreatePatientInDB) GetPatients() []model.Patient {
	return c.loadAll()
}

func (c *CreatePatientInDB) GetPatientByID(id int64) (model.Patient, error) {
	return c.loadOne(id)
}
